package command

import (
	"errors"
	"fmt"
	"os"
	"time"

	"code.dwrz.net/src/pkg/dqs/command/help"
	"code.dwrz.net/src/pkg/dqs/entry"
	"code.dwrz.net/src/pkg/dqs/portion"
	"code.dwrz.net/src/pkg/dqs/store"
)

var Add = &command{
	execute: addPortions,

	description: "add portions to an entry",
	help:        help.Add,
	name:        "add",
}

func addPortions(args []string, date time.Time, store *store.Store) error {
	u, err := store.GetUser()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to get user: %w", err)
	}
	if u == nil {
		return Config.execute(args, date, store)
	}

	e, err := store.GetEntry(date.Format(entry.DateFormat))
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to get entry: %w", err)
	}
	if e == nil {
		e = entry.New(date, u)
	}

	if len(args) < 2 {
		return fmt.Errorf("missing category and quantity")
	}
	if len(args)%2 != 0 {
		return fmt.Errorf("uneven number of arguments")
	}

	for i := 0; i < len(args); i += 2 {
		portionCategory := args[i]
		quantity := args[i+1]

		c, err := e.Category(portionCategory)
		if err != nil {
			return err
		}

		q, err := portion.Parse(quantity)
		if err != nil {
			return err
		}

		if err := c.Add(q); err != nil {
			return fmt.Errorf(
				"failed to add portions: %w", err,
			)
		}
	}

	if err := store.UpdateEntry(e); err != nil {
		return fmt.Errorf("failed to store entry: %w", err)
	}

	fmt.Println(e.FormatPrint(u))

	return nil
}
