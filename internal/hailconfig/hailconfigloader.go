package hailconfig

import (
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type hailconfigFile struct {
	*os.File
}
type ReadWriteResetCloser interface {
	io.ReadWriteCloser
	Reset() error
}

type Loader interface {
	Load() ([]ReadWriteResetCloser, error)
}

type StandardHailConfigLoader struct {
}

var DefaultLoader Loader = new(StandardHailConfigLoader)

func (StandardHailConfigLoader) Load() ([]ReadWriteResetCloser, error) {
	cfgPath, err := hailconfigPath()
	if err != nil {
		return nil, errors.Wrap(err, "cannot determine hailconfig path")
	}
	f, err := os.OpenFile(cfgPath, os.O_RDWR, 0)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.Wrap(err, ".hailconfig file is not found")
		}
		return nil, errors.Wrap(err, "failed to open the file")
	}
	return []ReadWriteResetCloser{ReadWriteResetCloser(&hailconfigFile{f})}, nil
}

func Init(title string) error {
	cfgfile, err := hailconfigPath()
	if err != nil {
		return errors.Wrap(err, "cannot determine .hailconfig path")
	}
	_, err = os.Open(cfgfile)
	if err != nil {
		if os.IsNotExist(err) {
			f, err := os.OpenFile(cfgfile, os.O_CREATE, 0755)
			if err != nil {
				return err
			}
			hc := new(Hailconfig).WithLoader(DefaultLoader)
			hc.config.Title = title
			hc.f = &hailconfigFile{f}
			return hc.Save()
		}
	}
	return errors.New(".hailconfig is already present, cannot do init")
}

func hailconfigPath() (string, error) {
	if v := os.Getenv("HAILCONFIG"); v != "" {
		return v, nil
	}
	// default path
	home := homeDir()
	if home == "" {
		return "", errors.New("HOME or USERPROFILE element is not set")
	}
	return filepath.Join(home, ".hailconfig"), nil
}

func homeDir() string {
	home := os.Getenv("HOME")
	if home == "" {
		home = os.Getenv("USERPROFILE") // windows
	}
	return home
}

// Reset truncates and set Seek to 0, 0 so that data can be written over the
// existing file from the top.
// It returns error "failed to truncate" and "failed to seek in file."
func (hc *hailconfigFile) Reset() error {
	err := hc.Truncate(0)
	if err != nil {
		return errors.Wrap(err, "failed to truncate")
	}
	_, err = hc.Seek(0, 0)
	return errors.Wrap(err, "failed to seek in file")
}
