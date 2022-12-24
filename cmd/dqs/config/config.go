package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"code.dwrz.net/src/pkg/dqs/entry"
)

const (
	defaultDir = ""
	dirUsage   = "the path to the dqs directory"

	defaultDate = ""
	dateUsage   = "the entry date to use, formatted as YYYYMMDD"
)

type Config struct {
	Date time.Time
	Dir  string
}

func Get() (*Config, error) {
	var (
		cfg  = &Config{}
		date string
	)

	flag.StringVar(&cfg.Dir, "dir", defaultDir, dirUsage)
	flag.StringVar(&date, "d", defaultDate, dateUsage)
	flag.StringVar(&date, "date", defaultDate, dateUsage)

	flag.Parse()

	// If unset, use the default directory.
	if cfg.Dir == "" {
		dir, err := DefaultDirectory()
		if err != nil {
			return nil, err
		}

		cfg.Dir = dir
	}

	// Parse and set the date.
	// If unset, use the current date.
	switch date {
	case defaultDate:
		now := time.Now()
		cfg.Date = time.Date(
			now.Year(), now.Month(), now.Day(),
			0, 0, 0, 0, time.Local,
		)
	default:
		var err error
		cfg.Date, err = time.Parse(entry.DateFormat, date)
		if err != nil {
			return nil, fmt.Errorf("invalid date: %w", err)
		}
	}

	return cfg, nil
}

func DefaultDirectory() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf(
			"failed to get user config dir: %w", err,
		)
	}

	return fmt.Sprintf("%s/dqs", dir), nil
}
