package main

import (
	"os"
	"path/filepath"

	"code.dwrz.net/src/pkg/editor"
	"code.dwrz.net/src/pkg/log"
	"code.dwrz.net/src/pkg/terminal"
)

func main() {
	var l = log.New(os.Stderr)

	// Setup workspace and log file.
	cdir, err := os.UserCacheDir()
	if err != nil {
		l.Error.Fatalf(
			"failed to determine user cache directory: %v", err,
		)
	}
	wdir := filepath.Join(cdir, "wen")

	if err := os.MkdirAll(wdir, os.ModeDir|0700); err != nil {
		l.Error.Fatalf("failed to create tmp dir: %v", err)
	}

	f, err := os.Create(wdir + "/log")
	if err != nil {
		l.Error.Fatalf("failed to create log file: %v", err)
	}
	defer f.Close()

	// Retrieve terminal info.
	t, err := terminal.New(os.Stdin.Fd())
	if err != nil {
		l.Error.Fatalf("failed to get terminal attributes: %v", err)
	}

	// Create the editor.
	editor, err := editor.New(editor.Parameters{
		In:       os.Stdin,
		Log:      log.New(f),
		Out:      os.Stdout,
		TempDir:  wdir,
		Terminal: t,
	})
	if err != nil {
		l.Error.Fatalf("failed to create editor: %v", err)
	}

	// TODO: handle signals, sigwinch.

	// Run the editor.
	if err := editor.Run(os.Args[1:]); err != nil {
		l.Error.Fatal(err)
	}
}
