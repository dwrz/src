package user

import (
	"fmt"
	"time"

	"code.dwrz.net/src/pkg/dqs/diet"
	"code.dwrz.net/src/pkg/dqs/user/units"
)

type User struct {
	Birthday     time.Time    `json:"birthday"`
	BodyFat      float64      `json:"bodyFat"`
	Diet         diet.Diet    `json:"diet"`
	Height       float64      `json:"height"`
	Name         string       `json:"name"`
	TargetWeight float64      `json:"targetWeight"`
	Units        units.System `json:"units"`
	Weight       float64      `json:"weight"`
}

var DefaultUser = User{
	Diet:  diet.Vegetarian,
	Units: units.Metric,
}

func (u *User) SetDiet(d diet.Diet) error {
	switch d {
	case diet.Omnivore:
		u.Diet = diet.Omnivore
	case diet.Vegan:
		u.Diet = diet.Vegan
	case diet.Vegetarian:
		u.Diet = diet.Vegetarian
	default:
		return fmt.Errorf("unrecognized diet: %s", d)
	}

	return nil
}
