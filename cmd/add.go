package cmd

import (
	"fmt"
	"hail/internal/editor"
	"hail/internal/hailconfig"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:       "add [alias] [command]",
	Short:     "add is used to add a new command in collection",
	ValidArgs: []string{"alias", "command"},
	Run: func(cmd *cobra.Command, args []string) {
		// Get alias either from -a flag or from args.
		alias, err := getAlias(cmd, args)
		checkError("error in validation", err)

		// Get description from -d flag
		des, _ := cmd.Flags().GetString("description")

		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()

		err = hc.Parse()
		checkError("error in parsing", err)

		if hc.IsPresent(alias) {
			checkError("error in validation", fmt.Errorf("alias already present"))
		}

		// Get Command from arguments or from editor
		command, err := getCommand(cmd, args)
		checkError("error in getting cmd", err)

		if alias == "" || command == "" {
			checkError("error in validation", fmt.Errorf("no alias or command is present"))
		}

		hc.Add(alias, command, des)
		err = hc.Save()
		checkError("error in save", err)

		success(fmt.Sprintf("command with alias '%s' has been added\n", alias))
	},
}

func getAlias(cmd *cobra.Command, args []string) (string, error) {
	alias, err := cmd.Flags().GetString("alias")
	if err != nil {
		return "", errors.Wrap(err, "error while get alias from flag")
	}
	if alias == "" && len(args) > 0 {
		alias = args[0]
	}
	if alias == "" {
		return alias, fmt.Errorf("no alias is found")
	}
	return alias, nil
}

func getCommand(cmd *cobra.Command, args []string) (string, error) {
	command := ""
	// Read command from a file.
	f, _ := cmd.Flags().GetString("file")
	if f != "" {
		bCommand, err := os.ReadFile(f)
		if err != nil {
			return "", errors.Wrap(err, "error while reading file")
		}
		return string(bCommand), nil
	}
	// Read file from args.
	alias, _ := cmd.Flags().GetString("alias")
	if len(args) == 1 && alias == "" || len(args) == 0 && alias != "" {
		e := editor.NewDefaultEditor([]string{})
		bCommand, _, err := e.LaunchTempFile("hail", false, os.Stdout)
		checkError("error in launching temp file", err)
		command = string(bCommand)
	} else if alias != "" && len(args) > 0 {
		command = strings.Join(args[0:], "")
	} else if alias == "" && len(args) > 1 {
		command = strings.Join(args[1:], "")
	}
	if command == "" {
		return command, fmt.Errorf("no command is found")
	}
	return command, nil
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("alias", "a", "", "alias for the command")
	addCmd.Flags().StringP("description", "d", "", "description of the command")
	addCmd.Flags().StringP("file", "f", "", "path of the file that needs to be read as command")
}
