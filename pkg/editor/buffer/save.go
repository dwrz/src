package buffer

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

func (b *Buffer) Save() error {
	var buf bytes.Buffer
	switch len(b.lines) {
	case 1:
		if l := b.lines[0]; l.Length() != 0 {
			if _, err := buf.WriteString(l.String()); err != nil {
				return fmt.Errorf(
					"failed to write to buffer: %v", err,
				)
			}
			if _, err := buf.WriteRune('\n'); err != nil {
				return fmt.Errorf(
					"failed to write to buffer: %v", err,
				)
			}
		}
	default:
		for _, l := range b.lines {
			if _, err := buf.WriteString(l.String()); err != nil {
				return fmt.Errorf(
					"failed to write to buffer: %v", err,
				)
			}
			if _, err := buf.WriteRune('\n'); err != nil {
				return fmt.Errorf(
					"failed to write to buffer: %v", err,
				)
			}
		}
	}

	if err := os.WriteFile(
		b.name,
		buf.Bytes(),
		os.ModePerm,
	); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	b.saved = time.Now()

	return nil
}
