package cmd

import (
	"fmt"
	"hail/internal/hailconfig"
	"os"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [alias]",
	Short: "It retrieves command basics the alias.",
	RunE: func(cmd *cobra.Command, args []string) error {
		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err := hc.Parse()
		if err != nil {
			fmt.Println("error in get: ", err)
			os.Exit(2)
		}
		if len(args) < 1 || args[0] == "" || !hc.IsPresent(args[0]) {
			fmt.Println("invalid alias :", args)
			os.Exit(2)
		}
		command, err := hc.Get(args[0])
		fmt.Println(command)
		return err
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func validateGet(cmd *cobra.Command, args []string) error {
	return nil
}
