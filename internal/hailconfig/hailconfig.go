package hailconfig

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/pkg/errors"
)

// Script is each individual script. It has only one field Command.
type Script struct {
	Command     string `toml:"command"`
	Description string `toml:"description"`
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
	Scripts map[string]Script `toml:"scripts"`
}

// WithLoader takes in a loader and assigns into *Hailconfig type.
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

// NewHailconfig is a constuctor that take Loader and returns *Hailconfig and error.
// It loads into *Hailconfig basis the loader provided and calls parse which parses the file
// into config field of *Hailconfig and returns it.
// It checks if the loader provided is *MockHailconfigLoader, if yes then it adds,
// data from TestScripts as well and then returns it. It is used for testing.
func NewHailconfig(l Loader) (*Hailconfig, error) {
	hc := new(Hailconfig).WithLoader(Loader(l))
	err := hc.Parse()
	if err != nil {
		return nil, errors.Wrap(err, "error in parsing")
	}
	// Checking to see if is called from test case func.
	switch l.(type) {
	case *MockHailconfigLoader:
		// if yes, then add TestScripts data at the same time.
		for k, v := range TestScripts {
			hc.Add(k, v, "")
		}
	default:
		return hc, nil
	}
	return hc, nil
}

// Add is used to add a new script to Scripts map.
// It takes alias and command as input and creates a type of script and adds
// to hc.Scripts map.
func (hc *Hailconfig) Add(alias, command, des string) {
	var sc Script
	if des != "" {
		sc = Script{
			Command:     command,
			Description: des,
		}
	} else {
		sc = Script{
			Command: command,
		}
	}

	if hc.Scripts == nil {
		hc.Scripts = make(map[string]Script)
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
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Alias", "Command", "Description"})
	t.SetColumnConfigs([]table.ColumnConfig{{}})
	for alias, script := range hc.Scripts {
		t.AppendRow([]interface{}{alias, script.Command, script.Description})
		t.AppendSeparator()
	}
	t.Render()
	return nil
}

// Update is used to update a command for already present alias.
// If the alias is not present which is to be updated then returns error.
func (hc *Hailconfig) Update(alias, command, des string) error {
	if found := hc.IsPresent(alias); !found {
		return fmt.Errorf("alias is not found")
	}
	hc.Add(alias, command, des)
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
		return fmt.Errorf("alias is not found")
	}
	delete(hc.Scripts, alias)
	return nil
}

// Copy checks if oldAlias is present and newAlias is not present or return error.
// If above conditions are passed then it adds command with newAlias.
func (hc *Hailconfig) Copy(oldAlias, newAlias string) error {
	if !hc.IsPresent(oldAlias) {
		return errors.New("old alias is not present")
	}
	if hc.IsPresent(newAlias) {
		return errors.New("new alias is already present")
	}
	hc.Add(newAlias, hc.Scripts[oldAlias].Command, hc.Scripts[oldAlias].Description)
	return nil
}

// Move checks for both oldAlias and new alias is present, if not then returns error.
// Otherwise adds command with newAlias and then deletes the command with old alias.
func (hc *Hailconfig) Move(oldAlias, newAlias string) error {
	if !hc.IsPresent(oldAlias) {
		return errors.New("old alias is not present")
	}
	if hc.IsPresent(newAlias) {
		return errors.New("new alias is already present")
	}
	hc.Add(newAlias, hc.Scripts[oldAlias].Command, hc.Scripts[oldAlias].Description)
	return hc.Delete(oldAlias)
}

// Get returns command and error based on the alias name.
// It checks for the alias in Scipts map. Before calling Get we should check if
// the alias is present in map or not.
func (hc *Hailconfig) Get(alias string) (string, error) {
	return hc.Scripts[alias].Command, nil

}

// Parse Loads the file and then decodes the content from file field of *Hailconfig
// into config field. It returns error if any.
func (hc *Hailconfig) Parse() error {
	files, err := hc.loader.Load()
	if err != nil {
		return errors.Wrap(err, "failed to load")
	}
	f := files[0] // Currently only on .hailconfig file is supported.
	hc.f = f
	_, err = toml.DecodeReader(hc.f, &hc.config)
	if err != nil {
		return errors.Wrap(err, "failed to decode")
	}
	return nil
}
