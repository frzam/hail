package cmd

import (
	"hail/internal/hailconfig"

	"github.com/spf13/cobra"
)

type ConfigOptions struct {
	List  bool
	Name  string
	Value string
}

func NewConfigOption() *ConfigOptions {
	return &ConfigOptions{}
}

func NewCmdConfig(loader hailconfig.Loader) *cobra.Command {
	cmd := &cobra.Command{
		//o := NewConfigOption()

	}

	return cmd
}
