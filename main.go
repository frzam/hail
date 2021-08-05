package main

import (
	"hail/cmd"
	"hail/internal/hailconfig"
	"io"
)

// Op interface contains Run method which executes commands
type Op interface {
	Run(hc *hailconfig.Hailconfig, stdout io.Writer) error
}

func main() {
	cmd.Execute()
}
