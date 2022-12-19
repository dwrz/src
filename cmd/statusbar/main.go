package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"code.dwrz.net/src/pkg/log"
	"code.dwrz.net/src/pkg/statusbar"
	"code.dwrz.net/src/pkg/statusbar/cpu"
	"code.dwrz.net/src/pkg/statusbar/datetime"
	"code.dwrz.net/src/pkg/statusbar/disk"
	"code.dwrz.net/src/pkg/statusbar/eth"
	"code.dwrz.net/src/pkg/statusbar/light"
	"code.dwrz.net/src/pkg/statusbar/memory"
	"code.dwrz.net/src/pkg/statusbar/mic"
	"code.dwrz.net/src/pkg/statusbar/power"
	"code.dwrz.net/src/pkg/statusbar/temp"
	"code.dwrz.net/src/pkg/statusbar/volume"
	"code.dwrz.net/src/pkg/statusbar/wlan"
	"code.dwrz.net/src/pkg/statusbar/wwan"
)

var once = flag.Bool("once", false, "print output once")

// TODO: block labels; rendered in case of error.
func main() {
	flag.Parse()

	var (
		ctx = context.Background()
		l   = log.New(os.Stderr)
	)

	// Setup workspace and log file.
	cdir, err := os.UserCacheDir()
	if err != nil {
		l.Error.Fatalf(
			"failed to determine user cache directory: %v", err,
		)
	}
	wdir := filepath.Join(cdir, "statusbar")

	if err := os.MkdirAll(wdir, os.ModeDir|0700); err != nil {
		l.Error.Fatalf("failed to create tmp dir: %v", err)
	}

	flog, err := os.Create(wdir + "/log")
	if err != nil {
		l.Error.Fatalf("failed to create log file: %v", err)
	}
	defer flog.Close()

	// Prepare the datetime format.
	now := time.Now()
	yearEnd := time.Date(
		now.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC,
	).AddDate(0, 0, -1)

	dfmt := fmt.Sprintf(
		"+%%Y-%%m-%%d %%u/7 %%W/52 %%j/%d %%H:%%M %%Z",
		yearEnd.YearDay(),
	)

	var bar = statusbar.New(statusbar.Parameters{
		Blocks: []statusbar.Block{
			cpu.New(),
			temp.New(),
			memory.New(),
			disk.New("/", "/home"),
			eth.New("eth0"),
			wlan.New(),
			wwan.New("wwan0"),
			power.New(power.Path),
			light.New(),
			volume.New(),
			mic.New(),
			datetime.New(datetime.Parameters{
				Format:   "+%d %H:%M",
				Label:    "🇺🇸",
				Timezone: "America/New_York",
			}),
			datetime.New(datetime.Parameters{
				Format:   "+%d %H:%M",
				Label:    "🇮🇹",
				Timezone: "Europe/Rome",
			}),
			datetime.New(datetime.Parameters{
				Format:   "+%d %H:%M",
				Label:    "🇨🇳",
				Timezone: "Asia/Shanghai",
			}),
			datetime.New(datetime.Parameters{
				Format:   dfmt,
				Timezone: "UTC",
			}),
		},
		Log:       log.New(flog),
		Separator: "┃",
	})

	// Initial print.
	fmt.Println(bar.Render(ctx))
	if *once {
		os.Exit(0)
	}

	// Main loop.
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			// now := time.Now()
			fmt.Println(bar.Render(ctx))
			// fmt.Println(time.Since(now))
		}
	}
}
