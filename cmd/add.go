package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:       "add [alias] [command]",
	Short:     "add is used to add a new command in collection",
	Args:      cobra.ExactArgs(2),
	ValidArgs: []string{"alias", "command"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add is called!")

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
