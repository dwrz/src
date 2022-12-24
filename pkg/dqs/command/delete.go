package command

import (
	"fmt"
	"time"

	"code.dwrz.net/src/pkg/dqs/command/help"
	"code.dwrz.net/src/pkg/dqs/entry"
	"code.dwrz.net/src/pkg/dqs/store"
)

var Delete = &command{
	execute: func(args []string, date time.Time, store *store.Store) error {
		id := date.Format(entry.DateFormat)

		if err := store.DeleteEntry(id); err != nil {
			return fmt.Errorf("failed to delete entry: %w", err)
		}

		return nil
	},

	description: "delete an entry",
	help:        help.Delete,
	name:        "delete",
}
