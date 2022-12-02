package line

import "strings"

func (l *Line) Append(s string) {
	l.data = l.data + s
}

func (l *Line) DeleteRune(index int) (deleted rune) {
	if index < 0 || index >= len(l.data) {
		return
	}

	var str strings.Builder
	for i, r := range l.data {
		if i == index {
			deleted = r
			continue
		}
		str.WriteRune(r)
	}

	l.data = str.String()

	return deleted
}

func (l *Line) Insert(index int, nr rune) {
	if index < 0 || index > len(l.data) {
		return
	}

	var str strings.Builder
	switch {

	// Handle empty lines and the end of the line.
	case index == len(l.data):
		str.WriteString(l.data)
		str.WriteRune(nr)

	default:
		for i, r := range l.data {
			if i == index {
				str.WriteRune(nr)
			}
			str.WriteRune(r)
		}
	}

	l.data = str.String()
}
