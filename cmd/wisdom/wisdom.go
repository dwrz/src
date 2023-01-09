package main

import (
	"flag"
	"os"
	"path/filepath"

	"code.dwrz.net/src/pkg/log"
	"code.dwrz.net/src/pkg/wisdom"
)

var (
	wrap = flag.Int("w", 72, "word wrap; disabled if negative")
)

func main() {
	var l = log.New(os.Stderr)

	flag.Parse()

	dir, err := os.UserConfigDir()
	if err != nil {
		l.Error.Fatalf("failed to get user config dir: %v", err)
	}

	w, err := wisdom.New(wisdom.Parameters{
		Log:  l,
		Path: filepath.Join(dir, "wisdom"),
		Wrap: *wrap,
	})

	if err := w.Command(flag.CommandLine.Args()); err != nil {
		l.Error.Fatal(err)
	}
}
