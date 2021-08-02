package cmd

import (
	"hail/cmd/cmdutil"
	"hail/internal/hailconfig"
	"io"

	"github.com/spf13/cobra"
)

// ListOptions is an empty struct since list need no options.
type ListOptions struct{}

// NewListOption is an empty constructor.
func NewListOption() *ListOptions {
	return &ListOptions{}
}

// NewCmdList creates a list cmd, which when called list out all the alias,
// with command and description in tabular form.
func NewCmdList(loader hailconfig.Loader, w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "list/ls prints all the alias and commands",
		Example: cmdutil.ListExample,
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, args []string) {

			o := NewListOption()

			hc, err := hailconfig.NewHailconfig(loader)
			cmdutil.CheckErr("error in new hailconfig", err)
			defer hc.Close()

			cmdutil.CheckErr("error in run", o.Run(hc, w))
		},
	}
	return cmd
}

// Run calls List method and prints out the table containing all the alias,
// command and descriptions.
func (o *ListOptions) Run(hc *hailconfig.Hailconfig, _ io.Writer) error {
	return hc.List()
}
