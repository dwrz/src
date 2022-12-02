package buffer

type Output struct {
	Glyph int
	Line  int
	Lines []string
}

func (b *Buffer) Render(height, width int) *Output {
	// Vertical scroll.
	if b.cursor.line < b.offset.line {
		b.offset.line = b.cursor.line
	}
	if b.cursor.line > height+b.offset.line {
		b.offset.line = b.cursor.line - height
	}

	// Horizontal scroll.
	if b.cursor.glyph < b.offset.glyph {
		b.offset.glyph = b.cursor.glyph
	}
	if b.cursor.glyph > width+b.offset.glyph {
		b.offset.glyph = b.cursor.glyph - width
	}

	// Generate lines.
	var lines = []string{}
	for i := b.offset.line; i <= height+b.offset.line; i++ {
		// Return empty lines for indexes past the buffer's lines.
		if i >= len(b.lines) {
			lines = append(lines, "")
			continue
		}

		line := b.lines[i].Render(b.offset.glyph, width)
		lines = append(lines, line)
	}

	return &Output{
		// Terminals are 1-indexed.
		Glyph: b.cursor.glyph - b.offset.glyph + 1,
		Line:  b.cursor.line - b.offset.line + 1,
		Lines: lines,
	}
}
