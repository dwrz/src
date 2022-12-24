package help

import "strings"

const deleteEntry = `The delete command deletes an entry.

The entry date is specified via the -date flag (run dqs -help to read
more about the flag).

There are no arguments.

There is no way to undo a deletion.

The command returns an error if no entry exists for the specified date.

Example:
dqs -date 20210620 delete
`

func Delete() string {
	var str strings.Builder

	str.WriteString(deleteEntry)

	return str.String()
}
