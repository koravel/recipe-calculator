package main

import (
	"fmt"
	sqlite "recipe-calculator/data/sqlite"
)

func main() {
	sqliteDriver := sqlite.NewSQLiteDriver("./my.db")

	recipeExecuter := sqlite.NewSQLiteRecipeExecuter(sqliteDriver)

	materialExecuter := sqlite.NewSQLiteMaterialExecuter(sqliteDriver)

	recipeExecuter.CreateTable()
	materialExecuter.CreateTable()

	err := materialExecuter.CreatePrimaryMaterial("iron ore")
	fmt.Println(err)
}
