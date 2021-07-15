package cmd

import (
	"hail/internal/hailconfig"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [title]",
	Short: "init initializes an empty .hailconfig file with title as provided",
	RunE: func(cmd *cobra.Command, args []string) error {
		title := ""
		if len(args) < 1 {
			title = "default"
		} else {
			title = args[0]
		}

		return hailconfig.Init(title)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
