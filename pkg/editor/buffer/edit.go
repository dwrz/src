package buffer

import (
	"unicode"
	"unicode/utf8"

	"code.dwrz.net/src/pkg/editor/buffer/glyph"
	"code.dwrz.net/src/pkg/editor/buffer/line"
)

func (b *Buffer) Backspace() {
	var cl = b.CursorLine()

	switch {
	// Don't do anything at the beginning of the buffer.
	case b.cursor.line == 0 && b.cursor.index == 0:

	// Delete empty lines.
	case cl.Length() == 0:
		b.lines = append(
			b.lines[:b.cursor.line],
			b.lines[b.cursor.line+1:]...,
		)

		b.CursorUp()
		b.CursorEnd()

	// Append to the previous line.
	case b.cursor.line != 0 && b.cursor.index == 0:
		index := b.cursor.line

		b.CursorUp()
		b.CursorEnd()

		b.lines[index-1].Append(cl.String())

		b.lines = append(
			b.lines[:index],
			b.lines[index+1:]...,
		)

	// Delete a rune.
	default:
		b.CursorLeft()
		b.CursorLine().DeleteRune(b.cursor.index)
	}

}

func (b *Buffer) Insert(r rune) {
	switch {
	case r == '\n' || r == '\r':
		b.Newline()

	case r == '\t':
		b.CursorLine().Insert(b.cursor.index, r)
		b.cursor.index += utf8.RuneLen(r)
		b.cursor.glyph += glyph.Width(r)

	// Ignore all other non-printable characters.
	case !unicode.IsPrint(r):
		return

	default:
		b.CursorLine().Insert(b.cursor.index, r)
		b.cursor.index += utf8.RuneLen(r)
		b.cursor.glyph += glyph.Width(r)
	}
}

// At the start of a line, create a new line above.
// At the end of the line, create a new line below.
// In the middle of a line, any remaining runes go to the next line.
// TODO: using the cursor index will probably break with marks.
func (b *Buffer) Newline() {
	var cl = b.CursorLine()

	switch {
	case b.cursor.index == 0:
		b.lines = append(
			b.lines[:b.cursor.line+1],
			b.lines[b.cursor.line:]...,
		)
		b.lines[b.cursor.line] = line.New("")

	default:
		text := cl.String()
		rest := line.New(text[b.cursor.index:])

		b.lines = append(
			b.lines[:b.cursor.line+1],
			b.lines[b.cursor.line:]...,
		)
		b.lines[b.cursor.line] = line.New(text[:b.cursor.index])
		b.lines[b.cursor.line+1] = rest
	}

	b.CursorDown()
	b.CursorHome()
}
