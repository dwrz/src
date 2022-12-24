package main

import (
	"flag"
	"fmt"
	"os"

	"code.dwrz.net/src/cmd/dqs/config"
	"code.dwrz.net/src/pkg/dqs"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "bad config: %v\n", err)
		return
	}

	app, err := dqs.New(dqs.Config{Dir: cfg.Dir})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if err := app.Run(flag.Args(), cfg.Date); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
