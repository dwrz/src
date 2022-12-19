package editor

import (
	"context"
	"fmt"

	"code.dwrz.net/src/pkg/build"
	"code.dwrz.net/src/pkg/terminal/input"
)

func (e *Editor) Run(ctx context.Context, files []string) error {
	if err := e.terminal.SetRaw(); err != nil {
		return fmt.Errorf("failed to set raw mode: %w", err)
	}
	e.canvas.Reset()

	// Log build info.
	e.log.Debug.Printf(
		"文 version %s; built at %s on %s",
		build.Commit, build.Time, build.Hostname,
	)

	// Reset the terminal before exiting.
	defer e.quit()

	// Start reading user input.
	go func() {
		if err := e.reader.Run(ctx); err != nil {
			e.errs <- err
		}
	}()

	// Open the files.
	go e.load(files)

	// Main loop.
	for {
		select {
		case <-ctx.Done():
			e.log.Debug.Printf("context done: stopping")
			return nil

		case err := <-e.errs:
			return err

		case msg := <-e.messages:
			e.log.Debug.Printf("%s", msg.Text)
			if err := e.canvas.Render(e.active, msg); err != nil {
				return fmt.Errorf("failed to render: %w", err)
			}

		case event := <-e.input:
			if event.Rune == 'q'&input.Control {
				return nil
			}
			if err := e.bufferInput(event); err != nil {
				return fmt.Errorf(
					"failed to process input: %w", err,
				)
			}
			if err := e.canvas.Render(e.active, nil); err != nil {
				return fmt.Errorf("failed to render: %w", err)
			}
		}
	}
}
