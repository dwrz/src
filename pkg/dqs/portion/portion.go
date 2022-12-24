package portion

import (
	"fmt"
	"math"
	"strconv"
)

type Amount string

const (
	None Amount = ""
	Full Amount = "full"
	Half Amount = "half"
)

func (a Amount) Claimed() bool {
	switch a {
	case Half, Full:
		return true
	default:
		return false
	}
}

type Portion struct {
	Points int    `json:"points"`
	Amount Amount `json:"amount"`
}

func (p *Portion) Score() float64 {
	switch p.Amount {
	case Full:
		return float64(p.Points)
	case Half:
		return float64(p.Points) / 2
	case None:
		fallthrough
	default:
		return float64(0)
	}
}

func Parse(quantity string) (float64, error) {
	q, err := strconv.ParseFloat(quantity, 64)
	if err != nil {
		return 0, fmt.Errorf(
			"failed to parse quantity %s: %w", quantity, err,
		)
	}
	if q < 0 {
		return 0, fmt.Errorf(
			"negative portions (%s) are not allowed", quantity,
		)
	}
	if math.Mod(q, 0.5) != 0 {
		return 0, fmt.Errorf(
			"quantity %f is not a full or half portion", q,
		)
	}

	return q, nil
}
