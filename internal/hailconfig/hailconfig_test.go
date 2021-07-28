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

	hc, _ := NewHailconfig(WithMockHailconfigLoader(""))
	for k, v := range scripts {
		hc.Add(k, v, "")
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

	hc := NewHailConfigDummy()

	want := "kubectl logs -f --tail=100"
	hc.Update("kube-logs", want, "")
	got := hc.Scripts["kube-logs"].Command
	if got != want {
		t.Errorf("got: %q want: %q", got, want)
	}
}

func TestUpdateEmpty(t *testing.T) {
	hc := new(Hailconfig).WithLoader(WithMockHailconfigLoader(""))
	hc.Parse()
	err := hc.Update("kube-logs", "kubectl logs -f --tail=100", "")
	if err == nil {
		t.Errorf("expected error")
	}
	if hc.Scripts != nil {
		t.Errorf("should have been a nil map")
	}
}

func TestIsPresent(t *testing.T) {
	hc := new(Hailconfig).WithLoader(WithMockHailconfigLoader(""))
	hc.Parse()

	assertIsPresent(t, hc.IsPresent("kube-logs"), false)
	for k, v := range scripts {
		hc.Add(k, v, "")
	}
	assertIsPresent(t, hc.IsPresent("oc-login"), true)
	assertIsPresent(t, hc.IsPresent("pv"), true)
}

func assertIsPresent(t *testing.T, got, want bool) {
	if got != want {
		t.Errorf("got: %t want: %t", got, want)
	}
}

func TestDelete(t *testing.T) {
	hc := NewHailConfigDummy()

	err := hc.Delete("pv")
	if err != nil {
		t.Errorf("recieved error: %v, but error should have been nil", err)
	}
	if hc.IsPresent("pv") {
		t.Errorf("alias should not be present")
	}
	// aliasNotFoundErr
	err = hc.Delete("k-logs")
	if err == nil {
		t.Errorf("should get aliasNotFoundErr")
	}
}
func TestCopy(t *testing.T) {
	hc := NewHailConfigDummy()

	oldAlias := "oc-login"
	newAlias := "login"

	// Basic Copy
	err := hc.Copy(oldAlias, newAlias)
	if err != nil {
		t.Errorf("no error as expected but found err: %v\n", err)
	}
	got := hc.Scripts[newAlias].Command
	want := hc.Scripts[oldAlias].Command
	assertGotWant(t, got, want)

	// When old alias is not present
	oldAlias = "dlogin"
	err = hc.Copy(oldAlias, newAlias)
	if err == nil {
		t.Errorf("expected error but no error is returned")
	}
	got = err.Error()
	want = "old alias is not present"
	assertGotWant(t, got, want)

	// When new alias is already present
	oldAlias = "pv"
	newAlias = "oc-login"
	err = hc.Copy(oldAlias, newAlias)
	if err == nil {
		t.Errorf("expected error but not error is returned")
	}
	got = err.Error()
	want = "new alias is already present"
	assertGotWant(t, got, want)
}

func TestMove(t *testing.T) {

	hc := NewHailConfigDummy()

	// Basic Move
	oldAlias := "pv"
	newAlias := "pvc"
	want, _ := hc.Get(oldAlias)
	err := hc.Move(oldAlias, newAlias)
	if err != nil {
		t.Errorf("no error as expected but found err: %v\n", err)
	}
	got := hc.Scripts[newAlias].Command
	assertGotWant(t, got, want)
	if hc.IsPresent(oldAlias) {
		t.Errorf("old alias is still present after move")
	}
	// When old alias is not present
	oldAlias = "dlogin"
	newAlias = "dlogin-new"
	err = hc.Move(oldAlias, newAlias)
	if err == nil {
		t.Errorf("expected error but no error is returned")
	}
	got = err.Error()
	want = "old alias is not present"
	assertGotWant(t, got, want)

	if hc.IsPresent(newAlias) {
		t.Errorf("new alias is present")
	}
	// When new alias is alread present
	oldAlias = "oc-login"
	newAlias = "kube-logs"

	err = hc.Move(oldAlias, newAlias)
	if err == nil {
		t.Errorf("expected error but no error is returned")
	}
	got = err.Error()
	want = "new alias is already present"
	assertGotWant(t, got, want)

}

func NewHailConfigDummy() *Hailconfig {
	hc := new(Hailconfig).WithLoader(WithMockHailconfigLoader(""))
	hc.Parse()
	for k, v := range scripts {
		hc.Add(k, v, "")
	}
	return hc
}

func assertGotWant(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' while want '%s'\n", got, want)
	}
}
