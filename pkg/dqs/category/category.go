package category

import (
	"fmt"

	"code.dwrz.net/src/pkg/color"
	"code.dwrz.net/src/pkg/dqs/portion"
)

type Category struct {
	Name        string            `json:"name"`
	Portions    []portion.Portion `json:"portions"`
	HighQuality bool              `json:"highQuality"`
}

func (c *Category) Add(q float64) error {
	for i := range c.Portions {
		p := &c.Portions[i]
		// Done if there's less than a half-portion left.
		if q < 0.5 {
			break
		}

		switch p.Amount {
		case portion.Full:
			continue
		case portion.Half:
			p.Amount = portion.Full
			q -= 0.5
		case portion.None:
			if q >= 1.0 {
				p.Amount = portion.Full
				q -= 1.0
			} else if q >= 0.5 {
				p.Amount = portion.Half
				q -= 0.5
			}
		}
	}

	// Check if we were able to add all the portions.
	if q > 0.5 {
		return fmt.Errorf("too many portions of %s", c.Name)
	}

	return nil
}

func (c *Category) Score() (total float64) {
	for _, p := range c.Portions {
		total += p.Score()
	}

	return total
}

func (c *Category) FormatPrint() string {
	var str = fmt.Sprintf("│%-26s│", c.Name)

	for _, p := range c.Portions {
		var cellColor color.Color

		switch {
		case p.Points > 0 && p.Amount == portion.Full:
			cellColor = color.BackgroundBrightGreen
		case p.Points > 0 && p.Amount == portion.Half:
			cellColor = color.BrightGreen
		case p.Points == 0 && p.Amount == portion.Full:
			cellColor = color.BackgroundBrightYellow
		case p.Points == 0 && p.Amount == portion.Half:
			cellColor = color.BrightYellow
		case p.Points < 0 && p.Amount == portion.Full:
			cellColor = color.BackgroundBrightRed
		case p.Points < 0 && p.Amount == portion.Half:
			cellColor = color.BrightRed
		}

		str += fmt.Sprintf("%s%+d%s│", cellColor, p.Points, color.Reset)
	}

	return str
}

func (c *Category) Remove(q float64) error {
	for i := len(c.Portions) - 1; i >= 0; i-- {
		p := &c.Portions[i]
		// Done if there's less than a half-portion left.
		if q < 0.5 {
			break
		}

		switch p.Amount {
		case portion.None:
			continue
		case portion.Half:
			p.Amount = portion.None
			q -= 0.5
		case portion.Full:
			if q >= 1.0 {
				p.Amount = portion.None
				q -= 1.0
			} else if q >= 0.5 {
				p.Amount = portion.Half
				q -= 0.5
			}
		}
	}

	// Check if we were able to remove all the portions.
	if q > 0.5 {
		return fmt.Errorf("too many portions of %s", c.Name)
	}

	return nil
}
