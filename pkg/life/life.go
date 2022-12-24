package life

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"code.dwrz.net/src/pkg/color"
	"code.dwrz.net/src/pkg/log"
	"code.dwrz.net/src/pkg/terminal"
	"code.dwrz.net/src/pkg/terminal/input"
)

const (
	alive = color.BackgroundBlack + "  " + color.Reset
	born  = color.BackgroundGreen + "  " + color.Reset
	died  = color.BackgroundRed + "  " + color.Reset
	dead  = "  "
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

type Life struct {
	errs     chan error
	events   chan *input.Event
	height   int
	in       *os.File
	input    *input.Reader
	last     [][]bool
	log      *log.Logger
	out      *os.File
	terminal *terminal.Terminal
	tick     time.Duration
	ticker   *time.Ticker
	turn     uint
	width    int
	world    [][]bool
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

func New(p Parameters) *Life {
	// Account for two spaces to represent each cell.
	p.Width /= 2
	// Account one line to render the turn count.
	p.Height -= 1

	var life = &Life{
		errs:     make(chan error),
		events:   make(chan *input.Event),
		height:   p.Height,
		in:       p.In,
		last:     make([][]bool, p.Height),
		log:      p.Log,
		out:      p.Out,
		terminal: p.Terminal,
		tick:     p.Tick,
		ticker:   time.NewTicker(p.Tick),
		width:    p.Width,
		world:    make([][]bool, p.Height),
	}
	for i := range life.world {
		life.last[i] = make([]bool, p.Width)
		life.world[i] = make([]bool, p.Width)
	}
	for i := 0; i < (p.Height * p.Width / 4); i++ {
		x := random.Intn(p.Width)
		y := random.Intn(p.Height)

		life.world[y][x] = true
	}

	life.input = input.New(input.Parameters{
		Chan: life.events,
		In:   os.Stdin,
		Log:  life.log,
	})

	return life
}

func (l *Life) Run(ctx context.Context) error {
	// Setup the terminal.
	if err := l.terminal.SetRaw(); err != nil {
		return fmt.Errorf(
			"failed to set terminal attributes: %v", err,
		)
	}

	// Clean up before exit.
	defer l.quit()

	// TODO: refactor to use a terminal output package.
	l.out.WriteString(terminal.ClearScreen)
	l.out.WriteString(terminal.CursorHide)
	l.render()

	// Start reading user input.
	go func() {
		if err := l.input.Run(ctx); err != nil {
			l.errs <- err
		}
	}()

	// Main loop.
	for {
		select {
		case <-ctx.Done():
			return nil

		case <-l.ticker.C:
			l.step()
			l.render()

		case event := <-l.events:
			switch event.Rune {
			case 'q': // Quit
				return nil
			}
		}
	}
}

func (l *Life) alive(x, y int) bool {
	// Wrap coordinates toroidally.
	x += l.width
	x %= l.width
	y += l.height
	y %= l.height

	return l.world[y][x]
}

func (l *Life) born(x, y int) bool {
	// Wrap coordinates toroidally.
	x += l.width
	x %= l.width
	y += l.height
	y %= l.height

	return !l.last[y][x] && l.world[y][x]
}

func (l *Life) died(x, y int) bool {
	// Wrap coordinates toroidally.
	x += l.width
	x %= l.width
	y += l.height
	y %= l.height

	return l.last[y][x] && !l.world[y][x]
}

func (l *Life) next(x, y int) bool {
	// Count the numbers of living neighbors.
	var alive = 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			// Ignore self.
			if h == 0 && v == 0 {
				continue
			}
			if l.alive(x+v, y+h) {
				alive++
			}
		}
	}

	// Determine the next state.
	switch alive {
	case 3: // Turn on.
		return true
	case 2: // Maintain state.
		return l.alive(x, y)
	default: // Turn off.
		return false
	}
}

func (l *Life) quit() {
	l.out.Write([]byte(terminal.CursorShow))

	if err := l.terminal.Reset(); err != nil {
		l.log.Error.Printf(
			"failed to reset terminal attributes: %v", err,
		)
	}
}

func (l *Life) render() {
	var buf bytes.Buffer

	buf.WriteString(terminal.CursorTopLeft)

	for y := 0; y < l.height; y++ {
		for x := 0; x < l.width; x++ {
			switch {
			case l.born(x, y):
				buf.WriteString(born)
			case l.died(x, y):
				buf.WriteString(died)
			case l.alive(x, y):
				buf.WriteString(alive)
			default:
				buf.WriteString(dead)

			}
		}
		buf.WriteString("\r\n")
	}
	fmt.Fprintf(
		&buf, "%sTurn:%s %d\r\n",
		color.Bold, color.Reset, l.turn,
	)

	l.out.Write(buf.Bytes())
}

func (l *Life) step() {
	// Create the next world.
	var next = make([][]bool, l.height)
	for i := range next {
		next[i] = make([]bool, l.width)
	}

	// Set the new cells based on the existing cells.
	for y := 0; y < l.height; y++ {
		for x := 0; x < l.width; x++ {
			next[y][x] = l.next(x, y)
		}
	}

	// Increment the turn and set the next world.
	l.turn++
	l.last = l.world
	l.world = next
}
