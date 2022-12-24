package command

import (
	"fmt"
	"time"

	"code.dwrz.net/src/pkg/dqs/command/help"
	"code.dwrz.net/src/pkg/dqs/report"
	"code.dwrz.net/src/pkg/dqs/store"
)

var Report = &command{
	execute: func(args []string, date time.Time, store *store.Store) error {
		entries, err := store.GetAllEntries()
		if err != nil {
			return err
		}

		fmt.Println(report.New(entries).Format())

		return nil
	},

	description: "report user and entry statistics",
	help:        help.Report,
	name:        "report",
}
