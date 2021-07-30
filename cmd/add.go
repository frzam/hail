package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/editor"
	"hail/internal/hailconfig"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// AddOptions contains all the fields that are needed for add cmd.
type AddOptions struct {
	Alias           string
	Command         string
	AliasFlag       string
	DescriptionFlag string
	FileFlag        string
}

// NewAddOptions returns an empty AddOptions.
func NewAddOptions() *AddOptions {
	return &AddOptions{}
}

func NewCmdAdd(loader hailconfig.Loader, w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:       "add [alias] [command]",
		Short:     "add is used to add a new command in collection",
		ValidArgs: []string{"alias", "command"},
		Run: func(cmd *cobra.Command, args []string) {
			o := NewAddOptions()

			var err error

			// Get alias either from -a flag or from args.
			o.Alias, err = getAlias(cmd, args)
			cmdutil.CheckErr("error in validation", err)

			// Get description from -d flag
			o.DescriptionFlag, _ = cmd.Flags().GetString("description")

			// Get file from -d flag
			o.FileFlag, _ = cmd.Flags().GetString("file")

			hc, err := hailconfig.NewHailconfig(loader)
			cmdutil.CheckErr("error in new hailconfig", err)
			defer hc.Close()

			if hc.IsPresent(o.Alias) {
				cmdutil.CheckErr("error in validation", fmt.Errorf("alias already present"))
			}
			// Get Command from arguments or from editor
			o.Command, err = getCommand(cmd, args)
			cmdutil.CheckErr("error in getting cmd", err)

			if o.Alias == "" || o.Command == "" {
				cmdutil.CheckErr("error in validation", fmt.Errorf("no alias or command is present"))
			}
			cmdutil.CheckErr("error in run", o.Run(hc, w))
			cmdutil.Success(fmt.Sprintf("command with alias '%s' has been added\n", o.Alias))
		},
	}
	cmd.Flags().StringP("alias", "a", "", "alias for the command")
	cmd.Flags().StringP("description", "d", "", "description of the command")
	cmd.Flags().StringP("file", "f", "", "path of the file that needs to be read as command")
	return cmd
}

// Run adds command and discription with alias, it then calls Save() which saves trucates
// data in hailconfig file.
func (o *AddOptions) Run(hc *hailconfig.Hailconfig, w io.Writer) error {
	hc.Add(o.Alias, o.Command, o.DescriptionFlag)
	return hc.Save()
}

// getAlias finds alias from flag or args. If no alias is found then it
// returns error.
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

// getCommand finds the command either from args or file or from editor, if no command is found
// then it returns error.
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
		cmdutil.CheckErr("error in launching temp file", err)
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
