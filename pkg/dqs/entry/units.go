package entry

import (
	"code.dwrz.net/src/pkg/dqs/stats"
	"code.dwrz.net/src/pkg/dqs/user/units"
)

func (e *Entry) UnitBodyFatWeight(system units.System) float64 {
	bfw := stats.BodyFatWeight(e.Weight, e.BodyFat)

	if system == units.Metric {
		return bfw
	}

	return units.KilogramToPounds(bfw)
}

func (e *Entry) UnitHeight(system units.System) float64 {
	if system == units.Metric {
		return e.Height
	}

	return units.CentimeterToInches(e.Height)
}

func (e *Entry) UnitWeight(system units.System) float64 {
	if system == units.Metric {
		return e.Weight
	}

	return units.KilogramToPounds(e.Weight)
}
