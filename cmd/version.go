package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "v0.1.0"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version prints the current version of hail",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
