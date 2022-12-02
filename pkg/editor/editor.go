package editor

import (
	"fmt"
	"os"
	"sync"

	"code.dwrz.net/src/pkg/editor/buffer"
	"code.dwrz.net/src/pkg/log"
	"code.dwrz.net/src/pkg/terminal"
)

type Editor struct {
	errs     chan error
	in       *os.File
	input    chan input
	log      *log.Logger
	messages chan Message
	out      *os.File
	terminal *terminal.Terminal
	tmpdir   string

	mu      sync.Mutex
	active  *buffer.Buffer
	buffers map[string]*buffer.Buffer
}

type Parameters struct {
	In       *os.File
	Log      *log.Logger
	Out      *os.File
	TempDir  string
	Terminal *terminal.Terminal
}

func New(p Parameters) (*Editor, error) {
	var editor = &Editor{
		buffers:  map[string]*buffer.Buffer{},
		errs:     make(chan error),
		in:       p.In,
		input:    make(chan input),
		log:      p.Log,
		messages: make(chan Message),
		out:      p.Out,
		terminal: p.Terminal,
		tmpdir:   p.TempDir,
	}

	// Create the initial scratch buffer.
	scratch, err := buffer.Create(buffer.NewBufferParams{
		Name: editor.tmpdir + "scratch",
		Log:  editor.log,
	})
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create scratch buffer: %v", err,
		)
	}
	editor.active = scratch
	editor.buffers[scratch.Name()] = scratch

	return editor, nil
}