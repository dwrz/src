package cell

import (
	"code.dwrz.net/src/pkg/minotaur/maze/direction"
)

type Cell uint8

func (c Cell) HasDirection(d direction.Direction) bool {
	return c&(1<<d) != 0
}

func (c *Cell) SetDirection(d direction.Direction) {
	*c |= 1 << d
}
