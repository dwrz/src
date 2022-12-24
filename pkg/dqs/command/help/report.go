package help

import "strings"

const report = `The report command displays a report on your statistics
and usage.

The command will display the statistics on the following:

Entries
Body Fat
DQS
Weight

It will also display recommendations as to which food categories you
should eat more or less of, given past entries.
`

func Report() string {
	var str strings.Builder

	str.WriteString(report)

	return str.String()
}
