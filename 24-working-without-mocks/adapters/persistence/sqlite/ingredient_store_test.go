package sqlite_test

import (
	"testing"

	"github.com/ashwnacharya/working-without-mocks/adapters/persistence/sqlite"
	"github.com/ashwnacharya/working-without-mocks/domain/planner"
)

func TestSqliteIngredientStore(t *testing.T) {
	planner.IngredientStoreContract{
		NewStore: func() planner.CloseableIngredientStore {
			return sqlite.NewIngredientStore()
		},
	}.Test(t)
}
