package fuzzy

import (
	"hail/internal/hailconfig"
	"testing"
)

func Test_NewIterativeGet(t *testing.T) {
	hc, _ := hailconfig.NewHailconfig(hailconfig.WithMockHailconfigLoader(""))
	defer hc.Close()

	ig := NewIterativeGet(hc)
	if len(ig.FuzzyScripts) != len(hailconfig.TestScripts) {
		t.Errorf("want length: %d, got length %d\n", len(hailconfig.TestScripts), len(ig.FuzzyScripts))
	}
	for _, s := range ig.FuzzyScripts {
		if _, ok := hailconfig.TestScripts[s.alias]; !ok {
			t.Errorf("alias is not present : '%s'", s.alias)
		}
	}
}
