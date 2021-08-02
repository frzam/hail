package cmd

import (
	"hail/internal/hailconfig"
	"os"
	"testing"
)

func Test_RunDelete(t *testing.T) {
	o := NewDeleteOptions()
	hc, _ := hailconfig.NewHailconfig(hailconfig.WithMockHailconfigLoader(""))
	o.Alias = "debug-pod"
	err := o.Run(hc, os.Stdout)
	if err != nil {
		t.Errorf("did not expect error, but got %q", err)
	}
	if hc.IsPresent(o.Alias) {
		t.Errorf("alias is not deleted %s", o.Alias)
	}
	o.Alias = "abc"
	err = o.Run(hc, os.Stdout)
	if err == nil {
		t.Errorf("expected error but got nil error")
	}
}

func Test_CmdDelete(t *testing.T) {
	cmd := NewCmdDelete(hailconfig.WithMockHailconfigLoader(""), os.Stdout)
	cmd.SetArgs([]string{"pv"})
	err := cmd.Execute()
	if err != nil {
		t.Errorf("expected nil error, got %q error", err)
	}
	cmd.Flags().Set("alias", "debug-pod")
	err = cmd.Execute()
	if err != nil {
		t.Errorf("expected nil error, got %q error", err)
	}
}
