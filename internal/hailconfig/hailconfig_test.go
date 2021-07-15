package hailconfig

import (
	"testing"
)

var scripts = map[string]string{
	"oc-login":  "oc login -u farzam -p",
	"kube-logs": "kubectl logs -f --tail=00 ",
	"pv":        "oc get pv",
}

func TestAdd(t *testing.T) {

	hc := new(Hailconfig).WithLoader(WithMockHailconfigLoader(""))
	hc.Parse()

	for k, v := range scripts {
		hc.Add(k, v)
	}
	for key, want := range scripts {
		got := hc.Scripts[key].Command
		if got != want {
			assertAddError(t, got, want, key)
		}
	}
}
func assertAddError(t *testing.T, got, want, key string) {
	t.Helper()
	t.Errorf("got: %q want: %q given: %q", got, want, key)
}

func TestUpdate(t *testing.T) {
	hc := new(Hailconfig).WithLoader(WithMockHailconfigLoader(""))
	hc.Parse()
	for k, v := range scripts {
		hc.Add(k, v)
	}
	want := "kubectl logs -f --tail=100"
	hc.Update("kube-logs", want)
	got := hc.Scripts["kube-logs"].Command
	if got != want {
		t.Errorf("got: %q want: %q", got, want)
	}
}

func TestUpdateEmpty(t *testing.T) {
	hc := new(Hailconfig).WithLoader(WithMockHailconfigLoader(""))
	hc.Parse()
	err := hc.Update("kube-logs", "kubectl logs -f --tail=100")
	if err == nil {
		t.Errorf("expected error")
	}
	if hc.Scripts != nil {
		t.Errorf("should have been a nil map")
	}
}