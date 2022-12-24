package help

import "strings"

const bodyFat = `The body-fat command sets the body fat percentage on an entry.

It accepts a single argument, a number which represents the body fat
percentage. The number may have a decimal component.

If the date used is the current date, your user data will also be
updated to use the body fat percentage. This will default future entries
to the inputted percentage.

Example:
dqs body-fat 24.25
`

func BodyFat() string {
	var str strings.Builder

	str.WriteString(bodyFat)

	return str.String()
}
