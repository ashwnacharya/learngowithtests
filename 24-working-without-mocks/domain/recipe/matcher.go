package recipe

import (
	"context"
	"github.com/ashwnacharya/working-without-mocks/domain/ingredients"
)

type Book interface {
	GetRecipes() []Recipe
}

type Matcher struct {
	recipeBook Book
	ingredientStore Store
}

func NewMatcher(recipeBook Book, ingredientStore Store) *Matcher {
	return &Matcher{recipeBook: recipeBook, ingredientStore: ingredientStore}
}

func (m *Matcher) SuggestRecipes(ctx context.Context) ([]Recipe, error) {
	
	availableIngredients, err := m.ingredientStore.GetIngredients(ctx)

	if err != nil {
		return nil, err
	}
	
	var suggestions []Recipe

	for _, recipe := range m.recipeBook.GetRecipes() {
		if m.haveIngredients(availableIngredients, recipe) {
			suggestions = append(suggestions, recipe)
		}
	}

	return suggestions, nil
}

func (m *Matcher) haveIngredients(availableIngredients ingredients.Ingredients, recipe Recipe) bool {
	for _, recipeIngredient := range recipe.Ingredients {
		if !ingredients.Ingredients(availableIngredients).Has(recipeIngredient) {
			return false
		}
	}
	return true
}

