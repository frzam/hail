package cmd

import (
	"fmt"
	"hail/internal/editor"
	"hail/internal/hailconfig"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [alias]",
	Short: "It is used to edit previously added command or script in text editor.",
	Run: func(cmd *cobra.Command, args []string) {
		alias, err := cmd.Flags().GetString("alias")
		// TODO: validations
		if err != nil || (alias == "" && len(args) < 2) {
			checkError("error in validation", fmt.Errorf("no alias or command is present"))
		}
		command := ""
		if len(args) < 2 {
			e := editor.NewDefaultEditor([]string{})
			bCommand, _, err := e.LaunchTempFile("hail")
			checkError("error in creating temp file: ", err)
			command = string(bCommand)
		}
		if alias == "" {
			alias = args[0]
		}
		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()
		// TODO: Update for edit, currently logic is like add
		err = hc.Parse()
		checkError("error in parsing", err)
		des, _ := cmd.Flags().GetString("description")
		if hc.IsPresent(alias) {
			checkError("error in validation", fmt.Errorf("alias already present"))
		}
		hc.Add(alias, command, des)
		err = hc.Save()
		checkError("error in save", err)

		success(fmt.Sprintf("command with alias '%s' has been added\n", alias))
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringP("alias", "a", "", "alias for the command/script")
}
