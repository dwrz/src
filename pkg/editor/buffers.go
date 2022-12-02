package editor

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"code.dwrz.net/src/pkg/editor/buffer"
)

// load files into editor buffers.
func (e *Editor) load(files []string) {
	// Attempt to deduplicate any files.
	var unique = map[string]struct{}{}
	for _, f := range files {
		path, err := filepath.Abs(f)
		if err != nil {
			e.log.Error.Printf(
				"failed to get absolute path for %s: %v",
				f, err,
			)

			path = filepath.Clean(f)
		}

		unique[path] = struct{}{}
	}

	// Load the files.
	// Set the first successfully loaded file as the active buffer.
	var setActive bool
	for name := range unique {
		p := buffer.NewBufferParams{
			Name: name,
			Log:  e.log,
		}

		b, err := buffer.Open(p)
		// Create the file if it doesn't exist.
		if errors.Is(err, os.ErrNotExist) {
			b, err = buffer.Create(p)
		}
		// If there was an error, report it to the user.
		if err != nil {
			e.messages <- Message{
				Text: fmt.Sprintf(
					"failed to load buffer %s: %v",
					name, err,
				),
			}
			continue
		}

		e.log.Debug.Printf("loaded buffer %s", name)

		e.setBuffer(b)

		if !setActive {
			e.setActiveBuffer(name)
		}
	}

	e.messages <- Message{Text: "loaded buffers"}
}

// setBuffer stores a buffer in the editor's buffer map.
func (e *Editor) setBuffer(b *buffer.Buffer) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.buffers[b.Name()] = b
}

// setActiveBuffer sets the named buffer as the active buffer.
func (e *Editor) setActiveBuffer(name string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	b, exists := e.buffers[name]
	if !exists {
		e.errs <- fmt.Errorf(
			"failed to set active buffer: buffer %s does not exist",
			name,
		)
	}

	e.active = b
	e.log.Debug.Printf("set active buffer %s", b.Name())
}
