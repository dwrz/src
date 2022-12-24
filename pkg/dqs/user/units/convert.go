package units

func InchesToCentimeter(inches float64) float64 {
	return inches * 2.54
}

func CentimeterToInches(cm float64) float64 {
	return cm / 2.54
}

func PoundsToKilogram(lbs float64) float64 {
	return lbs * 0.45359237
}

func KilogramToPounds(kg float64) float64 {
	return kg * 2.20462262185
}
