package inmemory_test

import (
	"testing"

	"github.com/ashwnacharya/working-without-mocks/adapters/persistence/inmemory"
	"github.com/ashwnacharya/working-without-mocks/domain/planner"
)

func TestInMemoryIngredientsStore(t *testing.T) {
	planner.IngredientStoreContract{
		NewStore: func() planner.CloseableIngredientStore {
			return inmemory.NewIngredientStore()
		},
	}.Test(t)
}
