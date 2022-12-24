package help

import "strings"

const note = `The note command is used to set a note on an entry.

There are four subcommands available:

append — used to add text to a note.

delete — used to delete a note.

edit — used to edit a note with a text editor (specified by the EDITOR
environment variable).

set — used to set (or overwrite) a note.

If not subcommand is specified, the entry's note is displayed.

Examples:
dqs note append Hello World
dqs -date 20111111 note delete
dqs note edit
dqs note set This is a note.
`

func Note() string {
	var str strings.Builder

	str.WriteString(note)

	return str.String()
}
