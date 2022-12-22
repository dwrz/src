package cell

import (
	"testing"

	"code.dwrz.net/src/pkg/minotaur/maze/direction"
)

func TestSetDirection(t *testing.T) {
	var c Cell
	t.Logf("start: %08b", c)

	c.SetDirection(direction.Right)
	t.Logf("set right: %08b", c)

	c.SetDirection(direction.Down)
	t.Logf("set down: %08b", c)

	var (
		hasRight = c.HasDirection(direction.Right)
		hasDown  = c.HasDirection(direction.Down)
		hasLeft  = c.HasDirection(direction.Left)
		hasUp    = c.HasDirection(direction.Up)
	)

	t.Logf("has right: %v", hasRight)
	t.Logf("has down: %v", hasDown)
	t.Logf("has left: %v", hasLeft)
	t.Logf("has up: %v", hasUp)

	if !hasRight {
		t.Errorf("expected cell to have right")
	}
	if !hasDown {
		t.Errorf("expected cell to have down")
	}
	if hasLeft {
		t.Errorf("expected cell to not have left")
	}
	if hasUp {
		t.Errorf("expected cell to not have up")
	}
}
