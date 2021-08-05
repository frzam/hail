package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/editor"
	"hail/internal/hailconfig"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// RunOptions contains fields to perform Run cmd.
type RunOptions struct {
	Alias   string
	Command string
	Path    string
}

func NewRunOption() *RunOptions {
	return &RunOptions{}
}

func NewCmdRun(loader hailconfig.Loader, w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run [alias]",
		Short: "it is used to directly run a command from alias",
		Run: func(cmd *cobra.Command, args []string) {
			o := NewRunOption()
			var err error

			o.Path, err = editor.CreateTempFile("hail", false, os.Stdout)
			cmdutil.CheckErr("error while creating temp file", err)

			hc, err := hailconfig.NewHailconfig(loader)
			cmdutil.CheckErr("error in new hailconfig", err)
			defer hc.Close()

			if len(args) == 0 {
				o.Alias, err = cmdutil.FindFuzzyAlias(hc)
				cmdutil.CheckErr("error while finding alias", err)
			}

			if o.Alias == "" && len(args) > 0 {
				err = cmdutil.ValidateArgss(args)
				cmdutil.CheckErr("error in validation", err)
				o.Alias = args[0]
			}
			if o.Alias == "" {
				cmdutil.CheckErr("", fmt.Errorf("error in validation: no alias is found"))
			}
			if !hc.IsPresent(o.Alias) {
				cmdutil.CheckErr("", fmt.Errorf("alias is not present: no command is found with '%s' alias", o.Alias))
			}
			cmdutil.CheckErr("error in run", o.Run(hc, w))
		},
	}
	return cmd
}

func (o *RunOptions) Run(hc *hailconfig.Hailconfig, w io.Writer) error {
	o.Command, _ = hc.Get(o.Alias)

	e := editor.NewDefaultEditor([]string{})

	output, err := e.RunScript(o.Path, string(o.Command), hc)
	if err != nil {
		return err
	}
	cmdutil.Success(string(output))
	return nil
}
