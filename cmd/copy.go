package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/hailconfig"
	"io"

	"github.com/spf13/cobra"
)

// CopyOptions constain all the fields that are needed for copy cmd.
type CopyOptions struct {
	OldAlias     string
	NewAlias     string
	OldAliasFlag string
	NewAliasFlag string
}

// NewCopyOption return an empty *CopyOption
func NewCopyOption() *CopyOptions {
	return &CopyOptions{}
}

// NewCmdCopy returns a command that create copy command.
func NewCmdCopy(loader hailconfig.Loader, w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "copy [old-alias] [new-alias]",
		Short:   "copy/cp is  used to copy one command/script to a new alias.",
		Example: cmdutil.CopyExample,
		Aliases: []string{"cp"},
		Run: func(cmd *cobra.Command, args []string) {

			o := NewCopyOption()

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
			cmdutil.Success(fmt.Sprintf("command with alias '%s' has been copied to alias '%s'\n", o.OldAlias, o.NewAlias))

		},
	}
	cmd.Flags().StringP("oldAlias", "o", "", "old alias to be copied from")
	cmd.Flags().StringP("newAlias", "n", "", "new alias to be copied to")
	return cmd
}

// Run copies command with old alias to new alias.
func (o *CopyOptions) Run(hc *hailconfig.Hailconfig, w io.Writer) error {
	if o.OldAliasFlag != "" {
		o.OldAlias = o.OldAliasFlag
	}
	if o.NewAliasFlag != "" {
		o.NewAlias = o.NewAliasFlag
	}
	err := hc.Copy(o.OldAlias, o.NewAlias)
	if err != nil {
		return err
	}
	return hc.Save()
}

// validataeCopyOrMove validates if two arguments are present
func validateCopyOrMove(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("invalid number of arguments. expected 2, received %d", len(args))
	}
	return nil
}
