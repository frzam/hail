package hailconfig

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// hailconfigFile has file type embedded into it.
type hailconfigFile struct {
	*os.File
}

// ReadWriteResetCloser has ReadWriteCloser interface and Reset method
// defined in it.
type ReadWriteResetCloser interface {
	io.ReadWriteCloser
	Reset() error
}

// Loader interface has only Load method define on it,
// it takes nothing and returns a slice of ReadWriterResetCloser interface
// error.
type Loader interface {
	Load() ([]ReadWriteResetCloser, error)
}

// StandardHailConfigLoader is an empty struct to make
// Load as a method.
type StandardHailConfigLoader struct {
}

// DefaultLoader ...
var DefaultLoader Loader = new(StandardHailConfigLoader)

// Load returns .hailconfig file as ReadWriteResetCloser. It gets the path from
// hailconfigPath then opens the file as RDWR and returns the file as ReadWriteResetCloser.
// Otherwise returns nil and error.
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

// Init func is used to create a new .hailconfig file with title given. It checks if file is not
// present then only it creates a .hailconfig file with title and saves it. Otherwise returns
// error that .hailconfig is already present.
func Init(title string) (string, error) {
	// default interpreter is bash
	interpreter := "bash"
	cfgfile, err := hailconfigPath()
	if err != nil {
		return "", errors.Wrap(err, "cannot determine .hailconfig path")
	}
	if _, err = os.Stat(cfgfile); os.IsNotExist(err) {
		f, err := os.OpenFile(cfgfile, os.O_CREATE, 0755)
		if err != nil {
			return "", err
		}
		hc := new(Hailconfig).WithLoader(DefaultLoader)
		hc.config.Title = title
		hc.config.Interpreter = interpreter
		hc.f = &hailconfigFile{f}
		return cfgfile, nil
	}
	// If file is already present
	return "", fmt.Errorf(".hailconfig is already present, can't do init at loc: %s", cfgfile)
}

// hailconfigPath looks for HAILCONFIG env variable, if it is not present then
// it looks for home path and if not found returns error.
// If home path is found then it add .hailconfig name returns the complete path.
func hailconfigPath() (string, error) {
	if v := os.Getenv("HAILCONFIG"); v != "" {
		return filepath.Join(v, ".hailconfig"), nil
	}
	// default path
	home := homeDir()
	if home == "" {
		return "", errors.New("HOME or USERPROFILE element is not set")
	}
	return filepath.Join(home, ".hailconfig"), nil
}

// homeDir is used find the home path of the system.
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
