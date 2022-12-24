package command

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"code.dwrz.net/src/pkg/dqs/command/help"
	"code.dwrz.net/src/pkg/dqs/entry"
	"code.dwrz.net/src/pkg/dqs/store"
)

var BodyFat = &command{
	execute: setBodyFat,

	description: "set the user's body fat percentage on an entry",
	help:        help.BodyFat,
	name:        "body-fat",
}

func setBodyFat(args []string, date time.Time, store *store.Store) error {
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

	if len(args) == 0 {
		return fmt.Errorf("missing body fat percentage")
	}

	bf, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return err
	}

	e.BodyFat = bf

	if err := store.UpdateEntry(e); err != nil {
		return fmt.Errorf("failed to store entry: %w", err)
	}

	// If the entry is for today, update the user.
	var currentDate = time.Now().Format(entry.DateFormat)
	if currentDate == date.Format(entry.DateFormat) {
		u.BodyFat = bf

		if err := store.UpdateUser(u); err != nil {
			return fmt.Errorf("failed to store user: %w", err)
		}
	}

	fmt.Println(u.FormatPrint())

	return nil
}
