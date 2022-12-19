package editor

import (
	"fmt"
	"os"
	"sync"

	"code.dwrz.net/src/pkg/editor/buffer"
	"code.dwrz.net/src/pkg/editor/canvas"
	"code.dwrz.net/src/pkg/editor/message"
	"code.dwrz.net/src/pkg/log"
	"code.dwrz.net/src/pkg/terminal"
	"code.dwrz.net/src/pkg/terminal/input"
)

type Editor struct {
	canvas   *canvas.Canvas
	errs     chan error
	input    chan *input.Event
	log      *log.Logger
	messages chan *message.Message
	reader   *input.Reader
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
		input:    make(chan *input.Event),
		log:      p.Log,
		messages: make(chan *message.Message),
		terminal: p.Terminal,
		tmpdir:   p.TempDir,
	}

	// Setup user input.
	editor.reader = input.New(input.Parameters{
		Chan: editor.input,
		In:   p.In,
		Log:  p.Log,
	})

	// Setup the canvas.
	editor.canvas = canvas.New(canvas.Parameters{
		Log:      p.Log,
		Out:      p.Out,
		Terminal: p.Terminal,
	})

	// Create the initial scratch buffer.
	scratch, err := buffer.Create(buffer.NewBufferParams{
		Name: editor.tmpdir + "/scratch",
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
