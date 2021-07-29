package main

import (
	"hail/cmd"
	"hail/internal/hailconfig"
	"io"
)

type Op interface {
	Run(hc *hailconfig.Hailconfig, stdout io.Writer) error
}

func main() {
	cmd.Execute()
}
