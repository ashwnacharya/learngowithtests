package planner

import (
	"context"
	"github.com/ashwnacharya/working-without-mocks/domain/ingredients"
	"github.com/ashwnacharya/working-without-mocks/domain/recipe"
)

type Book interface {
	GetRecipes() []recipe.Recipe
}

type Planner struct {
	recipeBook Book
	ingredientStore Store
}

func New(recipeBook Book, ingredientStore Store) *Planner {
	return &Planner{recipeBook: recipeBook, ingredientStore: ingredientStore}
}

func (p Planner) SuggestRecipes(ctx context.Context) ([]recipe.Recipe, error) {
	
	availableIngredients, err := p.ingredientStore.GetIngredients(ctx)

	if err != nil {
		return nil, err
	}
	
	var suggestions []recipe.Recipe

	for _, recipe := range p.recipeBook.GetRecipes() {
		if haveIngredients(availableIngredients, recipe) {
			suggestions = append(suggestions, recipe)
		}
	}

	return suggestions, nil
}

func haveIngredients(availableIngredients ingredients.Ingredients, recipe recipe.Recipe) bool {
	for _, recipeIngredient := range recipe.Ingredients {
		if !availableIngredients.Has(recipeIngredient) {
			return false
		}
	}
	return true
}

