package cmd

import (
	"fmt"
	"hail/cmd/cmdutil"
	"hail/internal/hailconfig"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd with package level scope.
var rootCmd *cobra.Command

// RootOptions contains all fields needed to run hail cmd.
type RootOptions struct {
	Alias   string
	Command string
}

// NewRootOption is an empty constructor that returns *RootOptions
func NewRootOption() *RootOptions {
	return &RootOptions{}
}

// NewCmdRoot returns a cmd that when executed will get output.
func NewCmdRoot(loader hailconfig.Loader, w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hail",
		Short: "hail is a cross-platform script management tool",
		Run: func(cmd *cobra.Command, args []string) {
			o := NewRootOption()

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
			cmdutil.CheckErr("error in run", o.Run(hc, w))
		},
	}
	return cmd
}

// Run validates the alias to see if is not empty and is present in hailconfig otherwise returns error.
// If validation is ok, then it retrives command by calling Get method. If not error is found then
// it prints command. It is called with os.Stdout writer.
func (o *RootOptions) Run(hc *hailconfig.Hailconfig, w io.Writer) error {
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

// Execute runs the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// init func adds all the subcommands into our rootCmd
func init() {
	rootCmd = NewCmdRoot(hailconfig.DefaultLoader, os.Stdout)
	rootCmd.AddCommand(NewCmdGet(hailconfig.DefaultLoader, os.Stdout))
	rootCmd.AddCommand(NewCmdAdd(hailconfig.DefaultLoader, os.Stdout))
	rootCmd.AddCommand(NewCmdCopy(hailconfig.DefaultLoader, os.Stdout))
	rootCmd.AddCommand(NewCmdDelete(hailconfig.DefaultLoader, os.Stdout))
	rootCmd.AddCommand(NewCmdEdit(hailconfig.DefaultLoader, os.Stdout))
	rootCmd.AddCommand(NewCmdInit(hailconfig.DefaultLoader, os.Stdout))
	rootCmd.AddCommand(NewCmdList(hailconfig.DefaultLoader, os.Stdout))
	rootCmd.AddCommand(NewCmdMove(hailconfig.DefaultLoader, os.Stdout))
	rootCmd.AddCommand(NewCmdUpdate(hailconfig.DefaultLoader, os.Stdout))
	rootCmd.AddCommand(NewCmdVersion(hailconfig.DefaultLoader, os.Stdout))
}
