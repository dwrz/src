package command

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"code.dwrz.net/src/pkg/dqs/command/help"
	"code.dwrz.net/src/pkg/dqs/entry"
	"code.dwrz.net/src/pkg/dqs/store"
)

var Note = &command{
	execute: setNote,

	description: "set a note on an entry",
	help:        help.Note,
	name:        "note",
}

func setNote(args []string, date time.Time, store *store.Store) error {
	u, err := store.GetUser()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to get user: %w", err)
	}
	if u == nil {
		return Config.execute(args, date, store)
	}

	e, err := store.GetEntry(date.Format(entry.DateFormat))
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to get entry: %w", err)
	}
	if e == nil {
		e = entry.New(date, u)
	}

	if len(args) == 0 {
		fmt.Println(e.Note)

		return nil
	}

	update, args := args[0], args[1:]
	switch update {
	case "append":
		if len(args) >= 1 {
			if len(e.Note) == 0 {
				e.Note = strings.Join(args, " ")
			} else {
				e.Note = fmt.Sprintf(
					"%s\n%s", e.Note,
					strings.Join(args, " "),
				)
			}
		}

	case "delete":
		e.Note = ""

	case "edit":
		var err error
		e.Note, err = editNote(e.Note)
		if err != nil {
			return err
		}

	case "set":
		if len(args) >= 1 {
			e.Note = strings.Join(args, " ")
		}

	default:
		fmt.Println(e.Note)
	}

	if err := store.UpdateEntry(e); err != nil {
		return fmt.Errorf("failed to store entry: %w", err)
	}

	return nil
}

func editNote(note string) (string, error) {
	// If no note is specified, take input from the user's editor.
	// Write the entry to the temporary file.
	temp, err := os.CreateTemp(os.TempDir(), "dqs-*")
	if err != nil {
		return "", err
	}
	defer temp.Close()

	_, err = temp.Write([]byte(note))
	if err != nil {
		return "", err
	}

	editor := os.Getenv("EDITOR")

	args := strings.Split(editor, " ")
	args = append(args, temp.Name())

	path, err := exec.LookPath(args[0])
	if err != nil {
		return "", fmt.Errorf(
			"failed to find $EDITOR %s: %w", editor, err,
		)
	}

	cmd := exec.Cmd{
		Path:   path,
		Args:   args,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	if err := cmd.Run(); err != nil {
		return "", err
	}

	if _, err := temp.Seek(0, 0); err != nil {
		return "", err
	}

	data, err := io.ReadAll(temp)
	if err != nil {
		return "", err
	}

	if err := os.Remove(temp.Name()); err != nil {
		return "", err
	}

	return string(data), nil
}
