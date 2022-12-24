package command

import (
	"fmt"
	"time"

	"code.dwrz.net/src/pkg/dqs/store"
)

type command struct {
	execute func(args []string, date time.Time, store *store.Store) error

	description string
	help        func() string
	name        string
}

func (c command) Execute(p Parameters) error {
	if err := p.Validate(); err != nil {
		return err
	}

	return c.execute(p.Args, p.Date, p.Store)
}

var commands []*command

func init() {
	commands = []*command{
		Add,
		BodyFat,
		Config,
		Delete,
		Entry,
		Export,
		Help,
		Note,
		Remove,
		Report,
		User,
		Weight,
	}
}

func Match(name string) (*command, error) {
	for _, command := range commands {
		if name == command.name {
			return command, nil
		}
	}

	return nil, fmt.Errorf("unrecognized command: %s", name)
}
