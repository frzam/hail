package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/hailconfig"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func NewCmdRoot() *cobra.Command {
	return &cobra.Command{
		Use:   "hail",
		Short: "hail is a cross-platform script management tool",
		Run:   run,
	}

}

func run(cmd *cobra.Command, args []string) {
	hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
	defer hc.Close()

	err := hc.Parse()
	cmdutil.CheckErr("error in parsing", err)

	alias, err := cmdutil.FindFuzzyAlias(hc)
	cmdutil.CheckErr("error while finding alias", err)

	if alias == "" || !hc.IsPresent(alias) {
		cmdutil.CheckErr("alias is not present", fmt.Errorf("no command is found with '%s' alias", alias))
	}

	command, err := hc.Get(alias)
	cmdutil.CheckErr("error in get", err)
	fmt.Fprintln(os.Stdout, command)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd = NewCmdRoot()
	rootCmd.AddCommand(NewCmdGet(hailconfig.DefaultLoader, os.Stdout))
	rootCmd.AddCommand(NewCmdAdd(hailconfig.DefaultLoader, os.Stdout))
	rootCmd.AddCommand(NewCmdCopy(hailconfig.DefaultLoader, os.Stdout))
	rootCmd.AddCommand(NewCmdDelete(hailconfig.DefaultLoader, os.Stdout))
	rootCmd.AddCommand(NewCmdEdit(hailconfig.DefaultLoader, os.Stdout))
}
