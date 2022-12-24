package dqs

import (
	"time"

	"code.dwrz.net/src/pkg/dqs/command"
	"code.dwrz.net/src/pkg/dqs/store"
)

type App struct {
	store *store.Store
}

func New(cfg Config) (*App, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	store, err := store.Open(cfg.Dir)
	if err != nil {
		return nil, err
	}

	return &App{
		store: store,
	}, nil
}

func (a *App) Run(args []string, date time.Time) error {
	params := command.Parameters{
		Args:  args,
		Date:  date,
		Store: a.store,
	}

	// If no arguments were provided, print the entry.
	if len(args) == 0 {
		return command.Entry.Execute(params)
	}

	// Process the command.
	params.Args = args[1:]

	cmd, err := command.Match(args[0])
	if err != nil {
		return err
	}

	return cmd.Execute(params)
}
