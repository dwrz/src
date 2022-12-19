package input

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"unicode/utf8"

	"code.dwrz.net/src/pkg/log"
)

type Event struct {
	Key  Key
	Rune rune
}

type Reader struct {
	buf    *bufio.Reader
	events chan *Event
	log    *log.Logger
}

type Parameters struct {
	Chan chan *Event
	In   *os.File
	Log  *log.Logger
}

func New(p Parameters) *Reader {
	return &Reader{
		buf:    bufio.NewReader(p.In),
		events: p.Chan,
		log:    p.Log,
	}
}

// TODO: reading one rune at a time is slow, especially when pasting large
// quantities of text into the active buffer. It would be nice to take more
// input at once, while still being able to handle escape sequences without
// too many edge cases.
func (i *Reader) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			r, size, err := i.buf.ReadRune()
			if err != nil {
				return fmt.Errorf("failed to read: %w", err)
			}
			i.log.Debug.Printf(
				"read rune %s %v (%d)",
				string(r), []byte(string(r)), size,
			)
			switch r {
			case utf8.RuneError:
				i.log.Error.Printf(
					"rune error: %s (%d)", string(r), size,
				)

				// Handle escape sequences.
			case Escape:
				if err := i.parseEscapeSequence(); err != nil {
					return fmt.Errorf(
						"failed to read: %w", err,
					)
				}

			default:
				i.events <- &Event{Rune: r}
			}
		}
	}
}

func (i *Reader) parseEscapeSequence() error {
	r1, _, err := i.buf.ReadRune()
	if err != nil {
		return fmt.Errorf("failed to read: %w", err)
	}

	// Ignore invalid escape sequences.
	if r1 != '[' && r1 != 'O' {
		i.events <- &Event{Rune: r1}
		return nil
	}

	// We've received an input of Esc + [ or Esc + O.
	// Determine the escape sequence.
	r2, _, err := i.buf.ReadRune()
	if err != nil {
		return fmt.Errorf("failed to read: %w", err)

	}

	// Check letter escape sequences.
	switch r2 {
	case 'A':
		i.events <- &Event{Key: Up}
		return nil
	case 'B':
		i.events <- &Event{Key: Down}
		return nil
	case 'C':
		i.events <- &Event{Key: Right}
		return nil
	case 'D':
		i.events <- &Event{Key: Left}
		return nil

	case 'O':
		r3, _, err := i.buf.ReadRune()
		if err != nil {
			return fmt.Errorf("failed to read: %w", err)
		}
		switch r3 {
		case 'P': // F1
			return nil
		case 'Q': // F2
			return nil
		case 'R': // F3
			return nil
		case 'S': // F4
			return nil
		default:
			// No match.
			i.events <- &Event{Rune: r1}
			i.events <- &Event{Rune: r2}
			i.events <- &Event{Rune: r3}
			return nil
		}
	}

	// Check for single digit numerical escape sequences.
	r3, _, err := i.buf.ReadRune()
	if err != nil {
		return fmt.Errorf("failed to read: %w", err)
	}
	switch {
	case r2 == '1' && r3 == '~':
		i.events <- &Event{Key: Home}
		return nil
	case r2 == '2' && r3 == '~':
		i.events <- &Event{Key: Insert}
		return nil
	case r2 == '3' && r3 == '~':
		i.events <- &Event{Rune: Delete}
		return nil
	case r2 == '4' && r3 == '~':
		i.events <- &Event{Key: End}
		return nil
	case r2 == '5' && r3 == '~':
		i.events <- &Event{Key: PageUp}
		return nil
	case r2 == '6' && r3 == '~':
		i.events <- &Event{Key: PageDown}
		return nil
	case r2 == '7' && r3 == '~':
		i.events <- &Event{Key: Home}
		return nil
	case r2 == '8' && r3 == '~':
		i.events <- &Event{Key: End}
		return nil
	case r2 == '9' && r3 == '~':
		i.events <- &Event{Key: End}
		return nil
	}

	// Check for double digit numerical escape sequences.
	r4, _, err := i.buf.ReadRune()
	if err != nil {
		return fmt.Errorf("failed to read: %w", err)
	}
	switch {
	case r2 == '1' && r3 == '0' && r4 == '~':
		return nil
	case r2 == '1' && r3 == '1' && r4 == '~':
		return nil
	case r2 == '1' && r3 == '2' && r4 == '~':
		return nil
	case r2 == '1' && r3 == '3' && r4 == '~':
		return nil
	case r2 == '1' && r3 == '4' && r4 == '~':
		return nil
	case r2 == '1' && r3 == '4' && r4 == '~':
		return nil
	case r2 == '1' && r3 == '6' && r4 == '~':
		return nil
	case r2 == '1' && r3 == '7' && r4 == '~':
		return nil
	case r2 == '1' && r3 == '8' && r4 == '~':
		return nil
	case r2 == '1' && r3 == '9' && r4 == '~':
		return nil
	case r2 == '2' && r3 == '0' && r4 == '~':
		return nil
	case r2 == '2' && r3 == '1' && r4 == '~':
		return nil
	case r2 == '2' && r3 == '2' && r4 == '~':
		return nil
	case r2 == '2' && r3 == '3' && r4 == '~':
		return nil
	case r2 == '2' && r3 == '4' && r4 == '~':
		return nil
	case r4 == '~':
		return nil
	}

	// No match.
	i.events <- &Event{Rune: r1}
	i.events <- &Event{Rune: r2}
	i.events <- &Event{Rune: r3}
	i.events <- &Event{Rune: r4}

	return nil
}
