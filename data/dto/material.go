package dto

type IMaterial interface {
	CalculateRecipe()
}

type Material struct {
	name   string
	recipe IRecipe
}
