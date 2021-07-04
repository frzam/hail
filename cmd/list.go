package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list is used to print all the alias and commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list is called!")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
