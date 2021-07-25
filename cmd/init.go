package cmd

import (
	"fmt"
	"hail/internal/hailconfig"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [title]",
	Short: "init initializes an empty .hailconfig file with title as provided",
	Run: func(cmd *cobra.Command, args []string) {
		title := ""
		if len(args) < 1 {
			title = "default"
		} else {
			title = args[0]
		}
		cfgfile, err := hailconfig.Init(title)
		checkError("error in init", err)
		success(fmt.Sprintf("Initialized a file '%s'", cfgfile))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
