package diet

import "code.dwrz.net/src/pkg/dqs/category"

type Diet string

const (
	Omnivore   Diet = "omnivore"
	Vegan      Diet = "vegan"
	Vegetarian Diet = "vegetarian"
)

func Valid(d Diet) bool {
	switch d {
	case Omnivore, Vegan, Vegetarian:
		return true
	default:
		return false
	}
}

func (d Diet) Template() map[string]category.Category {
	switch d {
	case Omnivore:
		return omnivore
	case Vegan:
		return vegan
	case Vegetarian:
		return vegetarian
	default:
		return nil
	}
}

func (d Diet) MaxScore() int {
	tmpl := d.Template()

	var max int
	for _, category := range tmpl {
		for _, portion := range category.Portions {
			if portion.Points > 0 {
				max += portion.Points
			}
		}
	}

	return max
}
