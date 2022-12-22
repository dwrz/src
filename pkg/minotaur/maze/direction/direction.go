package direction

import (
	"math/rand"
	"time"
)

type Direction int

const (
	Down Direction = iota
	Left
	Right
	Up
)

func (d Direction) Opposite() Direction {
	switch d {
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	case Up:
		return Down
	default:
		return d
	}
}

func (d Direction) String() string {
	switch d {
	case Down:
		return "down"
	case Left:
		return "left"
	case Right:
		return "right"
	case Up:
		return "up"
	default:
		return ""
	}
}

var (
	directions = [4]Direction{Down, Left, Right, Up}
	random     = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func Directions() [4]Direction {
	return [4]Direction{Down, Left, Right, Up}
}

func Random() Direction {
	return directions[random.Intn(len(directions))]
}

func RandomDirections() [4]Direction {
	var dirs = [...]Direction{Down, Left, Right, Up}

	random.Shuffle(len(dirs), func(i, j int) {
		dirs[i], dirs[j] = dirs[j], dirs[i]
	})

	return dirs
}
