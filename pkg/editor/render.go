package editor

import (
	"bytes"
	"fmt"
	"path"
	"strings"

	"github.com/mattn/go-runewidth"

	"code.dwrz.net/src/pkg/color"
	"code.dwrz.net/src/pkg/terminal"
)

func (e *Editor) render(msg *Message) error {
	size, err := e.terminal.Size()
	if err != nil {
		return fmt.Errorf("failed to get terminal size: %w", err)
	}

	var (
		buf    bytes.Buffer
		cursor = e.active.Cursor()
		bars   = 2
		height = int(size.Rows) - bars
		width  = int(size.Columns)
		output = e.active.Render(height, width)
	)

	// Move the cursor to the top left.
	buf.Write([]byte(terminal.CursorHide))
	buf.Write([]byte(terminal.CursorTopLeft))

	// Print each line.
	for _, line := range output.Lines {
		buf.Write([]byte(terminal.EraseLine))
		buf.WriteString(line)
		buf.WriteString("\r\n")
	}

	// Draw the status bar.
	buf.Write([]byte(terminal.EraseLine))
	buf.WriteString(e.statusBar(width, cursor.Line(), cursor.Glyph()))
	buf.WriteString("\r\n")

	// Draw the message bar.
	buf.Write([]byte(terminal.EraseLine))
	buf.WriteString(e.messageBar(msg, width))

	// Set the cursor.
	buf.Write([]byte(
		fmt.Sprintf(terminal.SetCursorFmt, output.Line, output.Glyph)),
	)
	buf.Write([]byte(terminal.CursorShow))

	e.out.Write(buf.Bytes())

	return nil
}

// TODO: show a character cursor, not the terminal cursor.
func (e *Editor) statusBar(width, y, x int) string {
	var bar strings.Builder

	bar.WriteString(color.Inverse)

	// Icon
	icon := "文 "
	bar.WriteString(icon)
	width -= runewidth.StringWidth(icon)

	// Calculate the length of the cursor, so we can determine how much
	// space we have left for the name of the buffer.
	cursor := fmt.Sprintf(" %d:%d", y, x)
	width -= runewidth.StringWidth(cursor)

	// Filename.
	// TODO: handle long filenames (shorten filepath).
	name := path.Base(e.active.Name())
	nw := runewidth.StringWidth(name)
	if nw <= width {
		bar.WriteString(name)
		width -= nw
	} else {
		for _, r := range name {
			rw := runewidth.RuneWidth(r)
			if width-rw >= 0 {
				bar.WriteRune(r)
				width -= rw
			} else {
				break
			}
		}
	}

	// Add empty spaces to the end of the line.
	for i := width; i >= 0; i-- {
		bar.WriteRune(' ')
	}

	// Cursor
	bar.WriteString(cursor)

	bar.WriteString(color.Reset)

	return bar.String()
}

// TODO: handle messages that are too long for one line.
// TODO: should status bar render independently?
func (e *Editor) messageBar(msg *Message, width int) string {
	switch {
	case msg == nil:
		return fmt.Sprintf("%*s", width, " ")
	case len(msg.Text) > width:
		return fmt.Sprintf("%*s", width, msg.Text)
	default:
		return fmt.Sprintf("%s %*s", msg.Text, width-len(msg.Text), " ")
	}

}
