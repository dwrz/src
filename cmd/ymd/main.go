package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	now = time.Now()

	year = flag.Int(
		"y", 0,
		"year; defaults to current year",
	)
	month = flag.Int(
		"m", 0,
		"month; defaults to current month",
	)

	ym = flag.Bool("ym", false, "create YY-MM directories")
)

// TODO: path arguments.
func main() {
	flag.Parse()

	if *year < 0 {
		fmt.Fprintf(os.Stderr, "invalid year: %d\n", *year)
		return
	}
	if *year == 0 {
		y := now.Year()
		year = &y
	}

	// Create YY-MM directories.
	if *ym {
		mkdirYM(*year)
		return
	}

	// Otherwise, create YYYY-MM-DD directories.
	if *month < 0 || *month > 12 {
		fmt.Fprintf(os.Stderr, "invalid month: %d\n", month)
		return
	}
	if *month == 0 {
		m := int(now.Month())
		month = &m
	}

	mkdirYMD(*year, *month)
}

func mkdirYM(year int) {
	for m := 1; m <= 12; m++ {
		name := fmt.Sprintf("%d-%02d", year, m)
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}
}

func mkdirYMD(year, month int) {
	// Calculate the last day of the month in the input year.
	last := time.Date(
		year, time.Month(month), 1, 0, 0, 0, 0, time.Local,
	).AddDate(0, 1, 0).Add(-time.Nanosecond).Day()

	// Create directories for each day of the month.
	for d := 1; d <= last; d++ {
		name := fmt.Sprintf("%d-%02d-%02d", year, month, d)
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}
}
