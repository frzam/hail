package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/editor"
	"hail/internal/hailconfig"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

// EditOptions contains all the necessory fields that are needed for edit cmd.
type EditOptions struct {
	Alias           string
	Command         string
	DescriptionFlag string
}

// NewEditOption is a contructor that returns an empty *EditOptions
func NewEditOption() *EditOptions {
	return &EditOptions{}
}

// NewCmdEdit creates cobra command and it gets the alias then basis alias gets the command
// then it opens the command in default editor and updates it after the command.
func NewCmdEdit(loader hailconfig.Loader, w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "edit [alias]",
		Short:   "It is used to edit previously added command or script in text editor.",
		Example: cmdutil.EditExample,
		Run: func(cmd *cobra.Command, args []string) {
			o := NewEditOption()

			hc, err := hailconfig.NewHailconfig(loader)
			cmdutil.CheckErr("error in new hailconfig", err)
			defer hc.Close()

			// Get alias from flag or from args
			o.Alias, _ = getAlias(cmd, args)
			cmdutil.CheckErr("error in validation", err)

			if o.Alias == "" && len(args) == 0 {
				o.Alias, err = cmdutil.FindFuzzyAlias(hc)
				cmdutil.CheckErr("error while finding alias", err)
			}

			// Get description
			o.DescriptionFlag, _ = cmd.Flags().GetString("description")

			if !hc.IsPresent(o.Alias) {
				cmdutil.CheckErr("alias is not present", fmt.Errorf("no command is found with '%s' alias", o.Alias))
			}
			o.Command, err = hc.Get(o.Alias)
			cmdutil.CheckErr("error in get", err)

			e := editor.NewDefaultEditor([]string{})
			bCommand, _, err := e.LaunchTempFile("hail", true, strings.NewReader(o.Command))
			cmdutil.CheckErr("error in launching temp file", err)

			err = hc.Update(o.Alias, string(bCommand), o.DescriptionFlag)
			cmdutil.CheckErr("error in update", err)

			err = hc.Save()
			cmdutil.CheckErr("error in save", err)

			cmdutil.Success(fmt.Sprintf("command with alias '%s' has been edited\n", o.Alias))
		},
	}
	cmd.Flags().StringP("alias", "a", "", "alias for the command/script")
	cmd.Flags().StringP("description", "d", "", "description of the command")
	return cmd
}
