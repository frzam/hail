package cmd

import (
	"hail/internal/hailconfig"
	"os"
	"testing"
)

func Test_RunUpdate(t *testing.T) {
	o := NewUpdateOption()
	hc, _ := hailconfig.NewHailconfig(hailconfig.WithMockHailconfigLoader(""))
	o.Alias = "was-bin"
	o.Description = "was bin"
	o.Command = "cd /opt/IBM/BPM/bin"
	err := o.Run(hc, os.Stdout)
	if err != nil {
		t.Errorf("expected nil error, got %q error", err)
	}
	gotDes := hc.Scripts[o.Alias].Description
	wantDes := o.Description
	if gotDes != wantDes {
		t.Errorf("want description: %s, got descript: %s", wantDes, gotDes)
	}
	gotCmd := hc.Scripts[o.Alias].Command
	wantCmd := o.Command
	if gotCmd != wantCmd {
		t.Errorf("want command: %s, got command: %s", wantCmd, gotCmd)
	}

	// Test no alias is present
	o.Alias = "abcd"
	err = o.Run(hc, os.Stdout)
	if err == nil {
		t.Errorf("want error, got nil error")
	}
}
