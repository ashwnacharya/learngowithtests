package sqlite_test

import (
	"github.com/ashwnacharya/working-without-mocks/adapters/persistence/sqlite"
	"github.com/ashwnacharya/working-without-mocks/domain/recipe"
	"testing"
)

func TestSqliteIngredientStore(t *testing.T) {
	recipe.StoreContract {
		NewStore: func() recipe.Store {
			return sqlite.NewIngredientStore()
		},
	}.Test(t)
}