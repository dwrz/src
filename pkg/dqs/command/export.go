package command

import (
	"fmt"
	"strings"
	"time"

	"code.dwrz.net/src/pkg/dqs/command/help"
	"code.dwrz.net/src/pkg/dqs/store"
)

var Export = &command{
	execute: func(args []string, date time.Time, store *store.Store) error {
		entries, err := store.GetAllEntries()
		if err != nil {
			return err
		}

		var str strings.Builder
		str.WriteString("date,body-fat,dqs,weight,\n")

		for _, e := range entries {
			str.WriteString(fmt.Sprintf(
				"%s,%.2f,%.1f,%.2f,\n",
				e.Date.Format("20060102"),
				e.BodyFat,
				e.Score(),
				e.Weight,
			))
		}

		fmt.Println(str.String())

		return nil
	},

	description: "export entries to csv",
	help:        help.Export,
	name:        "export",
}
