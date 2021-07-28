package cmd

import (
	"hail/cmd/cmdutil"
	"hail/internal/hailconfig"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "list/ls prints all the alias and commands",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err := hc.Parse()
		cmdutil.CheckErr("error in parsing", err)

		err = hc.List()
		cmdutil.CheckErr("error in list", err)
	},
}

func init() {
	NewCmdRoot().AddCommand(listCmd)
}
