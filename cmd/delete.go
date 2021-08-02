package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/hailconfig"
	"io"

	"github.com/spf13/cobra"
)

// DeleteOptions  contains fields to perform delete cmd.
type DeleteOptions struct {
	Alias   string
	Command string
}

// NewDeleteOptions returns an empty *DeleteOptions
func NewDeleteOptions() *DeleteOptions {
	return &DeleteOptions{}
}

// NewCmdDelete creates a cobra.Command and flags associated with it. It does validation
// and calls Run method if validations are fine.
func NewCmdDelete(loader hailconfig.Loader, w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete [alias]",
		Short:   "delete/rm removes command from hail basis alias",
		Example: cmdutil.DeleteExample,
		Aliases: []string{"rm"},
		Run: func(cmd *cobra.Command, args []string) {
			o := NewDeleteOptions()

			hc, err := hailconfig.NewHailconfig(loader)
			cmdutil.CheckErr("error in new hailconfig", err)
			defer hc.Close()

			o.Alias, err = cmd.Flags().GetString("alias")
			cmdutil.CheckErr("error in flag parsing", err)

			if o.Alias == "" && len(args) == 0 {
				o.Alias, err = cmdutil.FindFuzzyAlias(hc)
				cmdutil.CheckErr("error while finding alias", err)
			}
			if o.Alias == "" {
				err = cmdutil.ValidateArgss(args)
				cmdutil.CheckErr("error in validation", err)
				o.Alias = args[0]
			}

			cmdutil.CheckErr("error in run", o.Run(hc, w))
			cmdutil.Success(fmt.Sprintf("command with alias '%s' has been deleted\n", o.Alias))
		},
	}
	cmd.Flags().StringP("alias", "a", "", "alias for the command")
	return cmd
}

// Run method deletes the command and alias and saves.
func (o *DeleteOptions) Run(hc *hailconfig.Hailconfig, w io.Writer) error {
	err := hc.Delete(o.Alias)
	if err != nil {
		return err
	}
	return hc.Save()
}
