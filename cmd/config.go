package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/hailconfig"
	"io"

	"github.com/spf13/cobra"
)

// ConfigOptions contains fields to perform configuration related operations
type ConfigOptions struct {
	List  bool
	Name  string
	Value string
}

// NewconfigOption creates an empty *ConfigOptions
func NewConfigOption() *ConfigOptions {
	return &ConfigOptions{}
}

// NewCmdConfig creates cmd to perform config operations
func NewCmdConfig(loader hailconfig.Loader, w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "config [flags]",
		Short:   "it is used to list or update configurations",
		Example: cmdutil.ConfigExample,
		Run: func(cmd *cobra.Command, args []string) {
			o := NewConfigOption()

			o.List, _ = cmd.Flags().GetBool("list")
			o.Name, _ = cmd.Flags().GetString("name")
			o.Value, _ = cmd.Flags().GetString("value")

			hc, err := hailconfig.NewHailconfig(loader)
			cmdutil.CheckErr("error in new hail config", err)

			cmdutil.CheckErr("error in run ", o.Run(hc, w))
			if !o.List {
				cmdutil.Success(fmt.Sprintf("config '%s' has been updated\n", o.Name))
			}
		},
	}
	cmd.Flags().BoolP("list", "l", false, "list all configurations.")
	cmd.Flags().StringP("name", "n", "", "name of the configuration to set")
	cmd.Flags().StringP("value", "v", "", "value of the configuration to set")
	return cmd
}

// Run either prints all the configuration properties or updates configs
func (o *ConfigOptions) Run(hc *hailconfig.Hailconfig, w io.Writer) error {
	if o.List {
		hc.ListConfigProperties(w)
		return nil
	}
	return hc.UpdateConfigProperties(o.Name, o.Value)
}
