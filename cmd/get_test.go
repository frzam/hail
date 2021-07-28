package cmd

import (
	"bytes"
	"hail/internal/hailconfig"
	"io/ioutil"
	"testing"
)

var scripts = map[string]string{
	"oc-login":  "oc login -u farzam -p",
	"kube-logs": "kubectl logs -f --tail=00 ",
	"pv":        "oc get pv",
}

func Test_Run(t *testing.T) {
	o := NewGetOptions()
	o.Alias = "pv"
	hc := NewHailConfigDummy()
	hc.Parse()
	b := bytes.NewBufferString("")
	o.Run(hc, b)
	out, _ := ioutil.ReadAll(b)
	want := "oc get pv\n"
	got := string(out)
	if got != want {
		t.Errorf("want: '%s' while got: '%s'", want, got)
	}
}

func NewHailConfigDummy() *hailconfig.Hailconfig {
	hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.WithMockHailconfigLoader(""))
	for k, v := range scripts {
		hc.Add(k, v, "")
	}
	return hc
}
