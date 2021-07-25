package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

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
