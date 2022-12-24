package category

import "code.dwrz.net/src/pkg/dqs/portion"

var (
	Dairy = Category{
		Name: "Dairy",
		Portions: []portion.Portion{
			{Points: 2},
			{Points: 1},
			{Points: 1},
			{Points: 0},
			{Points: -1},
			{Points: -2},
		},
		HighQuality: true,
	}

	Fruit = Category{
		Name: "Fruit",
		Portions: []portion.Portion{
			{Points: 2},
			{Points: 2},
			{Points: 2},
			{Points: 1},
			{Points: 0},
			{Points: 0},
		},
		HighQuality: true,
	}

	HQBeverages = Category{
		Name: "High Quality Beverages",
		Portions: []portion.Portion{
			{Points: 1},
			{Points: 1},
			{Points: 0},
			{Points: 0},
			{Points: -1},
			{Points: -2},
		},
		HighQuality: true,
	}

	HQProcessedFoods = Category{
		Name: "Processed Foods",
		Portions: []portion.Portion{
			{Points: 1},
			{Points: 0},
			{Points: -1},
			{Points: -2},
			{Points: -2},
			{Points: -2},
		},
		HighQuality: true,
	}

	LegumesPlantProteins = Category{
		Name: "Legumes & Plant Proteins",
		Portions: []portion.Portion{
			{Points: 2},
			{Points: 2},
			{Points: 1},
			{Points: 0},
			{Points: -1},
			{Points: -1},
		},
		HighQuality: true,
	}

	NutsSeedsHealthyOils = Category{
		Name: "Nuts, Seeds, Healthy Oils",
		Portions: []portion.Portion{
			{Points: 2},
			{Points: 2},
			{Points: 1},
			{Points: 0},
			{Points: 0},
			{Points: -1},
		},
		HighQuality: true,
	}

	UnprocessedMeatSeafood = Category{
		Name: "Unprocessed Meat & Seafood",
		Portions: []portion.Portion{
			{Points: 2},
			{Points: 1},
			{Points: 1},
			{Points: 0},
			{Points: -1},
			{Points: -2},
		},
		HighQuality: true,
	}

	Vegetables = Category{
		Name: "Vegetables",
		Portions: []portion.Portion{
			{Points: 2},
			{Points: 2},
			{Points: 2},
			{Points: 1},
			{Points: 0},
			{Points: 0},
		},
		HighQuality: true,
	}

	WholeGrains = Category{
		Name: "Whole Grains",
		Portions: []portion.Portion{
			{Points: 2},
			{Points: 2},
			{Points: 1},
			{Points: 0},
			{Points: 0},
			{Points: -1},
		},
		HighQuality: true,
	}
)
