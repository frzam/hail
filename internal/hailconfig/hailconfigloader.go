package hailconfig

import (
	"io"
	"os"

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

type StandardHailConfigLoader struct{}

var DefaultLoader Loader = new(StandardHailConfigLoader)

func (StandardHailConfigLoader) Load() ([]ReadWriteResetCloser, error) {
	cfgPath := ".hailconfig"
	f, err := os.OpenFile(cfgPath, os.O_RDWR, 0)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.Wrap(err, ".hailconfig file is not found")
		}
		return nil, errors.Wrap(err, "failed to open the file")
	}
	return []ReadWriteResetCloser{ReadWriteResetCloser(&hailconfigFile{f})}, nil
}

func hailconfigPath(path string) (string, error) {

	return "", nil
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
