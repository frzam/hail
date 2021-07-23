package cmd

import (
	"fmt"
	"hail/internal/hailconfig"
	"os"

	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:     "copy/cp [old-alias] [new-alias]",
	Short:   "It is used to copy one command/script to a new alias.",
	Aliases: []string{"cp"},
	Run: func(cmd *cobra.Command, args []string) {
		err := validateCopyOrMove(args)
		if err != nil {
			fmt.Printf("error in validate copy: %v\n", err)
			os.Exit(2)
		}
		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err = hc.Parse()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(2)
		}
		err = hc.Copy(args[0], args[1])
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(2)
		}
		if err = hc.Save(); err != nil {
			fmt.Printf("error in save : %v\n", err)
			os.Exit(2)
		}
		fmt.Printf("command with alias '%s' has been copied to alias '%s'\n", args[0], args[1])

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
}
