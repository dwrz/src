package editor

import (
	"fmt"
	"time"

	"code.dwrz.net/src/pkg/build"
	"code.dwrz.net/src/pkg/terminal"
)

func (e *Editor) Run(files []string) error {
	e.terminal.SetRaw()
	e.out.Write([]byte(terminal.ClearScreen))

	// Log build info.
	e.log.Debug.Printf(
		"文 version %s; built at %s on %s",
		build.Commit, build.Time, build.Hostname,
	)

	// Reset the terminal before exiting.
	defer e.quit()

	// Start reading user input.
	go e.readInput()

	// Open all the files.
	go e.load(files)

	// Set the initial message.
	go func() {
		time.Sleep(1 * time.Second)
		e.messages <- Message{Text: "Ctrl-Q: Quit"}
	}()

	// Main loop.
	for {
		select {
		case err := <-e.errs:
			return err

		case msg := <-e.messages:
			e.log.Debug.Printf("%s", msg.Text)
			if err := e.render(&msg); err != nil {
				return fmt.Errorf("failed to render: %w", err)
			}

		case input := <-e.input:
			if input.Command == Quit {
				return nil
			}
			if err := e.processInput(input); err != nil {
				return fmt.Errorf(
					"failed to process input: %w", err,
				)
			}
			if err := e.render(nil); err != nil {
				return fmt.Errorf("failed to render: %w", err)
			}
		}
	}
}
