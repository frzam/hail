package hailconfig

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type hailconfigFile struct {
	*os.File
}

type StandardHailConfigLoader struct{}

var DefaultLoader Loader = new(StandardHailConfigLoader)

func (StandardHailConfigLoader) Load() ([]ReadWriteResetCloser, error) {
	cfgPath := "hail.toml"
	f, err := os.OpenFile(cfgPath, os.O_RDWR, 0)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.Wrap(err, "hail.toml file is not found")
		}
		return nil, errors.Wrap(err, "failed to open the file")
	}
	return []ReadWriteResetCloser{ReadWriteResetCloser(&hailconfigFile{f})}, nil
}

func (hc *Hailconfig) Parse() error {
	files, err := hc.loader.Load()
	if err != nil {
		return errors.Wrap(err, "failed to load")
	}
	f := files[0]
	hc.f = f
	var sc script
	_, err = toml.DecodeReader(hc.f, &sc)
	if err != nil {
		return errors.Wrap(err, "failed to decode")
	}
	fmt.Println(sc)
	return nil
}

func hailconfigPath(path string) (string, error) {

	return "", nil
}

func (hc *hailconfigFile) Reset() error {
	return nil
}
