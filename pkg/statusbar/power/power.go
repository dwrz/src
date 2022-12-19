package power

import (
	"bufio"
	"context"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://www.kernel.org/doc/html/latest/power/power_supply_class.html
const Path = "/sys/class/power_supply"

type Block struct {
	path string
}

func New(path string) *Block {
	return &Block{path: path}
}

func (b *Block) Name() string {
	return "power"
}

func (b *Block) Render(ctx context.Context) (string, error) {
	// Get the power supplies.
	var supplies []string

	files, err := os.ReadDir(b.path)
	if err != nil {
		return "", fmt.Errorf("failed to read dir %s: %v", b.path, err)
	}
	for _, f := range files {
		name := f.Name()

		if name == "AC" || strings.Contains(name, "BAT") {
			supplies = append(supplies, f.Name())
		}
	}
	sort.Slice(supplies, func(i, j int) bool {
		return supplies[i] < supplies[j]
	})

	var sections = []string{}
	for _, s := range supplies {
		path := fmt.Sprintf("%s/%s/uevent", b.path, s)
		f, err := os.Open(path)
		if err != nil {
			return "", fmt.Errorf(
				"failed to open %s: %v", path, err,
			)
		}
		defer f.Close()

		switch {
		case s == "AC":
			sections = append(sections, outputAC(f))
		case strings.Contains(s, "BAT"):
			sections = append(sections, outputBAT(f))
		}
	}

	var output strings.Builder
	for i, s := range sections {
		if i > 0 {
			output.WriteRune(' ')
			output.WriteRune(' ')
		}
		output.WriteString(s)
	}

	return output.String(), nil
}

func outputAC(f *os.File) string {
	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		k, v, found := strings.Cut(scanner.Text(), "=")
		if !found {
			continue
		}
		if k == "POWER_SUPPLY_ONLINE" {
			online, err := strconv.ParseBool(v)
			if err != nil {
				break
			}
			if online {
				return ""
			} else {
				return ""
			}
		}
	}

	return ""
}

func outputBAT(f *os.File) string {
	// Compile the stats.
	var (
		stats   = map[string]string{}
		scanner = bufio.NewScanner(f)
	)
	for scanner.Scan() {
		k, v, found := strings.Cut(scanner.Text(), "=")
		if !found {
			continue
		}
		// Exit early if the battery is not present.
		if k == "POWER_SUPPLY_PRESENT" && v == "0" {
			return ""
		}

		stats[k] = v
	}

	// Assemble output.
	var icon rune
	switch stats["POWER_SUPPLY_STATUS"] {
	case "Charging":
		icon = ''
	case "Discharging":
		icon = ''
	case "Full":
		icon = ''
	default:
		icon = ''
	}

	// Get capacity.
	var capacity string
	if s, exists := stats["POWER_SUPPLY_CAPACITY"]; exists {
		capacity = s + "%"
	}

	// Get remaining.
	pn, _ := strconv.ParseFloat(stats["POWER_SUPPLY_POWER_NOW"], 64)
	en, _ := strconv.ParseFloat(stats["POWER_SUPPLY_ENERGY_NOW"], 64)
	if r := remaining(en, pn); r != "" {
		return fmt.Sprintf("%c %s %s", icon, capacity, r)
	}

	return fmt.Sprintf("%c %s", icon, capacity)
}

func remaining(energy, power float64) string {
	if energy == 0 || power == 0 {
		return ""
	}

	// Calculate the remaining hours.
	var hours = energy / power
	if hours == 0 {
		return ""
	}

	return fmt.Sprintf(
		"%d:%02d",
		int(hours),
		int((hours-math.Floor(hours))*60),
	)
}
