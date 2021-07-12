package cmd

import (
	"errors"
	"fmt"
	"hail/internal/hailconfig"
	"os"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [alias]",
	Short: "It retrieves command basics the alias.",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := validateGet(args)
		if err != nil {
			fmt.Println("validation error:", err)
			os.Exit(2)
		}
		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err = hc.Parse()
		if err != nil {
			fmt.Println("error in get:", err)
			os.Exit(2)
		}
		if !hc.IsPresent(args[0]) {
			fmt.Printf("err: no command is found with this '%s' alias\n", args[0])
			os.Exit(2)
		}
		command, err := hc.Get(args[0])
		fmt.Println(command)
		return err
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func validateGet(args []string) error {
	if len(args) < 1 {
		return errors.New("no alias is present")
	}
	if len(args) > 1 {
		return errors.New("more than one alias is present")
	}
	return nil
}
