package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "hail",
	Short: "hail is a cross-platform script management tool",
	RunE:  run,
}

func run(cmd *cobra.Command, args []string) error {
	//return runGet(cmd, args)
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
