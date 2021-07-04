package cmd

import (
	"fmt"
	"hail/internal/hailconfig"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list is used to print all the alias and commands",
	Run: func(cmd *cobra.Command, args []string) {
		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()
		err := hc.Parse()
		if err != nil {
			fmt.Println("error in list : ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
