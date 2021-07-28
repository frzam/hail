package cmdutil

import (
	"errors"
	"fmt"
	"hail/internal/fuzzy"
	"hail/internal/hailconfig"
	"os"

	"github.com/fatih/color"
)

func ValidateArgss(args []string) error {
	if len(args) < 1 {
		return errors.New("no alias is present")
	}
	if len(args) > 1 {
		return errors.New("more than one alias is present")
	}
	return nil
}

func CheckErr(msg string, err error) {
	if err != nil {
		red := color.New(color.FgRed, color.Bold).SprintFunc()
		fmt.Printf("%s: %s: %v\n", red("Error"), msg, err)
		os.Exit(2)
	}
}

func Success(msg string) {
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	fmt.Printf("%s: %s", green("cmdutil.Success"), msg)
}

func FindFuzzyAlias(hc *hailconfig.Hailconfig) (string, error) {
	return fuzzy.NewIterativeGet(hc).FindAlias()
}
