package editor

import (
	"fmt"

	"code.dwrz.net/src/pkg/editor/message"
	"code.dwrz.net/src/pkg/terminal/input"
)

func (e *Editor) bufferInput(event *input.Event) error {
	size, err := e.terminal.Size()
	if err != nil {
		return fmt.Errorf("failed to get terminal size: %w", err)
	}

	if event.Rune != input.Null {
		switch event.Rune {
		case input.Delete:
			e.active.Backspace()

		case 's' & input.Control:
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
			e.active.Insert(event.Rune)
		}

		return nil
	}

	switch event.Key {
	case input.Down:
		e.active.CursorDown()

	case input.Left:
		e.active.CursorLeft()

	case input.Right:
		e.active.CursorRight()

	case input.Up:
		e.active.CursorUp()

	case input.Insert:

	case input.End:
		e.active.CursorEnd()

	case input.Home:
		e.active.CursorHome()

	case input.PageDown:
		e.active.PageDown(int(size.Rows))

	case input.PageUp:
		e.active.PageUp(int(size.Rows))

	default:
		e.log.Debug.Printf("unrecognized input: %#v", event)
	}

	return nil
}

func (e *Editor) promptInput(event *input.Event) error {
	switch event.Key {

	default:
		e.log.Debug.Printf("unrecognized input: %#v", event)
	}

	// If a newline was received, take the input.
	// Pass it back to the caller.
	// Make it the caller's job to set the prompt again if not happy
	// with return value.

	return nil
}
