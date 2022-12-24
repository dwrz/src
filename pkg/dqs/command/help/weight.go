package help

import "strings"

const weight = `The weight command sets your weight on an entry.

It accepts a single argument, a number which represents your weight.
The number may have a decimal component.

If the date used is the current date, your user data will also be updated.
This will default future entries to the inputted weight.

Example:
dqs weight 100.00
`

func Weight() string {
	var str strings.Builder

	return str.String()
}
