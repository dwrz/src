package buffer

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"code.dwrz.net/src/pkg/editor/buffer/line"
	"code.dwrz.net/src/pkg/log"
)

type Buffer struct {
	cursor *Cursor
	lines  []*line.Line
	log    *log.Logger
	file   *os.File
	name   string
	offset *offset
	stat   os.FileInfo
	saved  time.Time
}

// Offset is the first visible column and row of the buffer.
type offset struct {
	glyph int
	line  int
}

// TODO: return true if the underlying file has changed.
// Use file size and mod time.
// Need to distinguish underlying file changes from changes made via edits.
func (b *Buffer) Changed() {
	return
}

func (b *Buffer) Close() error {
	return b.file.Close()
}

func (b *Buffer) CursorLine() *line.Line {
	return b.lines[b.cursor.line]
}

func (b *Buffer) Line(i int) *line.Line {
	return b.lines[i]
}

func (b *Buffer) Name() string {
	return b.name
}

type NewBufferParams struct {
	Name string
	Log  *log.Logger
}

func Create(p NewBufferParams) (*Buffer, error) {
	var b = &Buffer{
		cursor: &Cursor{},
		lines:  []*line.Line{},
		log:    p.Log,
		name:   p.Name,
		offset: &offset{},
	}

	// Create the file.
	f, err := os.Create(b.name)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	b.file = f

	// Scan the lines.
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		b.lines = append(b.lines, line.New(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan: %v", err)
	}

	// Add a line to write in, if the file is empty.
	if len(b.lines) == 0 {
		b.lines = append(b.lines, &line.Line{})
	}

	return b, nil
}

func Open(p NewBufferParams) (*Buffer, error) {
	var b = &Buffer{
		cursor: &Cursor{},
		lines:  []*line.Line{},
		log:    p.Log,
		name:   p.Name,
		offset: &offset{},
	}

	// Open the file.
	f, err := os.Open(b.name)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	b.file = f

	// Scan the lines.
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		b.lines = append(b.lines, line.New(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan: %v", err)
	}

	// Add a line to write in, if the file is empty.
	if len(b.lines) == 0 {
		b.lines = append(b.lines, &line.Line{})
	}

	return b, nil
}
