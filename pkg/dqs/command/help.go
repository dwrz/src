package command

import (
	"fmt"
	"strings"
	"time"

	"code.dwrz.net/src/pkg/dqs/command/help"
	"code.dwrz.net/src/pkg/dqs/store"
)

const (
	helpIntro = `dqs — diet quality score

Usage:
	dqs [flags] <command> [arguments]


Run dqs -help or dqs -h to see information on available flags.

The following commands are available:

`
	helpMore = "Run dqs help <command> to see the command's documentation."
)

var Help = &command{
	execute: showHelp,

	description: "show built-in command documentation",
	help:        help.Help,
	name:        "help",
}

func showHelp(args []string, date time.Time, store *store.Store) error {
	if len(args) == 0 {
		var str strings.Builder

		str.WriteString(helpIntro)

		for _, c := range commands {
			str.WriteString(
				fmt.Sprintf("%-11s%s\n", c.name, c.description),
			)
		}

		str.WriteString("\n")
		str.WriteString(helpMore)

		fmt.Println(str.String())

		return nil
	}

	cmd, err := Match(args[0])
	if err != nil {
		return err
	}

	fmt.Println(cmd.help())

	return nil
}
