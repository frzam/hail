package cmd

import (
	"fmt"
	"hail/internal/fuzzy"
	"hail/internal/hailconfig"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hail",
	Short: "hail is a cross-platform script management tool",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
	defer hc.Close()

	err := hc.Parse()
	checkError("error in parsing", err)

	alias, err := findFuzzyAlias(hc)
	checkError("error while finding alias", err)
	if alias == "" || !hc.IsPresent(alias) {
		checkError("alias is not present", fmt.Errorf("no command is found with '%s' alias", alias))
	}

	command, err := hc.Get(alias)
	checkError("error in get", err)
	fmt.Fprintln(os.Stdout, command)
}

func findFuzzyAlias(hc *hailconfig.Hailconfig) (string, error) {
	ig := fuzzy.NewIterativeGet(hc)

	return ig.FindAlias()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkError(msg string, err error) {
	if err != nil {
		red := color.New(color.FgRed, color.Bold).SprintFunc()
		fmt.Printf("%s: %s: %v\n", red("Error"), msg, err)
		os.Exit(2)
	}
}

func success(msg string) {
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	fmt.Printf("%s: %s\n", green("Success"), msg)
}
