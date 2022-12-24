package stats

import (
	"math"
)

func BMI(height, weight float64) float64 {
	if height > 0 {
		return weight / math.Pow(height/100, 2)
	}

	return 0.0
}

func BodyFatWeight(weight, bodyFat float64) float64 {
	return (weight * bodyFat) / 100
}
