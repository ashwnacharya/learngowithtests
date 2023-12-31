package inmemory

import "github.com/ashwnacharya/working-without-mocks/domain/recipe"

type RecipeStore struct {
	Recipes []recipe.Recipe
}

func (s RecipeStore) GetRecipes() []recipe.Recipe {
	return s.Recipes
}