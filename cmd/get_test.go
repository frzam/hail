package cmd

import (
	"bytes"
	"hail/internal/hailconfig"
	"io/ioutil"
	"os"
	"os/exec"
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
	if os.Getenv("ERROR") == "1" {
		// TEST error with multiple aliases.
		b := bytes.NewBufferString("")
		cmd := NewCmdGet(hailconfig.WithMockHailconfigLoader(""), b)
		cmd.SetArgs([]string{"pv", "klogs"})
		_ = cmd.Execute()
	}

	c := exec.Command(os.Args[0], "-test.run=Test_CmdGet")
	c.Env = append(c.Env, "ERROR=1")
	stdout, _ := c.StdoutPipe()
	if err := c.Start(); err != nil {
		t.Fatal(err)
	}
	gotBytes, _ := ioutil.ReadAll(stdout)
	gotErr := string(gotBytes)
	wantErr := "Error: error in validation: more than one alias is present\n"
	assertErr(wantErr, gotErr, t)
}
func assertErr(wantErr, gotErr string, t *testing.T) {
	t.Helper()
	if gotErr != wantErr {
		t.Errorf("want error: '%q', got error: '%q'", wantErr, gotErr)
	}
}
