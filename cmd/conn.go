/*
Copyright © 2024 Infovista
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/dtylman/scp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CopyFunc func(client *ssh.Client, source string, target string) (int64, error)

func copyFiles(client *ssh.Client, sourceFiles []string, targetFileOrFolder string, performUpload bool) error {
	var copyFunction CopyFunc
	var emoji string
	var actionText string

	if performUpload {
		copyFunction = scp.CopyTo
		emoji = "🔼"
		actionText = "upload"
	} else {
		copyFunction = scp.CopyFrom
		emoji = "🔽"
		actionText = "download"
	}

	caser := cases.Title(language.AmericanEnglish)
	log.Printf("%s %sing ...\n", emoji, caser.String(actionText))
	if len(sourceFiles) == 1 {
		// Rename file if there is only one source file.
		if _, err := copyFunction(client, sourceFiles[0], targetFileOrFolder); err != nil {
			log.Printf("❌ Failed to %s file from remote: %v", actionText, err)
			return err
		}
		log.Println("📑 " + sourceFiles[0] + " >> " + targetFileOrFolder)

		log.Println("📡 Transferred 1 file")
	} else {
		transferredFiles := int64(0)

		for _, sourceFile := range sourceFiles {
			_, file := path.Split(sourceFile)
			targetFile := path.Join(targetFileOrFolder, file)

			if _, err := copyFunction(client, sourceFile, targetFile); err != nil {
				log.Printf("❌ Failed to %s file from remote: %v", actionText, err)
			}
			log.Println("📑 " + sourceFile + " >> " + targetFile)

			transferredFiles += 1
		}

		log.Printf("📡 Transferred %d files\n", transferredFiles)
	}
	return nil
}

func mainCopy(sourceFiles []string, destination string, performUpload bool) error {
	// Parse timeout.
	timeout, err := time.ParseDuration(Timeout)
	if err != nil {
		log.Fatalf("❌ Failed to parse timeout: %v", err)
	}

	if PassEnv != "" {
		Password = os.Getenv(PassEnv)
	}
	if Password == "" && Key == "" {
		log.Fatal("❌ Failed to configure authentication method: missing credentials")
	}

	// Create configuration for SSH target.
	targetConfig := &ssh.ClientConfig{
		Timeout:         timeout,
		User:            Username,
		Auth:            ConfigureAuthentication(Key, Passphrase, Password),
		HostKeyCallback: ConfigureHostKeyCallback(Fingerprint, SkipHostKeyVerification),
	}

	// Configure target address.
	targetAddress := fmt.Sprintf("%s:%d", Host, Port)

	// Initialize target SSH client.
	var targetClient *ssh.Client

	// Check if a proxy should be used.
	if Proxy != "" {
		// Create SSH config for SSH proxy.
		proxyConfig := &ssh.ClientConfig{
			Timeout:         timeout,
			User:            ProxyUsername,
			Auth:            ConfigureAuthentication(ProxyKey, ProxyPassphrase, ProxyPassword),
			HostKeyCallback: ConfigureHostKeyCallback(ProxyFingerprint, ProxySkipHostKeyVerification),
		}

		// Establish SSH session to proxy host.
		proxyAddress := fmt.Sprintf("%s:%d", Proxy, ProxyPort)
		proxyClient, err2 := ssh.Dial("tcp", proxyAddress, proxyConfig)
		if err2 != nil {
			log.Fatalf("❌ Failed to connect to proxy: %v", err)
		}
		defer proxyClient.Close()

		// Create a TCP connection to from the proxy host to the target.
		netConn, err2 := proxyClient.Dial("tcp", targetAddress)
		if err2 != nil {
			log.Printf("❌ Failed to dial to target: %v", err2)
			return err2
		}

		targetConn, channel, req, err3 := ssh.NewClientConn(netConn, targetAddress, targetConfig)
		if err3 != nil {
			log.Printf("❌ Failed to connect to target: %v", err3)
			return err3
		}

		targetClient = ssh.NewClient(targetConn, channel, req)
	} else {
		if targetClient, err = ssh.Dial("tcp", targetAddress, targetConfig); err != nil {
			log.Printf("❌ Failed to connect to target: %v", err)
			return err
		}
	}
	defer targetClient.Close()

	return copyFiles(targetClient, sourceFiles, destination, performUpload)
}
