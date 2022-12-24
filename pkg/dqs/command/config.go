package command

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"code.dwrz.net/src/pkg/color"
	"code.dwrz.net/src/pkg/dqs/command/help"
	"code.dwrz.net/src/pkg/dqs/diet"
	"code.dwrz.net/src/pkg/dqs/store"
	"code.dwrz.net/src/pkg/dqs/user"
	"code.dwrz.net/src/pkg/dqs/user/units"
)

var Config = &command{
	execute: configure,

	description: "configure the user settings",
	help:        help.Config,
	name:        "config",
}

func configure(args []string, date time.Time, store *store.Store) error {
	u, err := store.GetUser()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if u == nil {
		u = &user.DefaultUser
	}

	fmt.Printf(
		"%sPress enter to skip a field.%s\n",
		color.Italic, color.Reset,
	)

	var in = bufio.NewReader(os.Stdin)

	// Parse the name.
	fmt.Print("Name: ")
	line, err := in.ReadString('\n')
	if err != nil {
		return err
	}

	if line = strings.TrimSpace(line); line != "" {
		u.Name = line
	}

	// Parse the birthday.
	fmt.Print("Birthday (YYYYMMDD): ")
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}

	if line = strings.TrimSpace(line); line != "" {
		birthday, err := time.Parse("20060102", line)
		if err != nil {
			return err
		}

		u.Birthday = birthday
	}

	// Parse the units.
	fmt.Print("Units (imperial|metric): ")
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}

	if line = strings.TrimSpace(line); line != "" {
		system := units.System(line)

		if !units.Valid(system) {
			return fmt.Errorf("invalid units")
		}

		u.Units = system
	}

	// Parse the height.
	fmt.Print(fmt.Sprintf("Height (%s): ", u.Units.Height()))
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}

	if line = strings.TrimSpace(line); line != "" {
		height, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return err
		}
		if height < 0 {
			return fmt.Errorf("invalid height")
		}

		// Store all units in metric.
		if u.Units == units.Imperial {
			height = units.InchesToCentimeter(height)
		}

		u.Height = height
	}

	// Parse the weight.
	fmt.Print(fmt.Sprintf("Weight (%s): ", u.Units.Weight()))
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}

	if line = strings.TrimSpace(line); line != "" {
		weight, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return err
		}
		if weight < 0 {
			return fmt.Errorf("invalid weight")
		}

		// Store all units in metric.
		if u.Units == units.Imperial {
			weight = units.PoundsToKilogram(weight)
		}

		u.Weight = weight
	}

	// Parse the body fat.
	fmt.Print("Body Fat (%): ")
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}

	if line = strings.TrimSpace(line); line != "" {
		bf, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return err
		}

		if bf < 0 || bf > 100 {
			return fmt.Errorf("invalid body fat percentage")
		}

		u.BodyFat = bf
	}

	// Parse the diet.
	fmt.Print("Diet (omnivore|vegan|vegetarian): ")
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}

	if line = strings.TrimSpace(line); line != "" {
		if !diet.Valid(diet.Diet(line)) {
			return fmt.Errorf("unrecognized diet")
		}

		u.Diet = diet.Diet(line)
	}

	// Parse the Target Weight.
	fmt.Print(fmt.Sprintf("Target Weight (%s): ", u.Units.Weight()))
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}

	if line = strings.TrimSpace(line); line != "" {
		tw, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return err
		}

		if tw < 0 {
			return fmt.Errorf("invalid target weight")
		}

		// Store all units in metric.
		if u.Units == units.Imperial {
			tw = units.PoundsToKilogram(tw)
		}

		u.TargetWeight = tw
	}

	if err := store.UpdateUser(u); err != nil {
		return fmt.Errorf(
			"failed to store user: %w", err,
		)
	}

	return nil
}
