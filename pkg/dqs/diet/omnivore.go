package diet

import "code.dwrz.net/src/pkg/dqs/category"

var omnivore = map[string]category.Category{
	// High Quality
	category.Dairy.Name:                  category.Dairy,
	category.Fruit.Name:                  category.Fruit,
	category.HQBeverages.Name:            category.HQBeverages,
	category.HQProcessedFoods.Name:       category.HQProcessedFoods,
	category.NutsSeedsHealthyOils.Name:   category.NutsSeedsHealthyOils,
	category.UnprocessedMeatSeafood.Name: category.UnprocessedMeatSeafood,
	category.Vegetables.Name:             category.Vegetables,
	category.WholeGrains.Name:            category.WholeGrains,

	// Low Quality
	category.FriedFoods.Name:    category.FriedFoods,
	category.LQBeverages.Name:   category.LQBeverages,
	category.Other.Name:         category.Other,
	category.ProcessedMeat.Name: category.ProcessedMeat,
	category.RefinedGrains.Name: category.RefinedGrains,
	category.Sweets.Name:        category.Sweets,
}
