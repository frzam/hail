package cmd

import (
	"fmt"
	"hail/internal/hailconfig"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:       "add [alias] [command]",
	Short:     "add is used to add a new command in collection",
	Args:      cobra.ExactArgs(2),
	ValidArgs: []string{"alias", "command"},
	RunE: func(cmd *cobra.Command, args []string) error {
		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err := hc.Parse()
		if err != nil {
			fmt.Println("error while parsing: ", err)
			os.Exit(1)
		}
		hc.Add(args[0], args[1])
		fmt.Printf("command with alias '%s' has been added\n", args[0])
		return hc.Save()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
