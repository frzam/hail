package cmd

import (
	"fmt"
	"hail/internal/hailconfig"
	"os"

	"github.com/spf13/cobra"
)

const updateExample = `  # Update Command with 'delete-pods' alias
  hail update delete-pods 'kubectl delete pod $(kubectl get pods | grep Completed | awk '{print $1}')'`

var updateCmd = &cobra.Command{
	Use:     "update [alias] [command]",
	Short:   "updates already present command.",
	Example: updateExample,
	Args:    cobra.ExactArgs(2),
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
		fmt.Printf("command with alias '%s' has been updated\n", args[0])
		return hc.Save()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
