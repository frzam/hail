package hailconfig

import (
	"bytes"
	"io"
	"strings"
)

type MockHailconfigLoader struct {
	in  io.Reader
	out bytes.Buffer
}

func (t *MockHailconfigLoader) Read(p []byte) (n int, err error) {
	return t.in.Read(p)
}

func (t *MockHailconfigLoader) Write(p []byte) (n int, err error) {
	return t.out.Write(p)
}

func (t *MockHailconfigLoader) Close() error {
	return nil
}

func (t *MockHailconfigLoader) Reset() error {
	return nil
}

func (t *MockHailconfigLoader) Load() ([]ReadWriteResetCloser, error) {
	return []ReadWriteResetCloser{ReadWriteResetCloser(t)}, nil
}

func WithMockHailconfigLoader(hailconfig string) *MockHailconfigLoader {
	return &MockHailconfigLoader{in: strings.NewReader(hailconfig)}
}
