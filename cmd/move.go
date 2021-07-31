package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/hailconfig"
	"io"

	"github.com/spf13/cobra"
)

// MoveOptions contains all fields needed to execute move cmd
type MoveOptions struct {
	OldAlias     string
	NewAlias     string
	OldAliasFlag string
	NewAliasFlag string
}

// NewMoveOption is a constructor that returns an empty *MoveOption.
func NewMoveOption() *MoveOptions {
	return &MoveOptions{}
}

// NewCmdMove creates move cmd and assigns flags to it as well. It calls Run method,
// which performs the actual move and returns error if any.
func NewCmdMove(loader hailconfig.Loader, w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "move [old-alias] [new-alias]",
		Short:   "move/mv used to move command with old alias to new alias.",
		Aliases: []string{"mv"},
		Run: func(cmd *cobra.Command, args []string) {
			o := NewMoveOption()

			var err error

			o.OldAliasFlag, err = cmd.Flags().GetString("oldAlias")
			cmdutil.CheckErr("error in parsing flag", err)

			o.NewAliasFlag, err = cmd.Flags().GetString("newAlias")
			cmdutil.CheckErr("error in parsing flag", err)

			if o.OldAliasFlag == "" || o.NewAliasFlag == "" {
				err = validateCopyOrMove(args)
				cmdutil.CheckErr("error in validation", err)
				o.OldAlias = args[0]
				o.NewAlias = args[1]
			}

			hc, err := hailconfig.NewHailconfig(loader)
			cmdutil.CheckErr("error in new hailconfig", err)
			defer hc.Close()

			cmdutil.CheckErr("error in run", o.Run(hc, w))

			cmdutil.Success(fmt.Sprintf("command with alias '%s' has been moved to alias '%s'\n", o.OldAlias, o.NewAlias))
		},
	}
	cmd.Flags().StringP("oldAlias", "o", "", "old alias to be copied from")
	cmd.Flags().StringP("newAlias", "n", "", "new alias to be copied to")
	return cmd
}

// Run method performs the move cmd or returns an error.
func (o *MoveOptions) Run(hc *hailconfig.Hailconfig, _ io.Writer) error {
	if o.OldAliasFlag != "" {
		o.OldAlias = o.OldAliasFlag
	}
	if o.NewAliasFlag != "" {
		o.NewAlias = o.NewAliasFlag
	}
	err := hc.Move(o.OldAlias, o.NewAlias)
	if err != nil {
		return err
	}
	return hc.Save()
}
