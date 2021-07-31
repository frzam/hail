package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/editor"
	"os"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [alias]",
	Short: "it is used to directly run a command from alias",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := editor.CreateTempFile("hail", false, os.Stdout)
		cmdutil.CheckErr("error while creating temp file", err)
		//TODO:
		command := "" // get(cmd, args)
		e := editor.NewDefaultEditor([]string{})
		output, err := e.RunScript(path, command)
		cmdutil.CheckErr("error in run script", err)

		fmt.Fprintln(os.Stdout, "output", string(output))
	},
}
