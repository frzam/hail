package cmd

import (
	"fmt"
	"hail/internal/hailconfig"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete [alias]",
	Short:   "delete/rm removes command from hail basis alias",
	Aliases: []string{"rm"},
	Run: func(cmd *cobra.Command, args []string) {
		alias, err := cmd.Flags().GetString("alias")
		if err != nil || alias == "" {
			checkError("error in flag parsing", err)

			err = validateArgs(args)
			checkError("error in validation", err)
			alias = args[0]
		}

		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err = hc.Parse()
		checkError("error in delete", err)

		err = hc.Delete(alias)
		checkError("error in delete", err)
		err = hc.Save()
		checkError("error in save", err)

		success(fmt.Sprintf("command with alias '%s' has been deleted\n", alias))
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("alias", "a", "", "alias for the command")
}
