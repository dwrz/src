package help

import "strings"

const export = `The export command prints statistics from all entries.

The columns are date, body-fat, diet quality score, and weight.

The export format is Comma Separate Values (CSV).

Example:
dqs export
`

func Export() string {
	var str strings.Builder

	str.WriteString(export)

	return str.String()
}
