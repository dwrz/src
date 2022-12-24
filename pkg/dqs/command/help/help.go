package help

import "strings"

const help = `dqs is designed to help you keep track of your Diet
Quality Score, as well as your progress on reaching a target weight.

The typical use case is to use the *add* command to keep track of
portions consumed. By default the application uses the current day, but
if you need to catch up on a past date, the -d or -date flags can be
used to refer to a past date or entry. For example, if you want to add
two fruit portions for today, you can run:

dqs add fruit 2

To do the same for a past date -- for example, August 8, 2008 -- you
can run:

dqs -date 20080808 fruit 2

or

dqs -d 20080808 fruit 2

Refer to dqs help add for more information on the *add* command.

Weight and body fat may be tracked with the body-fat and weight
commands.

To check on progress towards goals, use the dqs report command.
`

func Help() string {
	var str strings.Builder

	str.WriteString(help)

	return str.String()
}
