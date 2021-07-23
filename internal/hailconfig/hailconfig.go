package hailconfig

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

var aliasNotFoundErr = errors.New("alias is not found")

// script is each individual script. It has only one field Command.
type script struct {
	Command string `toml:"command"`
}

// Hailconfig contains three field loader, f of type ReadWriteResetCloser interface and config
// Working on hailconfig file is done using value of type Hailconfig
type Hailconfig struct {
	loader Loader
	f      ReadWriteResetCloser
	config
}

// config represents the complete hailconfig file. It contains Title and
// map Scripts.
type config struct {
	Title   string
	Scripts map[string]script `toml:"scripts"`
}

func (hc *Hailconfig) WithLoader(l Loader) *Hailconfig {
	hc.loader = l
	return hc
}

// Close is used to close a haliconfig file.
// If file reference is nil that means file is already closed.
func (hc *Hailconfig) Close() error {
	if hc.f == nil {
		return nil
	}
	return hc.f.Close()
}

// Add is used to add a new script to Scripts map.
// It takes alias and command as input and creates a type of script and adds
// to hc.Scripts map.
func (hc *Hailconfig) Add(alias, command string) {
	sc := script{
		Command: command,
	}
	if hc.Scripts == nil {
		hc.Scripts = make(map[string]script)
	}

	hc.Scripts[alias] = sc
}

// Save writes Hailconfig data into .hailconfig file.
// It resets the file so that new data can be written over the existing file,
// it returns error in case of failed to reset or encoding failure.
func (hc *Hailconfig) Save() error {
	err := hc.f.Reset()
	if err != nil {
		return errors.Wrap(err, "failed to reset")
	}
	return toml.NewEncoder(hc.f).Encode(&hc.config)
}

// List is used to print all the aliases and commands in Scripts map.
func (hc *Hailconfig) List() error {
	for alias, script := range hc.Scripts {
		fmt.Fprintf(os.Stdout, "%s\t\t%s\n", alias, script.Command)
	}
	return nil
}

// Update is used to update a command for already present alias.
// If the alias is not present which is to be updated then returns error.
func (hc *Hailconfig) Update(alias, command string) error {
	if found := hc.IsPresent(alias); !found {
		return aliasNotFoundErr
	}
	hc.Add(alias, command)
	return nil
}

// IsPresent checks if the alias is prsent in Scripts map.
func (hc *Hailconfig) IsPresent(alias string) bool {
	_, found := hc.Scripts[alias]
	return found
}

// Delete removes a command basis the alias provided.
// It returns alias not found error when alias is not present
func (hc *Hailconfig) Delete(alias string) error {
	if !hc.IsPresent(alias) {
		return aliasNotFoundErr
	}
	delete(hc.Scripts, alias)
	return nil
}

func (hc *Hailconfig) Copy(oldAlias, newAlias string) error {
	if !hc.IsPresent(oldAlias) {
		return errors.New("old alias is not present.")
	}
	if hc.IsPresent(newAlias) {
		return errors.New("new alias is already present")
	}
	hc.Add(newAlias, hc.Scripts[oldAlias].Command)
	return nil
}

// Get returns command and error based on the alias name.
// It checks for the alias in Scipts map. Before calling Get we should check if
// the alias is present in map or not.
func (hc *Hailconfig) Get(alias string) (string, error) {
	return hc.Scripts[alias].Command, nil

}

func (hc *Hailconfig) Parse() error {
	files, err := hc.loader.Load()
	if err != nil {
		return errors.Wrap(err, "failed to load")
	}
	f := files[0]
	hc.f = f
	_, err = toml.DecodeReader(hc.f, &hc.config)
	if err != nil {
		return errors.Wrap(err, "failed to decode")
	}
	return nil
}
