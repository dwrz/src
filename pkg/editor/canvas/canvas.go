package canvas

import (
	"bytes"
	"os"

	"code.dwrz.net/src/pkg/terminal"
)

type Canvas struct {
	buf      bytes.Buffer
	out      *os.File
	terminal *terminal.Terminal
}

type Parameters struct {
	Out      *os.File
	Terminal *terminal.Terminal
}

func New(p Parameters) *Canvas {
	return &Canvas{
		out:      p.Out,
		terminal: p.Terminal,
	}
}
