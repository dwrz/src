package editor

import (
	"bufio"
	"fmt"
	"unicode/utf8"

	"code.dwrz.net/src/pkg/terminal"
)

type command int

const (
	Backspace command = iota
	CursorDown
	CursorLeft
	CursorRight
	CursorUp
	Delete
	End
	Home
	Insert
	Open
	PageDown
	PageUp
	Quit
	Save
)

type input struct {
	Command command
	Rune    rune
}

// TODO: reading one rune at a time is slow, especially when pasting large
// quantities of text into the active buffer. It would be nice to take more
// input at once, while still being able to handle escape sequences without
// too many edge cases.
func (e *Editor) readInput() {
	var buf = bufio.NewReader(e.in)
	for {
		r, size, err := buf.ReadRune()
		if err != nil {
			e.errs <- fmt.Errorf("failed to read stdin: %w", err)
		}
		e.log.Debug.Printf(
			"read rune %s %v (%d)",
			string(r), []byte(string(r)), size,
		)
		switch r {
		case utf8.RuneError:
			e.log.Error.Printf(
				"rune error: %s (%d)", string(r), size,
			)

		// Handle escape sequences.
		case terminal.Escape:
			e.parseEscapeSequence(buf)

		case terminal.Delete:
			e.input <- input{Command: Backspace}

		case 'q' & terminal.Control:
			e.input <- input{Command: Quit}

		case 's' & terminal.Control:
			e.input <- input{Command: Save}

		default:
			e.input <- input{Command: Insert, Rune: r}
		}
	}
}

func (e *Editor) parseEscapeSequence(buf *bufio.Reader) {
	r1, _, err := buf.ReadRune()
	if err != nil {
		e.errs <- fmt.Errorf("failed to read stdin: %w", err)
		return
	}

	// Ignore invalid escape sequences.
	if r1 != '[' && r1 != 'O' {
		e.input <- input{Command: Insert, Rune: r1}
		return
	}

	// We've received an input of Esc + [ or Esc + O.
	// Determine the escape sequence.
	r2, _, err := buf.ReadRune()
	if err != nil {
		e.errs <- fmt.Errorf("failed to read stdin: %w", err)
		return
	}

	// Check letter escape sequences.
	switch r2 {
	case 'A':
		e.input <- input{Command: CursorUp}
		return
	case 'B':
		e.input <- input{Command: CursorDown}
		return
	case 'C':
		e.input <- input{Command: CursorRight}
		return
	case 'D':
		e.input <- input{Command: CursorLeft}
		return

	case 'O':
		r3, _, err := buf.ReadRune()
		if err != nil {
			e.errs <- fmt.Errorf("failed to read stdin: %w", err)
		}
		switch r3 {
		case 'P': // F1
			return
		case 'Q': // F2
			return
		case 'R': // F3
			return
		case 'S': // F4
			return
		default:
			// No match.
			e.input <- input{Command: Insert, Rune: r1}
			e.input <- input{Command: Insert, Rune: r2}
			e.input <- input{Command: Insert, Rune: r3}
			return
		}
	}

	// Check for single digit numerical escape sequences.
	r3, _, err := buf.ReadRune()
	if err != nil {
		e.errs <- fmt.Errorf("failed to read stdin: %w", err)
	}
	switch {
	case r2 == '1' && r3 == '~':
		e.input <- input{Command: Home}
		return
	case r2 == '2' && r3 == '~':
		e.input <- input{Command: Insert}
		return
	case r2 == '3' && r3 == '~':
		e.input <- input{Command: Delete}
		return
	case r2 == '4' && r3 == '~':
		e.input <- input{Command: End}
		return
	case r2 == '5' && r3 == '~':
		e.input <- input{Command: PageUp}
		return
	case r2 == '6' && r3 == '~':
		e.input <- input{Command: PageDown}
		return
	case r2 == '7' && r3 == '~':
		e.input <- input{Command: Home}
		return
	case r2 == '8' && r3 == '~':
		e.input <- input{Command: End}
		return
	case r2 == '9' && r3 == '~':
		e.input <- input{Command: End}
		return
	}

	// Check for double digit numerical escape sequences.
	r4, _, err := buf.ReadRune()
	if err != nil {
		e.errs <- err
	}
	switch {
	case r2 == '1' && r3 == '0' && r4 == '~':
		return
	case r2 == '1' && r3 == '1' && r4 == '~':
		return
	case r2 == '1' && r3 == '2' && r4 == '~':
		return
	case r2 == '1' && r3 == '3' && r4 == '~':
		return
	case r2 == '1' && r3 == '4' && r4 == '~':
		return
	case r2 == '1' && r3 == '4' && r4 == '~':
		return
	case r2 == '1' && r3 == '6' && r4 == '~':
		return
	case r2 == '1' && r3 == '7' && r4 == '~':
		return
	case r2 == '1' && r3 == '8' && r4 == '~':
		return
	case r2 == '1' && r3 == '9' && r4 == '~':
		return
	case r2 == '2' && r3 == '0' && r4 == '~':
		return
	case r2 == '2' && r3 == '1' && r4 == '~':
		return
	case r2 == '2' && r3 == '2' && r4 == '~':
		return
	case r2 == '2' && r3 == '3' && r4 == '~':
		return
	case r2 == '2' && r3 == '4' && r4 == '~':
		return
	case r4 == '~':
		return
	}

	// No match.
	e.input <- input{Command: Insert, Rune: r1}
	e.input <- input{Command: Insert, Rune: r2}
	e.input <- input{Command: Insert, Rune: r3}
	e.input <- input{Command: Insert, Rune: r4}

}

func (e *Editor) processInput(input input) error {
	size, err := e.terminal.Size()
	if err != nil {
		return fmt.Errorf("failed to get terminal size: %w", err)
	}

	switch input.Command {
	case Backspace:
		e.active.Backspace()

	case CursorDown:
		e.active.CursorDown()

	case CursorLeft:
		e.active.CursorLeft()

	case CursorRight:
		e.active.CursorRight()

	case CursorUp:
		e.active.CursorUp()

	case Insert:
		e.active.Insert(input.Rune)

	case End:
		e.active.CursorEnd()

	case Home:
		e.active.CursorHome()

	case PageDown:
		e.active.PageDown(int(size.Rows))

	case PageUp:
		e.active.PageUp(int(size.Rows))

	case Save:
		if err := e.active.Save(); err != nil {
			go func() {
				e.messages <- Message{
					Text: fmt.Sprintf(
						"failed to save: %v", err,
					),
				}
			}()
		}
		go func() {
			e.messages <- Message{
				Text: fmt.Sprintf("saved file"),
			}
		}()

	default:
		e.log.Debug.Printf("unrecognized input: %#v", input)
	}

	return nil
}
