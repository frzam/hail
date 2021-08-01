package cmd

import (
	"fmt"
	"hail/internal/hailconfig"
	"io"

	"github.com/spf13/cobra"
)

var (
	version = "v0.1.8"
)

// NewCmdVersion creates version cmd, it prints out the latest version when runs.
func NewCmdVersion(_ hailconfig.Loader, w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "version prints the current version of hail",

		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Print version basis the latest release.
			fmt.Fprintln(w, version)
		},
	}
	return cmd
}
