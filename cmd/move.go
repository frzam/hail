package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/hailconfig"

	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	Use:     "move [old-alias] [new-alias]",
	Short:   "move/mv used to move command with old alias to new alias.",
	Aliases: []string{"mv"},
	Run: func(cmd *cobra.Command, args []string) {
		oldAlias, err := cmd.Flags().GetString("oldAlias")
		cmdutil.CheckErr("error in parsing flag", err)

		newAlias, err := cmd.Flags().GetString("newAlias")
		cmdutil.CheckErr("error in parsing flag", err)

		if oldAlias == "" || newAlias == "" {
			err = validateCopyOrMove(args)
			cmdutil.CheckErr("error in validation", err)
			oldAlias = args[0]
			newAlias = args[1]
		}

		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err = hc.Parse()
		cmdutil.CheckErr("error in parse", err)

		err = hc.Move(oldAlias, newAlias)
		cmdutil.CheckErr("error in move", err)

		err = hc.Save()
		cmdutil.CheckErr("error in save", err)

		cmdutil.Success(fmt.Sprintf("command with alias '%s' has been moved to alias '%s'\n", oldAlias, newAlias))
	},
}

func init() {
	NewCmdRoot().AddCommand(moveCmd)
	moveCmd.Flags().StringP("oldAlias", "o", "", "old alias to be copied from")
	moveCmd.Flags().StringP("newAlias", "n", "", "new alias to be copied to")
}
