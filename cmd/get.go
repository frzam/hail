package cmd

import (
	"fmt"
	"hail/internal/hailconfig"
	"io"
	"os"

	"github.com/spf13/cobra"
)

type GetOptions struct {
	Alias   string
	Command string
}

func NewGetOptions() *GetOptions {
	return &GetOptions{}
}

func NewCmdGet(loader hailconfig.Loader) *cobra.Command {
	return &cobra.Command{
		Use:   "get [alias]",
		Short: "get retrieves command basis the alias.",
		Run: func(cmd *cobra.Command, args []string) {
			o := NewGetOptions()

			hc, err := hailconfig.NewHailconfig(loader)
			checkError("error in new hailconfig", err)

			if len(args) == 0 {
				o.Alias, err = findFuzzyAlias(hc)
				checkError("error while finding alias", err)
			}

			if o.Alias == "" && len(args) > 0 {
				err = validateArgs(args)
				checkError("error in validation", err)
				o.Alias = args[0]
			}

			o.Run(hc, os.Stdout)
		},
	}
}

func (o *GetOptions) Run(hc *hailconfig.Hailconfig, w io.Writer) {
	if o.Alias == "" {
		checkError("error in validation", fmt.Errorf("no alias is found"))
	}

	if !hc.IsPresent(o.Alias) {
		checkError("alias is not present", fmt.Errorf("no command is found with '%s' alias", o.Alias))
	}
	var err error
	o.Command, err = hc.Get(o.Alias)
	checkError("error in get", err)

	fmt.Fprintln(w, o.Command)
}
