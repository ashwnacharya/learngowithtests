package sqlite_test

import (
	"github.com/ashwnacharya/working-without-mocks/adapters/persistence/sqlite"
	"github.com/ashwnacharya/working-without-mocks/domain/planner"
	"testing"
)

func TestSqliteIngredientStore(t *testing.T) {
	planner.StoreContract {
		NewStore: func() planner.CloseableStore {
			return sqlite.NewIngredientStore()
		},
	}.Test(t)
}