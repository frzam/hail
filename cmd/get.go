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
	Short: "get retrieves command basis the alias.",
	Run:   runGet,
}

func runGet(cmd *cobra.Command, args []string) {
	command := get(cmd, args)
	fmt.Fprintln(os.Stdout, command)
}

func get(cmd *cobra.Command, args []string) string {
	alias := ""

	hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
	defer hc.Close()

	err := hc.Parse()

	checkError("error in parsing", err)
	if len(args) == 0 {
		alias, err = findFuzzyAlias(hc)
		checkError("error while finding alias", err)
	}
	if alias == "" {
		err = validateArgs(args)
		checkError("error in validation", err)
		alias = args[0]
	}

	if !hc.IsPresent(alias) {
		checkError("alias is not present", fmt.Errorf("no command is found with '%s' alias", alias))
	}
	command, err := hc.Get(alias)
	checkError("error in get", err)
	return command
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
