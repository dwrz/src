package command

import (
	"errors"
	"fmt"
	"os"
	"time"

	"code.dwrz.net/src/pkg/dqs/command/help"
	"code.dwrz.net/src/pkg/dqs/store"
)

var User = &command{
	execute: func(args []string, date time.Time, store *store.Store) error {
		u, err := store.GetUser()
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("failed to get user: %w", err)
		}
		if u == nil {
			return Config.execute(args, date, store)
		}

		fmt.Println(u.FormatPrint())

		return nil
	},

	description: "display user data and settings",
	help:        help.User,
	name:        "user",
}
