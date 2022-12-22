package maze

import (
	"code.dwrz.net/src/pkg/minotaur/maze/cell"
	"code.dwrz.net/src/pkg/minotaur/maze/direction"
	"code.dwrz.net/src/pkg/minotaur/maze/position"
)

func inBounds(h, w, x, y int) bool {
	if x < 0 || x >= w {
		return false
	}
	if y < 0 || y >= h {
		return false
	}

	return true
}

// http://weblog.jamisbuck.org/2011/1/27/maze-generation-growing-tree-algorithm
func generate(h, w int) [][]cell.Cell {
	var layout = make([][]cell.Cell, h)

	// Create the rows.
	for i := range layout {
		layout[i] = make([]cell.Cell, w)
		for j := range layout[i] {
			layout[i][j] = 0
		}
	}

	// Connect the cells.
	var (
		current = position.Position{X: 0, Y: 0}
		stack   = []position.Position{current}
		visited = map[position.Position]struct{}{
			current: {},
		}
	)
	for len(stack) > 0 {
		// Get newest.
		current = stack[len(stack)-1]

		dirs := direction.RandomDirections()
		found := false
		for _, d := range dirs {
			var neighbor = current.WithDirection(d)

			// Ignore if out of bounds.
			if !inBounds(h, w, neighbor.X, neighbor.Y) {
				continue
			}

			// Ignore if already visited.
			if _, seen := visited[neighbor]; seen {
				continue
			}

			// We've found a cell to connect.
			found = true

			// Mark it as visited.
			visited[neighbor] = struct{}{}

			// Add it to the queue.
			stack = append(stack, neighbor)

			// Connect the cells.
			layout[current.Y][current.X].SetDirection(d)
			layout[neighbor.Y][neighbor.X].SetDirection(
				d.Opposite(),
			)

			break
		}

		// If no neighbor was found, remove this cell from the list.
		if !found {
			stack = stack[:len(stack)-1]
		}
	}

	return layout
}
