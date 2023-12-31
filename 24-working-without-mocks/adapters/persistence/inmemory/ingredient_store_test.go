package inmemory_test

import (
	"testing"
	"github.com/ashwnacharya/working-without-mocks/adapters/persistence/inmemory"
	"github.com/ashwnacharya/working-without-mocks/domain/recipe"
)

func TestInMemoryIngredientsStore(t *testing.T) {
	recipe.StoreContract{
		NewStore: func() recipe.Store {
			return inmemory.NewIngredientStore()
		},
	}.Test(t)
}