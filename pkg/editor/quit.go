package editor

func (e *Editor) quit() {
	e.canvas.Reset()

	if err := e.terminal.Reset(); err != nil {
		e.log.Error.Printf(
			"failed to reset terminal attributes: %v", err,
		)
	}
}
