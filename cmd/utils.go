package cmd

import (
	"errors"
	"log"
	"net"
	"os"
	"strconv"

	"golang.org/x/crypto/ssh"
)

// ConfigureAuthentication configures the authentication method.
func ConfigureAuthentication(key string, passphrase string, password string) []ssh.AuthMethod {
	// Create signer for public key authentication method.
	auth := make([]ssh.AuthMethod, 1)

	if key != "" {
		var err error
		var targetSigner ssh.Signer

		if passphrase != "" {
			targetSigner, err = ssh.ParsePrivateKeyWithPassphrase([]byte(key), []byte(passphrase))
		} else {
			targetSigner, err = ssh.ParsePrivateKey([]byte(key))
		}

		if err != nil {
			log.Fatalf("❌ Failed to parse private key: %v", err)
		}

		// Configure public key authentication.
		auth[0] = ssh.PublicKeys(targetSigner)
	} else if password != "" {
		// Fall back to password authentication.
		auth[0] = ssh.Password(password)
		log.Println("⚠️ Using a password for authentication is insecure!")
		log.Println("⚠️ Please consider using public key authentication!")
	} else {
		log.Fatal("❌ Failed to configure authentication method: missing credentials")
	}

	return auth
}

// ConfigureHostKeyCallback configures the SSH host key verification callback.
// Unless the `skip` option is set to `string("true")` it will return a function,
// which verifies the host key against the specified ssh key fingerprint.
func ConfigureHostKeyCallback(expected string, skip bool) ssh.HostKeyCallback {
	if skip {
		log.Println("⚠️ Skipping host key verification is insecure!")
		log.Println("⚠️ This allows for person-in-the-middle attacks!")
		log.Println("⚠️ Please consider using host key verification!")
		return ssh.InsecureIgnoreHostKey() //nolint:gosec
	}

	return func(hostname string, remote net.Addr, pubKey ssh.PublicKey) error {
		fingerprint := ssh.FingerprintSHA256(pubKey)
		if fingerprint != expected {
			return errors.New("fingerprint mismatch: server fingerprint: " + fingerprint)
		}

		return nil
	}
}

func AtoiWithDefault(s string, n int) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		return n
	}
	return i
}

func GetEnvWithDefault(s, def string) string {
	if v, ok := os.LookupEnv(s); ok {
		return v
	}
	return def
}
