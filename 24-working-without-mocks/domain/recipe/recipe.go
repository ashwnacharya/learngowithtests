package recipe

import "github.com/ashwnacharya/working-without-mocks/domain/ingredients"

type Recipe struct {
	Name string
	Ingredients []ingredients.Ingredient
}

