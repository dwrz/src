package main

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"code.dwrz.net/src/pkg/editor"
	"code.dwrz.net/src/pkg/log"
	"code.dwrz.net/src/pkg/terminal"
)

func main() {
	var l = log.New(os.Stderr)

	// Setup the main context.
	ctx, cancel := context.WithCancel(context.Background())

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

	// TODO: refactor; handle sigwinch.
	go func() {
		var signals = make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

		// Block until we receive a signal.
		s := <-signals
		l.Debug.Printf("received signal: %s", s)

		cancel()
	}()

	// Run the editor.
	if err := editor.Run(ctx, os.Args[1:]); err != nil {
		l.Error.Fatal(err)
	}
}
