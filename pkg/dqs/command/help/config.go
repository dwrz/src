package help

import "strings"

const config = `The config command is used to configure dqs.

It allows for entry of the following fields:

Name
Birthday
Units (imperial or metric*)
Height
Weight
Body Fat Percentage
Diet (omnivore, vegan, or vegetarian*)
Target Weight

All fields are optional. Units default to metric; the default diet is
vegetarian.

This command will be invoked automatically the first time dqs is run.

Example:
dqs config
`

func Config() string {
	var str strings.Builder

	str.WriteString(config)

	return str.String()
}
