package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/hailconfig"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

// UpdateOptions contain fields for update cmd.
type UpdateOptions struct {
	Alias       string
	Command     string
	Description string
}

// NewUpdateOption returns an empty *UpdateOptions
func NewUpdateOption() *UpdateOptions {
	return &UpdateOptions{}
}

// NewCmdUpdate creates a update command. It validates the args and check for alias
// and command.
func NewCmdUpdate(loader hailconfig.Loader, w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update [alias] [command]",
		Short:   "updates already present command",
		Example: cmdutil.UpdateExample,
		Run: func(cmd *cobra.Command, args []string) {
			o := NewUpdateOption()
			var err error
			o.Alias, err = cmd.Flags().GetString("alias")
			o.Description, _ = cmd.Flags().GetString("description")

			if err != nil || (o.Alias == "" && len(args) < 2) {
				cmdutil.CheckErr("error in validation", fmt.Errorf("no alias or command is present"))
			}
			if o.Alias == "" && len(args) > 1 {
				o.Alias = args[0]
				o.Command = strings.Join(args[1:], "")
			} else if o.Alias != "" && len(args) > 0 {
				o.Command = strings.Join(args[0:], "")
			} else {
				cmdutil.CheckErr("error in validation", fmt.Errorf("no alias or command is present"))
			}

			hc, err := hailconfig.NewHailconfig(loader)
			cmdutil.CheckErr("error in new hailconfig", err)

			cmdutil.CheckErr("error in run", o.Run(hc, w))
			cmdutil.Success(fmt.Sprintf("command with alias '%s' has been updated\n", o.Alias))
		},
	}
	cmd.Flags().StringP("alias", "a", "", "alias for the command")
	cmd.Flags().StringP("description", "d", "", "descrition of the command")
	return cmd
}

// Run updates command and description in hailconfig basis alias.
func (o *UpdateOptions) Run(hc *hailconfig.Hailconfig, w io.Writer) error {
	err := hc.Update(o.Alias, o.Command, o.Description)
	if err != nil {
		return err
	}
	return hc.Save()
}
