/*
Copyright © 2024 Infovista
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	Host                    string
	Port                    int
	Timeout                 string
	Key                     string
	Passphrase              string
	Password                string
	PassEnv                 string
	SkipHostKeyVerification bool
	Username                string
	Fingerprint             string
	// proxy
	Proxy                        string
	ProxyUsername                string
	ProxyPassword                string
	ProxyPort                    int
	ProxyKey                     string
	ProxyPassphrase              string
	ProxySkipHostKeyVerification bool
	ProxyFingerprint             string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scp-binary",
	Short: "A single-binary SCP client.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Host, "host", "H", os.Getenv("SCP_HOST"),
		"host to connect to ($SCP_HOST)")
	rootCmd.MarkPersistentFlagRequired("host")
	rootCmd.PersistentFlags().IntVar(&Port, "port", AtoiWithDefault(os.Getenv("SCP_PORT"), 22),
		"port to connect to ($SCP_PORT/22)")
	rootCmd.PersistentFlags().StringVar(&Timeout, "timeout", GetEnvWithDefault("SCP_TIMEOUT", "30s"),
		"timeout in seconds ($SCP_TIMEOUT/30s)")
	rootCmd.PersistentFlags().StringVarP(&Username, "username", "u", os.Getenv("SCP_USERNAME"),
		"username ($SCP_USERNAME)")
	rootCmd.MarkPersistentFlagRequired("username")
	rootCmd.PersistentFlags().StringVarP(&Password, "password", "p", "",
		"password ($INSECURE_PASSWORD)")
	rootCmd.PersistentFlags().StringVar(&PassEnv, "passenv", "INSECURE_PASSWORD",
		"envvar name containing password")
	// host key
	rootCmd.PersistentFlags().StringVar(&Key, "key", os.Getenv("SCP_KEY"),
		"private key file ($SCP_KEY)")
	rootCmd.PersistentFlags().StringVar(&Passphrase, "passphrase", os.Getenv("SCP_KEY_PASSPHRASE"),
		"private key passphrase ($SCP_KEY_PASSPHRASE)")
	rootCmd.PersistentFlags().StringVar(&Fingerprint, "fingerprint", os.Getenv("SCP_KEY_FINGERPRINT"),
		"expected host key fingerprint ($SCP_KEY_FINGERPRINT)")
	rootCmd.PersistentFlags().BoolVar(&SkipHostKeyVerification, "skip-fingerprint",
		os.Getenv("SCP_IGNORE_FINGERPRINT") == "true", "skip host key verification ($SCP_IGNORE_FINGERPRINT)")
	// proxy
	rootCmd.PersistentFlags().StringVar(&Proxy, "proxy", os.Getenv("SCP_PROXY"),
		"proxy ($SCP_PROXY)")
	rootCmd.PersistentFlags().IntVar(&ProxyPort, "proxy-port", AtoiWithDefault(os.Getenv("SCP_PROXY_PORT"), 22),
		"port to connect to ($SCP_PROXY_PORT/22)")
	rootCmd.PersistentFlags().StringVar(&ProxyUsername, "proxy-username", os.Getenv("SCP_PROXY_USERNAME"),
		"proxy username ($SCP_PROXY_USERNAME)")
	rootCmd.PersistentFlags().StringVar(&ProxyPassword, "proxy-password", os.Getenv("SCP_PROXY_PASSWORD"),
		"proxy password ($SCP_PROXY_PASSWORD)")
	// proxy key
	rootCmd.PersistentFlags().StringVar(&ProxyKey, "proxy-key", os.Getenv("SCP_PROXY_KEY"),
		"proxy key ($SCP_PROXY_KEY)")
	rootCmd.PersistentFlags().StringVar(&ProxyPassphrase, "proxy-key-passphrase", os.Getenv("SCP_PROXY_KEY_PASSPHRASE"),
		"private key passphrase ($SCP_PROXY_KEY_PASSPHRASE)")
	rootCmd.PersistentFlags().StringVar(&ProxyFingerprint, "proxy-key-fingerprint", os.Getenv("SCP_PROXY_KEY_FINGERPRINT"),
		"expected proxy host key fingerprint ($SCP_PROXY_KEY_FINGERPRINT)")
	rootCmd.PersistentFlags().BoolVar(&ProxySkipHostKeyVerification, "proxy-skip-fingerprint",
		os.Getenv("SCP_IGNORE_PROXY_FINGERPRINT") == "true", "skip proxy host key verification ($SCP_IGNORE_PROXY_FINGERPRINT)")
}
