package cmdutil

import "testing"

func Test_ValidateArgss(t *testing.T) {
	// Test one.
	gotErr := ValidateArgss([]string{"pv"})
	if gotErr != nil {
		t.Errorf("want: nil error, got: '%q'", gotErr)
	}
	// Test less than one.
	gotErr = ValidateArgss([]string{})
	wantErr := "no alias is present"
	assertErr(wantErr, gotErr.Error(), t)

	// Test more than one.
	gotErr = ValidateArgss([]string{"pv", "klog"})
	wantErr = "more than one alias is present"
	assertErr(wantErr, gotErr.Error(), t)
}

func assertErr(wantErr, gotErr string, t *testing.T) {
	t.Helper()
	if gotErr != wantErr {
		t.Errorf("want: '%s' got: '%s'\n", wantErr, gotErr)
	}
}
