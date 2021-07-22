package cmd

import (
	"errors"
	"fmt"
	"hail/internal/hailconfig"
	"os"

	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:     "copy [alias] [new-alias]",
	Short:   "It is used to copy one command/script to a new alias.",
	Aliases: []string{"cp"},
	RunE: func(cmd *cobra.Command, args []string) error {
		err := validateCopy(args)
		if err != nil {
			fmt.Printf("error in validate copy: %v\n", err)
			os.Exit(2)
		}
		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err = hc.Parse()
		if err != nil {
			fmt.Printf("error:%v\n", err)
			os.Exit(2)
		}
		//err = hc.Copy(args[0], args[1])
		if err != nil {
			fmt.Println("error:%v\n", err)
			os.Exit(2)
		}
		return nil
	},
}

func validateCopy(args []string) error {
	if len(args) != 2 {
		return errors.New(fmt.Sprintf("invalid number of arguments. expected 2, recieved %d", len(args)))
	}
	return nil
}

func init() {
	rootCmd.AddCommand(copyCmd)
}
