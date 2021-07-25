package cmd

import (
	"fmt"
	"hail/internal/editor"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [alias]",
	Short: "It is used to edit previously added command or script in text editor.",
	Run: func(cmd *cobra.Command, args []string) {
		e := editor.NewDefaultEditor([]string{})
		bytes, _, err := e.LaunchTempFile("hail")
		if err != nil {
			fmt.Println("error: ", err)
		}
		fmt.Println(string(bytes))

	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
