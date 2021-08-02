package cmd

import (
	"hail/internal/hailconfig"
	"os"
	"testing"
)

func Test_RunMove(t *testing.T) {
	o := NewMoveOption()
	hc, _ := hailconfig.NewHailconfig(hailconfig.WithMockHailconfigLoader(""))
	// Test general move
	o.OldAlias = "pv"
	o.NewAlias = "persistence-volume"
	err := o.Run(hc, os.Stdout)
	if err != nil {
		t.Errorf("want nil error, got %q error", err)
	}
	if !hc.IsPresent(o.NewAlias) {
		t.Errorf("new alias '%q' is not present", o.NewAlias)
	}
	if hc.IsPresent(o.OldAlias) {
		t.Errorf("old alias '%q' is still present", o.OldAlias)
	}
	// Test old alias is not present
	o.OldAlias = "abcd"
	err = o.Run(hc, os.Stdout)
	if err == nil {
		t.Errorf("expected error, got nil error")
	}
	// Test new alias is not present
	o.OldAlias = "pv"
	o.NewAlias = "debug-pod"
	err = o.Run(hc, os.Stdout)
	if err == nil {
		t.Errorf("got nil error, expected error")
	}
}
