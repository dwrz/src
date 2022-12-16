package editor

import (
	"fmt"

	"code.dwrz.net/src/pkg/editor/command"
	"code.dwrz.net/src/pkg/editor/input"
	"code.dwrz.net/src/pkg/editor/message"
)

func (e *Editor) bufferInput(input *input.Event) error {
	size, err := e.terminal.Size()
	if err != nil {
		return fmt.Errorf("failed to get terminal size: %w", err)
	}

	switch input.Command {
	case command.Backspace:
		e.active.Backspace()

	case command.CursorDown:
		e.active.CursorDown()

	case command.CursorLeft:
		e.active.CursorLeft()

	case command.CursorRight:
		e.active.CursorRight()

	case command.CursorUp:
		e.active.CursorUp()

	case command.Insert:
		e.active.Insert(input.Rune)

	case command.End:
		e.active.CursorEnd()

	case command.Home:
		e.active.CursorHome()

	case command.PageDown:
		e.active.PageDown(int(size.Rows))

	case command.PageUp:
		e.active.PageUp(int(size.Rows))

	case command.Save:
		// Get the filename.
		if err := e.active.Save(); err != nil {
			go func() {
				e.messages <- message.New(fmt.Sprintf(
					"failed to save: %v", err,
				))
			}()
		}
		go func() {
			e.messages <- message.New("saved file")
		}()

	default:
		e.log.Debug.Printf("unrecognized input: %#v", input)
	}

	return nil
}

func (e *Editor) promptInput(input *input.Event) error {
	switch input.Command {
	case command.Backspace:

	default:
		e.log.Debug.Printf("unrecognized input: %#v", input)
	}

	// If a newline was received, take the input.
	// Pass it back to the caller.
	// Make it the caller's job to set the prompt again if not happy
	// with return value.

	return nil
}
