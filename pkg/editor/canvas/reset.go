package canvas

import "code.dwrz.net/src/pkg/terminal"

func (c *Canvas) Reset() {
	c.out.Write([]byte(terminal.ClearScreen))
	c.out.Write([]byte(terminal.CursorTopLeft))
}
