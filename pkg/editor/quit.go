package editor

import (
	"code.dwrz.net/src/pkg/terminal"
)

func (e *Editor) quit() {
	e.out.Write([]byte(terminal.ClearScreen))
	e.out.Write([]byte(terminal.CursorTopLeft))

	if err := e.terminal.Reset(); err != nil {
		e.log.Error.Printf(
			"failed to reset terminal attributes: %v", err,
		)
	}
}
