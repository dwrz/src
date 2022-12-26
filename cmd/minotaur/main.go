package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"code.dwrz.net/src/pkg/log"
	"code.dwrz.net/src/pkg/minotaur"
	"code.dwrz.net/src/pkg/terminal"
)

var (
	height = flag.Int("h", 0, "height")
	width  = flag.Int("w", 0, "width")
	tick   = flag.Int("t", 250, "ms between minotaur movement")
)

func main() {
	var l = log.New(os.Stderr)

	// Parse flags.
	flag.Parse()
	if *height < 0 {
		l.Error.Fatalf("invalid height: %d", *height)
	}
	if *width < 0 {
		l.Error.Fatalf("invalid width: %d", *width)
	}
	if *tick <= 0 {
		l.Error.Fatalf("invalid tick: %d", *tick)
	}

	// Setup the main context.
	ctx, cancel := context.WithCancel(context.Background())

	// Setup workspace and log file.
	cdir, err := os.UserCacheDir()
	if err != nil {
		l.Error.Fatalf(
			"failed to determine user cache directory: %v", err,
		)
	}
	wdir := filepath.Join(cdir, "minotaur")

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
	size, err := t.Size()
	if err != nil {
		l.Error.Fatalf("failed to get terminal size: %v", err)
	}
	// If no dimensions were provided, use the terminal size.
	if *height == 0 {
		rows := int(size.Rows)
		height = &rows
	}
	if *width == 0 {
		cols := int(size.Columns)
		width = &cols
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

	// Setup the game.
	game, err := minotaur.New(minotaur.Parameters{
		Height:   *height,
		In:       os.Stdin,
		Log:      log.New(f),
		Out:      os.Stdout,
		Terminal: t,
		Tick:     time.Duration(*tick) * time.Millisecond,
		Width:    *width,
	})
	if err != nil {
		l.Error.Fatalf("failed to create game: %v", err)
	}

	// Run the game.
	if err := game.Run(ctx); err != nil {
		l.Error.Fatal(err)
	}
}
