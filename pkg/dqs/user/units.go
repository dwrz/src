package user

import (
	"code.dwrz.net/src/pkg/dqs/stats"
	"code.dwrz.net/src/pkg/dqs/user/units"
)

func (u *User) UnitBodyFatWeight() float64 {
	bfw := stats.BodyFatWeight(u.Weight, u.BodyFat)

	if u.Units == units.Metric {
		return bfw
	}

	return units.KilogramToPounds(bfw)
}

func (u *User) UnitHeight() float64 {
	if u.Units == units.Metric {
		return u.Height
	}

	return units.CentimeterToInches(u.Height)
}

func (u *User) UnitTargetWeight() float64 {
	if u.Units == units.Metric {
		return u.TargetWeight
	}

	return units.KilogramToPounds(u.TargetWeight)
}

func (u *User) UnitWeight() float64 {
	if u.Units == units.Metric {
		return u.Weight
	}

	return units.KilogramToPounds(u.Weight)
}
