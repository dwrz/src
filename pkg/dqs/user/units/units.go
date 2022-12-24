package units

type System string

const (
	Default  System = "metric"
	Imperial System = "imperial"
	Metric   System = "metric"
)

func Valid(s System) bool {
	switch s {
	case Imperial, Metric:
		return true
	default:
		return false
	}
}

type Unit string

const (
	Centimeter Unit = "cm"
	Kilogram   Unit = "kg"

	Inches Unit = "in"
	Pounds Unit = "lbs"
)

func (s System) Weight() Unit {
	switch s {
	case Imperial:
		return Pounds
	default:
		return Kilogram
	}
}

func (s System) Height() Unit {
	switch s {
	case Imperial:
		return Inches
	default:
		return Centimeter
	}
}
