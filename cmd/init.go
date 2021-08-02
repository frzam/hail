package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/hailconfig"
	"io"

	"github.com/spf13/cobra"
)

// InitOptions contains field needed to run init cmd.
type InitOptions struct {
	Title   string
	CfgFile string
}

// NewInitOptions return an empty *InitOptions
func NewInitOptions() *InitOptions {
	return &InitOptions{}
}

// NewCmdInit creates a cobra cmd, which when called will create a new .hailconfig file.
// If file is already present then it will throw error.
func NewCmdInit(loader hailconfig.Loader, w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [title]",
		Short: "init initializes an empty .hailconfig file with title as provided",
		Run: func(cmd *cobra.Command, args []string) {
			o := NewInitOptions()
			if len(args) < 1 {
				o.Title = "default"
			} else {
				o.Title = args[0]
			}
			cmdutil.CheckErr("error in init", o.Run(&hailconfig.Hailconfig{}, w))
			cmdutil.Success(fmt.Sprintf("Initialized a file '%s'", o.CfgFile))

		},
	}
	return cmd
}

// Run calls the Init func that validates the location of .hailconfig, and if not present
// then creates a file and returns it. Otherwise returns an error.
func (o *InitOptions) Run(_ *hailconfig.Hailconfig, _ io.Writer) error {
	var err error
	o.CfgFile, err = hailconfig.Init(o.Title)
	return err
}
