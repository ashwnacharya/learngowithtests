package inmemory

import (
	"context"
	"github.com/ashwnacharya/working-without-mocks/domain/ingredients"
)

type IngredientStore struct {
	Ingredients []ingredients.Ingredient
}

func (s *IngredientStore) GetIngredients(ctx context.Context) ([]ingredients.Ingredient, error) {
	return s.Ingredients, nil
}

func (s *IngredientStore) Store(ctx context.Context, ingredients ...ingredients.Ingredient) error {
	s.Ingredients = append(s.Ingredients, ingredients...)
	return nil
}
