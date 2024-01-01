package inmemory

import (
	"context"
	"github.com/ashwnacharya/working-without-mocks/domain/ingredients"
)

type IngredientStore struct {
	Ingredients ingredients.Ingredients
}

func NewIngredientStore() *IngredientStore {
	return &IngredientStore{}
}

func (s *IngredientStore) GetIngredients(ctx context.Context) ([]ingredients.Ingredient, error) {
	return s.Ingredients, nil
}

func (s *IngredientStore) Store(ctx context.Context, ingredients ...ingredients.Ingredient) error {
	for idx, ingredient := range ingredients {
		if s.Ingredients.Has(ingredient) {
			s.Ingredients[idx].Quantity += ingredient.Quantity
		} else {
			s.Ingredients = append(s.Ingredients, ingredient)
		}
	}
	return nil
}

func (i *IngredientStore) Close() {
}