package cmd

import (
	"fmt"
	"hail/internal/hailconfig"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const updateExample = `  # Update Command with 'delete-pods' alias
  hail update delete-pods 'kubectl delete pod $(kubectl get pods | grep Completed | awk '{print $1}')'`

var updateCmd = &cobra.Command{
	Use:     "update [alias] [command]",
	Short:   "updates already present command.",
	Example: updateExample,
	Run: func(cmd *cobra.Command, args []string) {
		alias, err := cmd.Flags().GetString("alias")
		des, _ := cmd.Flags().GetString("description")
		command := ""
		if err != nil || (alias == "" && len(args) < 2) {
			fmt.Println("error: no alias or command is present")
			os.Exit(2)
		}
		if alias == "" && len(args) > 1 {
			alias = args[0]
			command = strings.Join(args[1:], "")
		} else if alias != "" && len(args) > 0 {
			command = strings.Join(args[0:], "")
		} else {
			fmt.Println("error: no alias or command is present")
			os.Exit(2)
		}

		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err = hc.Parse()
		checkError("error in parse", err)

		err = hc.Update(alias, command, des)
		checkError("error in update", err)

		err = hc.Save()
		checkError("error in save", err)

		fmt.Printf("command with alias '%s' has been updated\n", alias)

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("alias", "a", "", "alias for the command")
	updateCmd.Flags().StringP("description", "d", "", "descrition of the command")
}
