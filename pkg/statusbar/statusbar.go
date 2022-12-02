package statusbar

import (
	"fmt"
	"strings"
	"sync"

	"code.dwrz.net/src/pkg/log"
)

type Block interface {
	Name() string
	Render() (string, error)
}

type Parameters struct {
	Blocks    []Block
	Log       *log.Logger
	Separator string
}

type StatusBar struct {
	b      *strings.Builder
	blocks []Block
	l      *log.Logger
	sep    string
}

func (s *StatusBar) Render() string {
	defer s.b.Reset()

	fmt.Fprintf(s.b, "%s ", s.sep)

	var (
		outputs = make([]string, len(s.blocks))
		wg      sync.WaitGroup
	)

	wg.Add(len(s.blocks))

	for i, b := range s.blocks {
		go func(i int, b Block) {
			defer wg.Done()

			text, err := b.Render()
			if err != nil {
				s.l.Error.Printf(
					"failed to render %s: %v",
					b.Name(), err,
				)
				outputs[i] = ""
			} else {
				outputs[i] = text
			}
		}(i, b)
	}

	wg.Wait()

	for i, o := range outputs {
		s.b.WriteString(o)

		if i < len(outputs)-1 {
			fmt.Fprintf(s.b, " %s ", s.sep)
		}
	}

	return s.b.String()
}

func New(p Parameters) *StatusBar {
	return &StatusBar{
		b:      &strings.Builder{},
		blocks: p.Blocks,
		l:      p.Log,
		sep:    p.Separator,
	}
}
