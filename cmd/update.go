package cmd

import (
	"fmt"
	"hail/internal/hailconfig"
	"os"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [alias] [command]",
	Short: "it updates already present command",
	RunE: func(cmd *cobra.Command, args []string) error {
		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()
		err := hc.Parse()
		if err != nil {
			fmt.Println("error while parsing: ", err)
			os.Exit(2)
		}
		err = hc.Update(args[0], args[1])
		if err != nil {
			fmt.Println("error while update: ", err)
			os.Exit(2)
		}
		return hc.Save()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
