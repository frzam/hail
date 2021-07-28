package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/hailconfig"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete [alias]",
	Short:   "delete/rm removes command from hail basis alias",
	Aliases: []string{"rm"},
	Run: func(cmd *cobra.Command, args []string) {
		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err := hc.Parse()
		cmdutil.CheckErr("error in parse", err)

		alias := ""
		if len(args) == 0 {
			alias, err = cmdutil.FindFuzzyAlias(hc)
			cmdutil.CheckErr("error while finding alias", err)
		}
		if alias == "" {
			alias, err = cmd.Flags().GetString("alias")
			if err != nil || alias == "" {
				cmdutil.CheckErr("error in flag parsing", err)

				err = cmdutil.ValidateArgss(args)
				cmdutil.CheckErr("error in validation", err)
				alias = args[0]
			}
		}

		err = hc.Delete(alias)
		cmdutil.CheckErr("error in delete", err)
		err = hc.Save()
		cmdutil.CheckErr("error in save", err)

		cmdutil.Success(fmt.Sprintf("command with alias '%s' has been deleted\n", alias))
	},
}

func init() {
	NewCmdRoot().AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("alias", "a", "", "alias for the command")
}
