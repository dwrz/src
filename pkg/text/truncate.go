package text

import "strings"

func Truncate(text string, width int) string {
	if len(text) < width {
		return text
	}

	return strings.TrimSpace(text[:width])
}
