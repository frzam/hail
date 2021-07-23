package cmd

import (
	"fmt"
	"hail/internal/hailconfig"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete/rm [alias]",
	Short:   "delete removes command from hail basis alias",
	Aliases: []string{"rm"},
	RunE: func(cmd *cobra.Command, args []string) error {
		err := validateArgs(args)
		if err != nil {
			fmt.Println("validation error:", err)
			os.Exit(2)
		}

		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err = hc.Parse()
		if err != nil {
			fmt.Println("error in delete:", err)
			os.Exit(2)
		}

		err = hc.Delete(args[0])
		if err != nil {
			fmt.Println("error in delete:", err)
			os.Exit(2)
		}
		fmt.Printf("command with alias '%s' has been deleted\n", args[0])
		return hc.Save()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
