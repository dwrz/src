package minotaur

import (
	"context"
	"fmt"
	"os"
	"time"

	"code.dwrz.net/src/pkg/color"
	"code.dwrz.net/src/pkg/log"
	"code.dwrz.net/src/pkg/minotaur/maze"
	"code.dwrz.net/src/pkg/minotaur/maze/direction"
	"code.dwrz.net/src/pkg/minotaur/maze/position"
	"code.dwrz.net/src/pkg/terminal"
	"code.dwrz.net/src/pkg/terminal/input"
)

const (
	// Game over messages.
	slain   = "You were slain by the Minotaur. Score: %v.\r\n"
	escaped = "You escaped the labyrinth. Score: %v.\r\n"

	// Blocks for rendering.
	end      = color.BackgroundGreen + "  " + color.Reset
	minotaur = color.BackgroundYellow + "  " + color.Reset
	passage  = "  "
	theseus  = color.BackgroundMagenta + "  " + color.Reset
	solution = color.BackgroundBlue + "  " + color.Reset
	wall     = color.BackgroundBlack + "  " + color.Reset
)

type Game struct {
	errs     chan error
	events   chan *input.Event
	input    *input.Reader
	log      *log.Logger
	maze     *maze.Maze
	out      *os.File
	start    time.Time
	terminal *terminal.Terminal
	ticker   *time.Ticker

	// Maze positions.
	end      position.Position
	theseus  position.Position
	minotaur position.Position
	solution []position.Position
}

type Parameters struct {
	Height   int
	In       *os.File
	Log      *log.Logger
	Out      *os.File
	Terminal *terminal.Terminal
	Tick     time.Duration
	Width    int
}

func New(p Parameters) (*Game, error) {
	// Account for two spaces per cell; horizontal passages and walls.
	// Account an extra row for the right wall.
	if p.Width%4 == 0 {
		p.Width -= 1
	}
	p.Width /= 4

	// Account extra rows for vertical passages and walls.
	if p.Height%2 == 0 {
		p.Height -= 1
	}
	p.Height /= 2

	// If either the terminal width or height are zero,
	// then the terminal is too small to render the maze.
	if p.Height == 0 || p.Width == 0 {
		return nil, fmt.Errorf("terminal too small")
	}

	var g = &Game{
		errs:     make(chan error),
		events:   make(chan *input.Event),
		log:      p.Log,
		out:      p.Out,
		terminal: p.Terminal,
		ticker:   time.NewTicker(p.Tick),
	}

	maze, err := maze.New(maze.Parameters{
		Height: p.Height,
		Log:    g.log,
		Width:  p.Width,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to generate maze: %v", err)
	}
	g.maze = maze

	g.input = input.New(input.Parameters{
		Chan: g.events,
		In:   os.Stdin,
		Log:  g.log,
	})

	// Set start and end.
	g.theseus = position.Random(p.Width, p.Height)
	g.end = position.Random(p.Width, p.Height)
	g.minotaur = position.Random(p.Width, p.Height)

	return g, nil
}

func (g *Game) Run(ctx context.Context) error {
	// Setup the terminal.
	if err := g.terminal.SetRaw(); err != nil {
		return fmt.Errorf(
			"failed to set terminal attributes: %v", err,
		)
	}

	// Clean up before exit.
	defer g.quit()

	// TODO: refactor to use a terminal output package.
	g.out.WriteString(terminal.ClearScreen)
	g.out.WriteString(terminal.CursorHide)
	g.render()

	// Start reading user input.
	go func() {
		if err := g.input.Run(ctx); err != nil {
			g.errs <- err
		}
	}()

	// Main loop.
	g.start = time.Now()
	for {
		select {
		case <-ctx.Done():
			return nil

		case err := <-g.errs:
			return err

		case <-g.ticker.C:
			g.moveMinotaur()
			g.render()

		case event := <-g.events:
			switch event.Rune {
			case 'q': // Quit
				return nil

			case 's': // Toggle the solution.
				if g.solution == nil {
					g.solve()
				} else {
					g.solution = nil
				}
				g.render()
				continue
			}
			switch event.Key {
			case input.Up:
				g.moveTheseus(direction.Up)
			case input.Left:
				g.moveTheseus(direction.Left)
			case input.Down:
				g.moveTheseus(direction.Down)
			case input.Right:
				g.moveTheseus(direction.Right)
			}
			// If the solution is set, update it.
			if g.solution != nil {
				g.solve()
			}

			g.render()

		default:
			// Check for game over conditions.
			switch {
			case g.escaped():
				fmt.Fprintf(g.out, escaped, g.score())
				return nil

			case g.slain():
				fmt.Fprintf(g.out, slain, g.score())
				return nil
			}
		}
	}
}

func (g *Game) render() {
	var params = maze.RenderParameters{
		Positions: map[position.Position]string{},
		Passage:   passage,
		Wall:      wall,
	}
	for _, p := range g.solution {
		params.Positions[p] = solution
	}
	params.Positions[g.end] = end
	params.Positions[g.theseus] = theseus
	params.Positions[g.minotaur] = minotaur

	g.out.Write(g.maze.Render(params))
}

func (g *Game) score() int {
	return int(time.Since(g.start).Round(time.Second).Seconds())
}

func (g *Game) quit() {
	g.out.Write([]byte(terminal.CursorShow))

	if err := g.terminal.Reset(); err != nil {
		g.log.Error.Printf(
			"failed to reset terminal attributes: %v", err,
		)
	}
}

func (g *Game) escaped() bool {
	return g.theseus == g.end
}

func (g *Game) moveMinotaur() {
	path := g.maze.Solve(g.minotaur, g.theseus)
	if len(path) < 2 {
		return
	}

	g.minotaur = path[1]
}

func (g *Game) moveTheseus(d direction.Direction) {
	if cell := g.maze.Cell(g.theseus); cell.HasDirection(d) {
		g.theseus = g.theseus.WithDirection(d)
	}
}

func (g *Game) slain() bool {
	return g.minotaur == g.theseus
}

func (g *Game) solve() {
	g.solution = g.maze.Solve(g.theseus, g.end)
}
