package glyph

import "github.com/mattn/go-runewidth"

func Width(r rune) int {
	switch r {
	case '\t':
		return 8
	default:
		return runewidth.RuneWidth(r)
	}
}
