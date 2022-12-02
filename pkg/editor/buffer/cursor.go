package buffer

import (
	"unicode/utf8"

	"code.dwrz.net/src/pkg/editor/buffer/glyph"
)

type Cursor struct {
	glyph int

	index int
	line  int
}

func (b *Buffer) Cursor() Cursor {
	return *b.cursor
}

func (c Cursor) Glyph() int {
	return c.glyph
}

func (c Cursor) Index() int {
	return c.index
}

func (c Cursor) Line() int {
	return c.line
}

func (b *Buffer) CursorDown() {
	b.log.Debug.Printf("↓ before: %#v", b.cursor)
	defer b.log.Debug.Printf("↓ after: %#v", b.cursor)

	// If we're on the last line, don't move down.
	if b.cursor.line >= len(b.lines)-1 {
		return
	}

	// Move to the next line.
	b.cursor.line++

	// Adjust the column.
	var (
		line   = b.CursorLine()
		length = line.Length()
		width  = line.Width()
	)
	switch {
	// If we've moved to a shorter line, snap to its end.
	case b.cursor.glyph > width:
		b.cursor.index = length
		b.cursor.glyph = width

	default:
		b.cursor.index, b.cursor.glyph = line.FindGlyphIndex(
			b.cursor.glyph,
		)
	}
}

func (b *Buffer) CursorLeft() {
	b.log.Debug.Printf("← before: %#v", b.cursor)
	defer b.log.Debug.Printf("← after: %#v", b.cursor)

	switch {
	// If we're at the beginning of the line, move to the line above;
	// unless we're at the start of the buffer.
	case b.cursor.index == 0 && b.cursor.line > 0:
		b.cursor.line--

		line := b.CursorLine()
		b.cursor.index = line.Length()
		b.cursor.glyph = line.Width()

	// Move left by one rune.
	case b.cursor.index > 0:
		var (
			line = b.CursorLine()
			r    rune
			size int
		)
		// Reverse until we hit the start of a rune.
		for i := b.cursor.index - 1; i >= 0; i-- {
			r, size = line.DecodeRune(i)
			if r != utf8.RuneError {
				b.cursor.index -= size
				b.cursor.glyph -= glyph.Width(r)
				break
			}
		}
	}
}

func (b *Buffer) CursorRight() {
	b.log.Debug.Printf("→ before: %#v", b.cursor)
	defer b.log.Debug.Printf("→ after: %#v", b.cursor)

	var (
		line   = b.CursorLine()
		length = line.Length()
	)

	switch {
	// If we're at the end of the line, move to the line below;
	// unless we're at the end of the buffer.
	case b.cursor.index == length && b.cursor.line < len(b.lines)-1:
		b.cursor.line++
		b.cursor.index = 0
		b.cursor.glyph = 0

	// Move the index right by one rune.
	case b.cursor.index < length:
		r, size := line.DecodeRune(b.cursor.index)
		if r == utf8.RuneError {
			b.cursor.index++
			b.cursor.glyph += glyph.Width(r)
		}

		b.cursor.index += size
		b.cursor.glyph += glyph.Width(r)
	}
}

func (b *Buffer) CursorUp() {
	b.log.Debug.Printf("↑ before: %#v", b.cursor)
	defer b.log.Debug.Printf("↑ after: %#v", b.cursor)

	// If we're on the first line, don't move up.
	if b.cursor.line == 0 {
		return
	}

	b.cursor.line--

	// Adjust the column.
	var (
		line   = b.CursorLine()
		length = line.Length()
		width  = line.Width()
	)
	switch {
	case b.cursor.glyph > width:
		b.cursor.index = length
		b.cursor.glyph = width

	default:
		b.cursor.index, b.cursor.glyph = line.FindGlyphIndex(
			b.cursor.glyph,
		)
	}
}

func (b *Buffer) PageDown(height int) {
	if b.cursor.line+height > len(b.lines) {
		b.cursor.line = len(b.lines) - 1
	} else {
		b.cursor.line += height
	}
}

func (b *Buffer) PageUp(height int) {
	if b.cursor.line-height > 0 {
		b.cursor.line -= height
	} else {
		b.cursor.line = 0
	}
}

func (b *Buffer) CursorHome() {
	b.cursor.glyph = 0
	b.cursor.index = 0
}

func (b *Buffer) CursorEnd() {
	var line = b.CursorLine()

	b.cursor.index = line.Length()
	b.cursor.glyph = line.Width()
}
