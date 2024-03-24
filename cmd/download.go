/*
Copyright © 2024 Infovista
*/
package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download SRC DEST",
	Short: "downloads a file from host:SRC to DEST. SRC is comma separated",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Printf("downloading %s:%s -> %s\n", Host, args[0], args[1])
		return mainCopy(strings.Split(args[0], ","), args[1], false)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Args = cobra.ExactArgs(2)
}
