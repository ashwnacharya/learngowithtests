package recipe

import (
	"context"
	"testing"
	"github.com/alecthomas/assert/v2"
	"github.com/ashwnacharya/working-without-mocks/domain/ingredients"
)

type Store interface {
	GetIngredients(ctx context.Context) ([]ingredients.Ingredient, error)
	Store(context.Context, ...ingredients.Ingredient) error
}

type StoreContract struct {
	NewStore func() Store
	Cleanup func()
}


func (s StoreContract) Test(t *testing.T) {

	t.Run("it returns what was stored", func(t *testing.T) {
	
		ctx := context.Background()

		want := []ingredients.Ingredient{
			{Name: "flour", Quantity: 1},
			{Name: "sugar", Quantity: 1},
			{Name: "eggs", Quantity: 2},
		}

		store := s.NewStore()

		err := store.Store(ctx, want...)
		assert.NoError(t, err)

		got, err := store.GetIngredients(ctx)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})
}