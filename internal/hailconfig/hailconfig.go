package hailconfig

import (
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
	Alias   string
	Command string
}

type Hailconfig struct {
	loader  Loader
	f       ReadWriteResetCloser
	scripts []script
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

func (hc *Hailconfig) Add(alias, command string) error {
	sc := script{
		Alias:   alias,
		Command: command,
	}
	hc.scripts = append(hc.scripts, sc)
	err := toml.NewEncoder(hc.f).Encode(&sc)
	if err != nil {
		return errors.Wrap(err, "error while encoding")
	}
	return nil
}
