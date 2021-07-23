package cmd

import (
	"fmt"
	"hail/internal/hailconfig"
	"os"

	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	Use:     "move/mv [old-alias] [new-alias]",
	Short:   "It is used to move command with old alias to new alias.",
	Aliases: []string{"mv"},
	Run: func(cmd *cobra.Command, args []string) {
		err := validateCopyOrMove(args)
		if err != nil {
			fmt.Printf("error in validate move: %v\n", err)
			os.Exit(2)
		}
		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err = hc.Parse()
		if err != nil {
			fmt.Printf("error in parse: %v\n", err)
			os.Exit(2)
		}
		err = hc.Move(args[0], args[1])
		if err != nil {
			fmt.Printf("error in move: %v\n", err)
			os.Exit(2)
		}
		if err = hc.Save(); err != nil {
			fmt.Printf("error in save: %v\n", err)
			os.Exit(2)
		}
		fmt.Printf("command with alias '%s' has been moved to alias '%s'\n", args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)
}
