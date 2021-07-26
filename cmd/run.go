package cmd

import (
	"hail/internal/editor"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [alias]",
	Short: "it is used to directly run a command from alias",
	Run: func(cmd *cobra.Command, args []string) {
		e := editor.NewDefaultEditor([]string{})
		e.RunScript()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
