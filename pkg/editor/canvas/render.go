package canvas

import (
	"fmt"
	"path"

	"github.com/mattn/go-runewidth"

	"code.dwrz.net/src/pkg/color"
	"code.dwrz.net/src/pkg/editor/buffer"
	"code.dwrz.net/src/pkg/editor/message"
	"code.dwrz.net/src/pkg/terminal"
)

func (c *Canvas) Render(b *buffer.Buffer, msg *message.Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	size, err := c.terminal.Size()
	if err != nil {
		return fmt.Errorf("failed to get terminal size: %w", err)
	}

	var (
		cursor = b.Cursor()
		bars   = 2
		height = int(size.Rows) - bars
		width  = int(size.Columns)
		output = b.Render(height, width)
	)

	// Move the cursor to the top left.
	c.buf.WriteString(terminal.CursorHide)
	c.buf.WriteString(terminal.CursorTopLeft)

	// Print each line.
	for _, line := range output.Lines {
		c.buf.WriteString(terminal.EraseLine)
		c.buf.WriteString(line)
		c.buf.WriteString("\r\n")
	}

	// Draw the status bar.
	c.statusBar(b, width, cursor.Line(), cursor.Glyph())

	// Draw the message bar.
	c.messageBar(msg, width)

	// Set the cursor.
	c.buf.WriteString(fmt.Sprintf(
		terminal.SetCursorFmt, output.Line, output.Glyph,
	))
	c.buf.WriteString(terminal.CursorShow)

	c.out.Write(c.buf.Bytes())

	c.buf.Reset()

	return nil
}

func (c *Canvas) statusBar(b *buffer.Buffer, width, y, x int) {
	c.buf.WriteString(terminal.EraseLine)
	c.buf.WriteString(color.Inverse)

	// Icon
	icon := "文 "
	c.buf.WriteString(icon)
	width -= runewidth.StringWidth(icon)

	// Calculate the length of the cursor, so we can determine how much
	// space we have left for the name of the buffer.
	cursor := fmt.Sprintf(" %d:%d", y, x)
	width -= runewidth.StringWidth(cursor)

	// Filename.
	// TODO: handle long filenames (shorten filepath).
	name := path.Base(b.Name())
	nw := runewidth.StringWidth(name)
	if nw <= width {
		c.buf.WriteString(name)
		width -= nw
	} else {
		for _, r := range name {
			rw := runewidth.RuneWidth(r)
			if width-rw >= 0 {
				c.buf.WriteRune(r)
				width -= rw
			} else {
				break
			}
		}
	}

	// Add empty spaces to the end of the line.
	for i := width; i >= 0; i-- {
		c.buf.WriteRune(' ')
	}

	// Cursor
	c.buf.WriteString(cursor)

	c.buf.WriteString(color.Reset)
	c.buf.WriteString("\r\n")
}

// TODO: handle messages that are too long for one line.
func (c *Canvas) messageBar(msg *message.Message, width int) {
	c.buf.WriteString(terminal.EraseLine)
	switch {
	case msg == nil:
		c.buf.WriteString(fmt.Sprintf("%*s", width, " "))
	case len(msg.Text) > width:
		c.buf.WriteString(fmt.Sprintf("%*s", width, msg.Text))
	default:
		c.buf.WriteString(fmt.Sprintf(
			"%s %*s", msg.Text, width-len(msg.Text), " ",
		))
	}
}
