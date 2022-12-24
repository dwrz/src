package command

import (
	"errors"
	"fmt"
	"os"
	"time"

	"code.dwrz.net/src/pkg/dqs/command/help"
	"code.dwrz.net/src/pkg/dqs/entry"
	"code.dwrz.net/src/pkg/dqs/store"
)

var Entry = &command{
	execute: func(args []string, date time.Time, store *store.Store) error {
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

		fmt.Println(e.FormatPrint(u))

		return nil
	},

	description: "display a date's entry (default)",
	help:        help.Entry,
	name:        "entry",
}
