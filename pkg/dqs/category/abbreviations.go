package category

var Abbreviations = map[string]string{
	// High Quality
	"d":    Dairy.Name,
	"f":    Fruit.Name,
	"hqb":  HQBeverages.Name,
	"hqpf": HQProcessedFoods.Name,
	"lpp":  LegumesPlantProteins.Name,
	"nsh":  NutsSeedsHealthyOils.Name,
	"ums":  UnprocessedMeatSeafood.Name,
	"v":    Vegetables.Name,
	"wg":   WholeGrains.Name,

	// Low Quality
	"ff":  FriedFoods.Name,
	"lqb": LQBeverages.Name,
	"o":   Other.Name,
	"p":   ProcessedMeat.Name,
	"rg":  RefinedGrains.Name,
	"s":   Sweets.Name,
}
