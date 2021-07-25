package cmd

import (
	"errors"
	"fmt"
	"hail/internal/hailconfig"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [alias]",
	Short: "get retrieves command basis the alias.",
	RunE:  runGet,
}

func runGet(cmd *cobra.Command, args []string) error {
	err := validateArgs(args)
	checkError("error in validation", err)

	hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
	defer hc.Close()

	err = hc.Parse()
	checkError("error in parsing", err)

	if !hc.IsPresent(args[0]) {
		checkError("alias is not present", fmt.Errorf("no command is found with '%s' alias", args[0]))
	}
	command, err := hc.Get(args[0])
	fmt.Println(command)
	return err
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func validateArgs(args []string) error {
	if len(args) < 1 {
		return errors.New("no alias is present")
	}
	if len(args) > 1 {
		return errors.New("more than one alias is present")
	}
	return nil
}
