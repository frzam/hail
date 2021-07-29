package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/hailconfig"
	"strings"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update [alias] [command]",
	Short:   "updates already present command.",
	Example: cmdutil.UpdateExample,
	Run: func(cmd *cobra.Command, args []string) {
		alias, err := cmd.Flags().GetString("alias")
		des, _ := cmd.Flags().GetString("description")
		command := ""
		if err != nil || (alias == "" && len(args) < 2) {
			cmdutil.CheckErr("error in validation", fmt.Errorf("no alias or command is present"))
		}
		if alias == "" && len(args) > 1 {
			alias = args[0]
			command = strings.Join(args[1:], "")
		} else if alias != "" && len(args) > 0 {
			command = strings.Join(args[0:], "")
		} else {
			cmdutil.CheckErr("error in validation", fmt.Errorf("no alias or command is present"))
		}

		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err = hc.Parse()
		cmdutil.CheckErr("error in parse", err)

		err = hc.Update(alias, command, des)
		cmdutil.CheckErr("error in update", err)

		err = hc.Save()
		cmdutil.CheckErr("error in save", err)

		cmdutil.Success(fmt.Sprintf("command with alias '%s' has been updated\n", alias))
	},
}

func init() {
	NewCmdRoot().AddCommand(updateCmd)
	updateCmd.Flags().StringP("alias", "a", "", "alias for the command")
	updateCmd.Flags().StringP("description", "d", "", "descrition of the command")
}
