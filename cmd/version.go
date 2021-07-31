package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "v0.1.8"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version prints the current version of hail",

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Print version basis the latest release.
		fmt.Println(version)
	},
}

func init() {
}
