package maze

import (
	"bytes"

	"code.dwrz.net/src/pkg/log"
	"code.dwrz.net/src/pkg/minotaur/maze/cell"
	"code.dwrz.net/src/pkg/minotaur/maze/direction"
	"code.dwrz.net/src/pkg/minotaur/maze/position"
	"code.dwrz.net/src/pkg/terminal"
)

type Maze struct {
	height int
	layout [][]cell.Cell
	log    *log.Logger
	width  int
}

type Parameters struct {
	Height int
	Log    *log.Logger
	Width  int
}

func New(p Parameters) (*Maze, error) {
	return &Maze{
		height: p.Height,
		layout: generate(p.Height, p.Width),
		log:    p.Log,
		width:  p.Width,
	}, nil
}

func (m *Maze) Cell(p position.Position) cell.Cell {
	return m.layout[p.Y][p.X]
}

func (m *Maze) InBounds(p position.Position) bool {
	return inBounds(m.height, m.width, p.X, p.Y)
}

func (m *Maze) Solve(start, end position.Position) (s []position.Position) {
	var q = [][]position.Position{{start}}
	for len(q) > 0 {
		// Get the current search from the queue.
		s, q = q[0], q[1:]

		// Use the last cell as the current candidate.
		var (
			current  = s[len(s)-1]
			cell     = m.Cell(current)
			previous *position.Position
		)

		// If we've reached the end, we have a solution.
		if current.Equal(end) {
			return s
		}

		// Explore new routes.
		// Determine the previous cell, to avoid backtracking.
		if len(s) > 1 {
			previous = &s[len(s)-2]
		}
		for _, d := range direction.Directions() {
			if !cell.HasDirection(d) {
				continue
			}

			// Get the next position with this direction.
			next := current.WithDirection(d)

			// Prevent backtracking.
			if previous != nil && next.Equal(*previous) {
				continue
			}

			// Next is valid; add it to the search queue.
			// TODO: is there a way to reuse the backing array?
			var route = make([]position.Position, len(s)+1)
			copy(route, s)
			route[len(route)-1] = next

			q = append(q, route)
		}
	}

	return nil
}

type RenderParameters struct {
	Passage string
	// Positions is used to render a cell with the mapped string.
	// The order in which you list positions matters.
	// If there are overlapping positions, the last set one wins.
	Positions map[position.Position]string
	Wall      string
}

func (m *Maze) Render(p RenderParameters) []byte {
	var buf bytes.Buffer

	buf.WriteString(terminal.CursorHide)
	buf.WriteString(terminal.CursorTopLeft)

	// Print the top wall.
	// We need to print each cell, and its right connection.
	// Thus the wall needs to be twice the width.
	// We also need to cover the right wall, thus +1.
	topWidth := 2*m.width + 1
	for i := 0; i < topWidth; i++ {
		buf.WriteString(p.Wall)
	}
	buf.WriteString("\r\n")

	for y, row := range m.layout {
		buf.WriteString(p.Wall) // left wall

		// Print each cell and the passage or wall to its right.
		for x, cell := range row {
			var block = p.Positions[position.Position{X: x, Y: y}]
			if block == "" {
				block = p.Passage
			}

			buf.WriteString(block)

			if cell.HasDirection(direction.Right) {
				buf.WriteString(p.Passage)
			} else {
				buf.WriteString(p.Wall)
			}
		}
		buf.WriteString("\r\n")

		// Print the cell's bottom passage or wall.
		// Print a wall to the cell's lower-right.
		buf.WriteString(p.Wall) // left wall
		for _, cell := range row {
			if cell.HasDirection(direction.Down) {
				buf.WriteString(p.Passage)
			} else {
				buf.WriteString(p.Wall)
			}

			buf.WriteString(p.Wall)
		}
		buf.WriteString("\r\n")
	}

	return buf.Bytes()
}
