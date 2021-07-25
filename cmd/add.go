package cmd

import (
	"fmt"
	"hail/internal/hailconfig"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:       "add [alias] [command]",
	Short:     "add is used to add a new command in collection",
	ValidArgs: []string{"alias", "command"},
	Run: func(cmd *cobra.Command, args []string) {
		// validating flags and args.
		alias, err := cmd.Flags().GetString("alias")
		command := ""
		if err != nil || (alias == "" && len(args) < 2) {
			checkError("error in validation", fmt.Errorf("no alias or command is present"))
		}
		des, _ := cmd.Flags().GetString("description")
		if alias == "" && len(args) > 1 {
			alias = args[0]
			command = strings.Join(args[1:], "")
		} else if alias != "" && len(args) > 0 {
			command = strings.Join(args[0:], "")
		} else {
			checkError("error in validation", fmt.Errorf("no alias or command is present"))
		}

		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err = hc.Parse()
		checkError("error in parsing", err)

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
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("alias", "a", "", "alias for the command")
	addCmd.Flags().StringP("description", "d", "", "describe the command")
}
