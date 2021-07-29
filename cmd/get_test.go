package cmd

import (
	"bytes"
	"hail/internal/hailconfig"
	"io/ioutil"
	"testing"
)

func Test_RunGet(t *testing.T) {
	o := NewGetOptions()
	hc, _ := hailconfig.NewHailconfig(hailconfig.WithMockHailconfigLoader(""))
	b := bytes.NewBufferString("")

	// Test that alias is getting returned properly.
	for alias, wantCmd := range hailconfig.TestScripts {
		o.Alias = alias
		o.Run(hc, b)

		out, _ := ioutil.ReadAll(b)

		got := string(out)
		if got != wantCmd+"\n" {
			t.Errorf("want: '%s' while got: '%s'", wantCmd, got)
		}
	}

	// Test validations
	// No alias is found i.e. Empty alias
	o.Alias = ""
	gotErr := o.Run(hc, b).Error()
	wantErr := "error in validation: no alias is found"
	if gotErr != wantErr {
		t.Errorf("want error: '%q', got error: '%q'", wantErr, gotErr)
	}

	// No alias is present
	o.Alias = "my-alias"
	gotErr = o.Run(hc, b).Error()
	wantErr = "alias is not present: no command is found with 'my-alias' alias"
	assertErr(wantErr, gotErr, t)
}

func Test_CmdGet(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdGet(hailconfig.WithMockHailconfigLoader(""), b)
	cmd.SetArgs([]string{"pv", "klogs"})
	_ = cmd.Execute()

	// Test validations

	// multiple alias is present
	cmd.SetArgs([]string{"pv", "klogs"})
	gotErr := cmd.Execute().Error()
	wantErr := "Error: error in validation: more than one alias is present"
	assertErr(wantErr, gotErr, t)
}
func assertErr(wantErr, gotErr string, t *testing.T) {
	t.Helper()
	if gotErr != wantErr {
		t.Errorf("want error: '%q', got error: '%q'", wantErr, gotErr)
	}
}
