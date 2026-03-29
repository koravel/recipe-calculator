package service

import dto "recipe-calculator/data/dto"

type IRecipeService interface {
	CreateRecipe(recipe dto.IRecipe) bool
	GetRecipe_ByResultMaterial(result_material dto.IMaterial) dto.IRecipe
	GetRecipe_ByIngridient(ingridient dto.IMaterial) dto.IRecipe
	DeleteRecipe(result_material dto.IMaterial) bool
	UpdateRecipe(recipe dto.IRecipe) bool
}
