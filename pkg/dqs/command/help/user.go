package help

import "strings"

const user = `The user command displays your user data.

This command may be used to confirm settings -- such as diet -- or the
user's current measurements -- such as height and weight.

Example:
dqs user
`

func User() string {
	var str strings.Builder

	str.WriteString(user)

	return str.String()
}
