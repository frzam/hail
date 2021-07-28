package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/hailconfig"
	"io"

	"github.com/spf13/cobra"
)

// GetOptions contains all the fields that are are used in Run method.
type GetOptions struct {
	Alias   string
	Command string
}

// NewGetOptions returns an empty GetOptions.
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

			cmdutil.CheckErr("error in run", o.Run(hc, w))
		},
	}
}

// Run validates the alias to see if is not empty and is present in hailconfig otherwise returns error.
// If validation is ok, then it retrives command by calling Get method. If not error is found then
// it prints command. It is called with os.Stdout writer.
func (o *GetOptions) Run(hc *hailconfig.Hailconfig, w io.Writer) error {
	// validating alias.
	if o.Alias == "" {
		return fmt.Errorf("error in validation: no alias is found")
	}
	if !hc.IsPresent(o.Alias) {
		return fmt.Errorf("alias is not present: no command is found with '%s' alias", o.Alias)
	}
	// call get method with alias.
	var err error
	o.Command, err = hc.Get(o.Alias)
	if err != nil {
		return fmt.Errorf("error in get: %q", err)
	}
	fmt.Fprintf(w, "%s\n", o.Command)
	return nil
}
