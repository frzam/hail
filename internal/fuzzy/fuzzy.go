package fuzzy

import (
	"fmt"
	"hail/internal/hailconfig"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	"github.com/pkg/errors"
)

// fuzzyScript contains alias and command as type string
type fuzzyScript struct {
	alias   string
	command string
}

// IterativeGet is struct that contains a slice of fuzzyScripts.
type IterativeGet struct {
	FuzzyScripts []fuzzyScript
}

// NewIterativeGet is a constructor that returns an IterativeGet type which constains
// slice of fuzzyScripts from *Hailconfig.Scripts.
func NewIterativeGet(hc *hailconfig.Hailconfig) IterativeGet {
	var fuzzyScripts []fuzzyScript
	for a, s := range hc.Scripts {
		fuzzyScripts = append(fuzzyScripts, fuzzyScript{a, s.Command})
	}
	return IterativeGet{fuzzyScripts}
}

// FindAlias returns ana alias basis fuzzy search or error.
// It launches a preview windows where user can search for the alias and
// select it.
func (ig IterativeGet) FindAlias() (string, error) {
	igx, err := fuzzyfinder.Find(ig.FuzzyScripts,
		func(i int) string {
			return ig.FuzzyScripts[i].alias
		},
		// Launch PreviewWindow
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
