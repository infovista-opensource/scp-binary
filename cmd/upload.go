/*
Copyright © 2024 Infovista
*/
package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload SRC DEST",
	Short: "uploads a file from SRC to host:DEST. SRC is comma separated",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Printf("uploading %s -> %s:%s\n", args[0], Host, args[1])
		return mainCopy(strings.Split(args[0], ","), args[1], true)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Args = cobra.ExactArgs(2)
}
