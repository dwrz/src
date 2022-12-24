package help

import "strings"

const entry = `The entry command displays the date's entry.

It is the default command run, if no command is specified.

The entry date is specified via the -date flag (run dqs -help to read
more about the flag).

Example:
dqs entry

The command is equivalent to:
dqs
`

func Entry() string {
	var str strings.Builder

	str.WriteString(entry)

	return str.String()
}
