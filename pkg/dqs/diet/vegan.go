package diet

import "code.dwrz.net/src/pkg/dqs/category"

var vegan = map[string]category.Category{
	// High Quality
	category.Fruit.Name:                category.Fruit,
	category.HQBeverages.Name:          category.HQBeverages,
	category.HQProcessedFoods.Name:     category.HQProcessedFoods,
	category.LegumesPlantProteins.Name: category.LegumesPlantProteins,
	category.NutsSeedsHealthyOils.Name: category.NutsSeedsHealthyOils,
	category.Vegetables.Name:           category.Vegetables,
	category.WholeGrains.Name:          category.WholeGrains,

	// Low Quality
	category.FriedFoods.Name:    category.FriedFoods,
	category.LQBeverages.Name:   category.LQBeverages,
	category.Other.Name:         category.Other,
	category.RefinedGrains.Name: category.RefinedGrains,
	category.Sweets.Name:        category.Sweets,
}
