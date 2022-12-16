package canvas

import "code.dwrz.net/src/pkg/terminal"

func (c *Canvas) Reset() {
	c.out.WriteString(terminal.ClearScreen)
	c.out.WriteString(terminal.CursorTopLeft)
}
