package entry

import (
	"fmt"

	"sort"
	"strconv"
	"strings"

	"code.dwrz.net/src/pkg/color"
	"code.dwrz.net/src/pkg/dqs/category"
	"code.dwrz.net/src/pkg/dqs/stats"
	"code.dwrz.net/src/pkg/dqs/user"
)

const (
	width  = 44
	top    = "┌────────────────────────────────────────────┐\n"
	hr     = "├────────────────────────────────────────────┤\n"
	bottom = "└────────────────────────────────────────────┘\n"
)

func centerPad(s string, width int) string {
	return fmt.Sprintf(
		"%[1]*s", -width, fmt.Sprintf("%[1]*s", (width+len(s))/2, s),
	)
}

// FormatPrint formats an entry for display to the user.
func (e *Entry) FormatPrint(u *user.User) string {
	// Assemble the data for display.
	var (
		highQuality []category.Category
		lowQuality  []category.Category
	)

	// Separate high quality and low quality categories.
	for _, c := range e.Categories {
		if c.HighQuality {
			highQuality = append(highQuality, c)
			continue
		}
		lowQuality = append(lowQuality, c)
	}

	// Sort alphabetically for consistent appearance.
	sort.Slice(highQuality, func(i, j int) bool {
		return highQuality[i].Name < highQuality[j].Name
	})
	sort.Slice(lowQuality, func(i, j int) bool {
		return lowQuality[i].Name < lowQuality[j].Name
	})

	// Prepare a string for display to the user.
	var str strings.Builder

	// Format the date.
	str.WriteString(top)
	str.WriteString(fmt.Sprintf(
		"│%s%s%s│\n",
		color.BrightBlack,
		centerPad(e.Date.Format(dateDisplayFormat), width),
		color.Reset,
	))
	str.WriteString(hr)

	// Format high-quality categories.
	str.WriteString(fmt.Sprintf("│%-44s│\n", "High Quality"))
	str.WriteString(hr)
	for _, c := range highQuality {
		str.WriteString(c.FormatPrint())
		str.WriteString("\n")
	}

	// Format low-quality categories.
	str.WriteString(hr)
	str.WriteString(fmt.Sprintf("│%-44s│\n", "Low Quality"))
	str.WriteString(hr)
	for _, c := range lowQuality {
		str.WriteString(c.FormatPrint())
		str.WriteString("\n")
	}
	str.WriteString(hr)

	// Format the total.
	var totalColor color.Color
	total := e.Score()
	switch {
	case total >= 15:
		totalColor = color.BrightGreen
	case total >= 0:
		totalColor = color.BrightYellow
	default:
		totalColor = color.BrightRed

	}

	max := e.Diet.MaxScore()
	formattedTotal := strconv.FormatFloat(total, 'f', -1, 64)
	digits := len(fmt.Sprintf("%s / %d", formattedTotal, max))
	space := width - digits
	format := fmt.Sprintf("%%-%ds%%s%%s%%s / %%d │\n", space)
	str.WriteString(fmt.Sprintf(
		format,
		"│Total / Max: ", totalColor, formattedTotal, color.Reset, max,
	))
	str.WriteString(hr)

	// Format the weight, delta, and body fat.
	digits = len(fmt.Sprintf(
		"%.2f / %.2f", e.UnitWeight(u.Units), u.UnitTargetWeight(),
	))
	space = width - digits
	format = fmt.Sprintf("%%-%ds%%.2f / %%.2f │\n", space)
	str.WriteString(fmt.Sprintf(
		format,
		"│Weight / Target:",
		e.UnitWeight(u.Units),
		u.UnitTargetWeight(),
	))

	diff := e.UnitWeight(u.Units) - u.UnitTargetWeight()
	digits = len(fmt.Sprintf("%+.2f", diff))
	space = width - digits
	format = fmt.Sprintf("%%-%ds%%+.2f │\n", space)
	str.WriteString(fmt.Sprintf(format, "│Weight Δ:", diff))

	bfw := e.UnitBodyFatWeight(u.Units)
	digits = len(fmt.Sprintf(
		"%.2f%% (%.2f %s)",
		e.BodyFat, bfw, u.Units.Weight(),
	))
	space = width - digits
	format = fmt.Sprintf("%%-%ds%%.2f%%%% (%%.2f %%s) │\n", space)
	str.WriteString(fmt.Sprintf(
		format,
		"│Body Fat:",
		e.BodyFat, bfw, u.Units.Weight(),
	))

	bmi := stats.BMI(u.Height, e.Weight)
	digits = len(fmt.Sprintf("%.2f", bmi))
	space = width - digits
	format = fmt.Sprintf("%%-%ds%%.2f │\n", space)
	str.WriteString(fmt.Sprintf(format, "│BMI:", bmi))

	str.WriteString(bottom)

	// If a note is set, format it.
	if e.Note != "" {
		str.WriteString(fmt.Sprintf(
			"\n%s%s%s\n", color.Italic, e.Note, color.Reset,
		))
	}

	return str.String()
}
