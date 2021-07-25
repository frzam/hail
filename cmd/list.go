package cmd

import (
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
		checkError("error in parsing", err)

		err = hc.List()
		checkError("error in list", err)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
