package cmd

import (
	"fmt"
	"hail/internal/hailconfig"

	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:     "copy [old-alias] [new-alias]",
	Short:   "copy/cp is  used to copy one command/script to a new alias.",
	Aliases: []string{"cp"},
	Run: func(cmd *cobra.Command, args []string) {
		oldAlias, err := cmd.Flags().GetString("oldAlias")
		checkError("error in parsing flag", err)

		newAlias, err := cmd.Flags().GetString("newAlias")
		checkError("error in parsing flag", err)

		if oldAlias == "" || newAlias == "" {
			err = validateCopyOrMove(args)
			checkError("error in validation", err)
			oldAlias = args[0]
			newAlias = args[1]
		}

		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err = hc.Parse()
		checkError("error in parsing", err)

		err = hc.Copy(oldAlias, newAlias)
		checkError("error in copy", err)

		err = hc.Save()
		checkError("error in save", err)

		fmt.Printf("command with alias '%s' has been copied to alias '%s'\n", oldAlias, newAlias)

	},
}

func validateCopyOrMove(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("invalid number of arguments. expected 2, recieved %d", len(args))
	}
	return nil
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.Flags().StringP("oldAlias", "o", "", "old alias to be copied from")
	copyCmd.Flags().StringP("newAlias", "n", "", "new alias to be copied to")
}
