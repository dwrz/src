package canvas

import (
	"bytes"
	"os"
	"sync"

	"code.dwrz.net/src/pkg/log"
	"code.dwrz.net/src/pkg/terminal"
)

type Canvas struct {
	log      *log.Logger
	out      *os.File
	terminal *terminal.Terminal

	mu  sync.Mutex
	buf bytes.Buffer
}

type Parameters struct {
	Log      *log.Logger
	Out      *os.File
	Terminal *terminal.Terminal
}

func New(p Parameters) *Canvas {
	return &Canvas{
		log:      p.Log,
		out:      p.Out,
		terminal: p.Terminal,
	}
}
