package line

import (
	"strings"
	"unicode/utf8"

	"code.dwrz.net/src/pkg/editor/buffer/glyph"
)

type Line struct {
	data string
}

func New(s string) *Line {
	return &Line{data: s}
}

func (l *Line) DecodeRune(index int) (r rune, size int) {
	return utf8.DecodeRuneInString(l.data[index:])
}

// GlyphIndex attempts to find the preceding index for a target glyph.
func (l *Line) FindGlyphIndex(target int) (index, g int) {
	var li, lg int
	for i, r := range l.data {
		// We've reached the target.
		if g == target {
			return i, g
		}
		// We've gone too far.
		// Return the preceding index and glyph.
		if g > target {
			return li, lg
		}

		// Otherwise, we haven't reached the target.
		// Save this index and glyph.
		li = i
		lg = g

		// Then increment the glyph.
		g += glyph.Width(r)
	}

	// We weren't able to find the glyph.
	// Return the last possible index and glyph.
	return len(l.data), g
}

func (l *Line) Length() int {
	return len(l.data)
}

func (l *Line) Render(offset, width int) string {
	var text = strings.ReplaceAll(l.data, "\t", "        ")
	if offset < 0 || offset > len(text) {
		return ""
	}

	var str strings.Builder
	for _, r := range text {
		rw := glyph.Width(r)
		if offset > 0 {
			offset -= rw
			continue
		}
		if r == utf8.RuneError {
			continue
		}

		// Exhausted column zero.
		if width-rw < -1 {
			break
		}

		str.WriteRune(r)
		width -= rw
	}

	return str.String()
}

func (l *Line) Runes() []rune {
	return []rune(l.data)
}

func (l *Line) String() string {
	return l.data
}

func (l *Line) Width() (count int) {
	for _, r := range l.data {
		rw := glyph.Width(r)
		if r == utf8.RuneError {
			continue
		}
		count += rw
	}

	return count
}
