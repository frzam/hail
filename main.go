package main

import (
	"hail/cmd"
	"hail/internal/hailconfig"
	"io"
)

// Op ...
type Op interface {
	Run(hc *hailconfig.Hailconfig, stdout io.Writer) error
}

func main() {
	cmd.Execute()
}
