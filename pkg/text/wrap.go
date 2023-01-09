package text

import "strings"

func Wrap(text string, width int) string {
	if width < 0 {
		return text
	}

	var words = strings.Fields(text)
	if len(words) == 0 {
		return ""
	}

	var (
		wrapped   = words[0]
		remaining = width - len(wrapped)
	)
	for _, word := range words[1:] {
		if len(word)+1 > remaining {
			wrapped += "\n" + word
			remaining = width - len(word)
		} else {
			wrapped += " " + word
			remaining -= 1 + len(word)
		}
	}

	return wrapped
}
