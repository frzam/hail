package cmd

import (
	"fmt"
	"hail/internal/editor"
	"hail/internal/hailconfig"
	"strings"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [alias]",
	Short: "It is used to edit previously added command or script in text editor.",
	Run: func(cmd *cobra.Command, args []string) {
		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err := hc.Parse()
		checkError("error in parsing", err)

		alias := ""
		if len(args) == 0 {
			alias, err = findFuzzyAlias(hc)
			checkError("error while finding alias", err)
		}
		if alias == "" {
			// Get alias from flag or from args
			alias, err = getAlias(cmd, args)
			checkError("error in validation", err)
		}

		// Get description
		des, _ := cmd.Flags().GetString("description")

		if !hc.IsPresent(alias) {
			checkError("alias is not present", fmt.Errorf("no command is found with '%s' alias", args[0]))
		}
		command, err := hc.Get(alias)
		checkError("error in get", err)

		e := editor.NewDefaultEditor([]string{})
		bCommand, _, err := e.LaunchTempFile("hail", true, strings.NewReader(command))
		checkError("error in launching temp file", err)

		err = hc.Update(alias, string(bCommand), des)
		checkError("error in update", err)

		err = hc.Save()
		checkError("error in save", err)

		success(fmt.Sprintf("command with alias '%s' has been edited\n", alias))
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringP("alias", "a", "", "alias for the command/script")
	editCmd.Flags().StringP("description", "d", "", "description of the command")
}
