package cmdutil

import (
	"errors"
	"fmt"
	"hail/internal/fuzzy"
	"hail/internal/hailconfig"
	"os"

	"github.com/fatih/color"
)

// ValidateArgss is used to validate number of args.
func ValidateArgss(args []string) error {
	if len(args) < 1 {
		return errors.New("no alias is present")
	}
	if len(args) > 1 {
		return errors.New("more than one alias is present")
	}
	return nil
}

// CheckErr is util func that is used to check error passed, if err is not nil
// then it logs error with 'Error" as red.
func CheckErr(msg string, err error) {
	if err != nil {
		red := color.New(color.FgRed, color.Bold).SprintFunc()
		fmt.Printf("%s: %s: %v\n", red("Error"), msg, err)
		os.Exit(2)
	}
}

// Success is called when the command works properly.
func Success(msg string) {
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	fmt.Printf("%s: %s", green("Success"), msg)
}

// FindFuzzyAlias returns a windows that is used to select an alias,
// in case of any problem, it returns an error.
func FindFuzzyAlias(hc *hailconfig.Hailconfig) (string, error) {
	return fuzzy.NewIterativeGet(hc).FindAlias()
}
