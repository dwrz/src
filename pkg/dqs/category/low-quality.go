package category

import dqs "code.dwrz.net/src/pkg/dqs/portion"

// Low Quality
var (
	FriedFoods = Category{
		Name: "Fried Foods",
		Portions: []dqs.Portion{
			{Points: -2},
			{Points: -2},
			{Points: -2},
			{Points: -2},
			{Points: -2},
			{Points: -2},
		},
		HighQuality: false,
	}

	LQBeverages = Category{
		Name: "Low Quality Beverages",
		Portions: []dqs.Portion{
			{Points: -2},
			{Points: -2},
			{Points: -2},
			{Points: -2},
			{Points: -2},
			{Points: -2},
		},
		HighQuality: false,
	}

	Other = Category{
		Name: "Other",
		Portions: []dqs.Portion{
			{Points: -1},
			{Points: -2},
			{Points: -2},
			{Points: -2},
			{Points: -2},
			{Points: -2},
		},
		HighQuality: false,
	}

	ProcessedMeat = Category{
		Name: "Processed Meat",
		Portions: []dqs.Portion{
			{Points: -2},
			{Points: -2},
			{Points: -2},
			{Points: -2},
			{Points: -2},
			{Points: -2},
		},
		HighQuality: false,
	}

	RefinedGrains = Category{
		Name: "Refined Grains",
		Portions: []dqs.Portion{
			{Points: -1},
			{Points: -1},
			{Points: -2},
			{Points: -2},
			{Points: -2},
			{Points: -2},
		},
		HighQuality: false,
	}

	Sweets = Category{
		Name: "Sweets",
		Portions: []dqs.Portion{
			{Points: -2},
			{Points: -2},
			{Points: -2},
			{Points: -2},
			{Points: -2},
			{Points: -2},
		},
		HighQuality: false,
	}
)
