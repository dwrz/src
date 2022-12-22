package position

import (
	"fmt"
	"math/rand"
	"time"

	"code.dwrz.net/src/pkg/minotaur/maze/direction"
)

type Position struct {
	X, Y int
}

func (p Position) Equal(v Position) bool {
	return p.X == v.X && p.Y == v.Y
}

func (p Position) Down() Position {
	return Position{X: p.X, Y: p.Y + 1}
}

func (p Position) Left() Position {
	return Position{X: p.X - 1, Y: p.Y}
}

func (p Position) Right() Position {
	return Position{X: p.X + 1, Y: p.Y}
}

func (p Position) Up() Position {
	return Position{X: p.X, Y: p.Y - 1}
}

func (p Position) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func (p Position) WithDirection(d direction.Direction) Position {
	switch d {
	case direction.Down:
		return p.Down()
	case direction.Left:
		return p.Left()
	case direction.Right:
		return p.Right()
	case direction.Up:
		return p.Up()
	default:
		return p
	}
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func Random(width, height int) Position {
	return Position{
		X: random.Intn(width),
		Y: random.Intn(height),
	}
}
