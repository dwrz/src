package user

import (
	"fmt"
	"strings"

	"code.dwrz.net/src/pkg/color"
	"code.dwrz.net/src/pkg/dqs/stats"
)

func (u *User) FormatPrint() string {
	var str strings.Builder

	str.WriteString(color.BrightBlack)
	str.WriteString("Name: ")
	str.WriteString(color.Reset)
	str.WriteString(u.Name)
	str.WriteString("\n")

	str.WriteString(color.BrightBlack)
	str.WriteString("Birthday: ")
	str.WriteString(color.Reset)
	str.WriteString(u.Birthday.Format("2006-01-02"))
	str.WriteString("\n")

	str.WriteString(color.BrightBlack)
	str.WriteString("Units: ")
	str.WriteString(color.Reset)
	str.WriteString(fmt.Sprintf("%s", u.Units))
	str.WriteString("\n")

	str.WriteString(color.BrightBlack)
	str.WriteString("Height: ")
	str.WriteString(color.Reset)
	str.WriteString(fmt.Sprintf(
		"%.2f %s", u.UnitHeight(), u.Units.Height()),
	)
	str.WriteString("\n")

	str.WriteString(color.BrightBlack)
	str.WriteString("Weight: ")
	str.WriteString(color.Reset)
	str.WriteString(fmt.Sprintf(
		"%.2f %s", u.UnitWeight(), u.Units.Weight()),
	)
	str.WriteString("\n")

	str.WriteString(color.BrightBlack)
	str.WriteString("Body Fat: ")
	str.WriteString(color.Reset)
	str.WriteString(fmt.Sprintf(
		"%.2f%% (%.2f %s)",
		u.BodyFat, u.UnitBodyFatWeight(), u.Units.Weight(),
	))
	str.WriteString("\n")

	str.WriteString(color.BrightBlack)
	str.WriteString("BMI: ")
	str.WriteString(color.Reset)
	str.WriteString(fmt.Sprintf("%.2f", stats.BMI(u.Height, u.Weight)))
	str.WriteString("\n")

	str.WriteString(color.BrightBlack)
	str.WriteString("Diet: ")
	str.WriteString(color.Reset)
	str.WriteString(fmt.Sprintf("%s", u.Diet))
	str.WriteString("\n")

	str.WriteString(color.BrightBlack)
	str.WriteString("Target Weight: ")
	str.WriteString(color.Reset)
	str.WriteString(fmt.Sprintf(
		"%.2f %s", u.UnitTargetWeight(), u.Units.Weight(),
	))
	str.WriteString("\n")

	return str.String()
}
