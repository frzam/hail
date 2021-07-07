package hailconfig

import (
	"fmt"
	"io"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type ReadWriteResetCloser interface {
	io.ReadWriteCloser
	Reset() error
}

type Loader interface {
	Load() ([]ReadWriteResetCloser, error)
}

type script struct {
	Command string `toml:"command"`
}

type Hailconfig struct {
	loader Loader
	f      ReadWriteResetCloser
	config config
}
type config struct {
	Title   string
	Scripts map[string]script `toml:"scripts"`
}

func (hc *Hailconfig) WithLoader(l Loader) *Hailconfig {
	hc.loader = l
	return hc
}

func (hc *Hailconfig) Close() error {
	if hc.f == nil {
		return nil
	}
	return hc.f.Close()
}

func (hc *Hailconfig) Add(alias, command string) {
	sc := script{
		Command: command,
	}
	hc.config.Scripts[alias] = sc
}

func (hc *Hailconfig) Save() error {
	err := hc.f.Reset()
	if err != nil {
		return errors.Wrap(err, "failed to reset")
	}
	return toml.NewEncoder(hc.f).Encode(&hc.config)
}

func (hc *Hailconfig) List() error {
	for alias, script := range hc.config.Scripts {
		fmt.Printf("%s\t\t%s\n", alias, script.Command)
	}
	return nil
}
func (hc *Hailconfig) Update(alias, command string) error {
	if found := hc.IsPresent(alias); !found {
		return errors.New("alias '" + alias + "' is not found")
	}
	hc.Add(alias, command)
	return nil
}

func (hc *Hailconfig) IsPresent(alias string) bool {
	_, found := hc.config.Scripts[alias]
	return found
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
