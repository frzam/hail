package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/hailconfig"
	"io"

	"github.com/spf13/cobra"
)

type GetOptions struct {
	Alias   string
	Command string
}

func NewGetOptions() *GetOptions {
	return &GetOptions{}
}

func NewCmdGet(loader hailconfig.Loader, w io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "get [alias]",
		Short: "get retrieves command basis the alias.",
		Run: func(cmd *cobra.Command, args []string) {
			o := NewGetOptions()

			hc, err := hailconfig.NewHailconfig(loader)
			cmdutil.CheckErr("error in new hailconfig", err)

			if len(args) == 0 {
				o.Alias, err = cmdutil.FindFuzzyAlias(hc)
				cmdutil.CheckErr("error while finding alias", err)
			}

			if o.Alias == "" && len(args) > 0 {
				err = cmdutil.ValidateArgss(args)
				cmdutil.CheckErr("error in validation", err)
				o.Alias = args[0]
			}

			o.Run(hc, w)
		},
	}
}

func (o *GetOptions) Run(hc *hailconfig.Hailconfig, w io.Writer) {
	if o.Alias == "" {
		cmdutil.CheckErr("error in validation", fmt.Errorf("no alias is found"))
	}

	if !hc.IsPresent(o.Alias) {
		cmdutil.CheckErr("alias is not present", fmt.Errorf("no command is found with '%s' alias", o.Alias))
	}
	var err error
	o.Command, err = hc.Get(o.Alias)
	cmdutil.CheckErr("error in get", err)

	fmt.Fprintf(w, "%s\n", o.Command)
}
