package cmd

import (
	"bytes"
	"hail/internal/hailconfig"
	"testing"
)

func Test_RunCopy(t *testing.T) {
	o := NewCopyOption()
	hc, _ := hailconfig.NewHailconfig(hailconfig.WithMockHailconfigLoader(""))
	b := bytes.NewBufferString("")

	// Test is that copy operation is running smoothly
	o.OldAlias = "debug-pod"
	o.NewAlias = "dp"
	err := o.Run(hc, b)
	if err != nil {
		t.Errorf("got error '%q'", err)
	}
	if _, ok := hc.Scripts[o.NewAlias]; !ok {
		t.Errorf("new alias '%s' is not found", o.NewAlias)
	}

	// Test error that old alias is not present
	o.OldAlias = ""
	err = o.Run(hc, b)
	if err == nil {
		t.Errorf("expected alias not found error")
	}

	// Test error that new alias is already present
	o.OldAlias = "debug-pod"
	o.NewAlias = "dp"
	err = o.Run(hc, b)
	if err == nil {
		t.Error("expected error")
	}
}
