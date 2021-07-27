package fuzzy

import (
	"fmt"
	"hail/internal/hailconfig"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	"github.com/pkg/errors"
)

type fuzzyScript struct {
	alias   string
	command string
}

type IterativeGet struct {
	FuzzyScripts []fuzzyScript
}

func NewIterativeGet(hc *hailconfig.Hailconfig) IterativeGet {
	var fuzzyScripts []fuzzyScript
	for a, s := range hc.Scripts {
		fuzzyScripts = append(fuzzyScripts, fuzzyScript{a, s.Command})
	}
	return IterativeGet{fuzzyScripts}
}

func (ig IterativeGet) FindAlias() (string, error) {
	igx, err := fuzzyfinder.Find(ig.FuzzyScripts,
		func(i int) string {
			return ig.FuzzyScripts[i].alias
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintln(ig.FuzzyScripts[i].command)
		}))
	if err != nil {
		if err.Error() == "abort" {
			return "", fmt.Errorf("no alias is found")
		}
		return "", errors.Wrap(err, "error while iterative get")
	}
	return ig.FuzzyScripts[igx].alias, nil
}
