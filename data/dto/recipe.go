package dto

type IRecipe interface {
	AddIngridients()
}

type Recipe struct {
	result      IMaterial
	ingridients []IMaterial
}
