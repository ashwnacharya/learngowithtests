package recipe_test

import (
	"context"
	"github.com/alecthomas/assert/v2"
	"github.com/ashwnacharya/working-without-mocks/adapters/persistence/inmemory"
	"github.com/ashwnacharya/working-without-mocks/adapters/persistence/sqlite"
	"github.com/ashwnacharya/working-without-mocks/domain/ingredients"
	"github.com/ashwnacharya/working-without-mocks/domain/recipe"
	"testing"
)


var (
	bananaBread = recipe.Recipe{
		Name: "Banana Bread",
		Ingredients: []ingredients.Ingredient{
			{Name: "Bananas", Quantity: 2},
			{Name: "Flour", Quantity: 1},
			{Name: "Eggs", Quantity: 2},
		},
	}
	bananaMilkshake = recipe.Recipe{
		Name: "Banana Milkshake",
		Ingredients: []ingredients.Ingredient{
			{Name: "Bananas", Quantity: 2},
			{Name: "Milk", Quantity: 1},
		},
	}
	recipeStore = inmemory.RecipeStore{Recipes: []recipe.Recipe{bananaBread, bananaMilkshake}}
)

func TestRecipeMatcher(t *testing.T) {
	t.Run("with in memory store", func(t *testing.T) {
		RecipeMatcherTest(t, func() CloseAbleStore {
			return inmemory.NewIngredientStore()
		})
	})

	t.Run("with sqlite", func(t *testing.T) {
		RecipeMatcherTest(t, func() CloseAbleStore {
			return sqlite.NewIngredientStore()
		})
	})
}

type CloseAbleStore interface {
	recipe.Store
	Close()
}

func RecipeMatcherTest(t *testing.T, newStore func() CloseAbleStore) {
	t.Run("if we have no ingredients we can't make anything", func(t *testing.T) {
		store := newStore()
		t.Cleanup(store.Close)
		assertAvailableRecipes(t, store, []recipe.Recipe{})
	})

	t.Run("if we have the ingredients for banana bread we can make it", func(t *testing.T) {
		store := newStore()
		t.Cleanup(store.Close)

		assert.NoError(t, store.Store(
			context.Background(),
			ingredients.Ingredient{Name: "Bananas", Quantity: 2},
			ingredients.Ingredient{Name: "Flour", Quantity: 1},
			ingredients.Ingredient{Name: "Eggs", Quantity: 2},
		))
		assertAvailableRecipes(t, store, []recipe.Recipe{bananaBread})
	})

	t.Run("if we have bananas and milk, we can make banana milkshake", func(t *testing.T) {
		store := newStore()
		t.Cleanup(store.Close)

		store.Store(
			context.Background(),
			ingredients.Ingredient{Name: "Bananas", Quantity: 2},
			ingredients.Ingredient{Name: "Milk", Quantity: 1},
		)
		assertAvailableRecipes(t, store, []recipe.Recipe{bananaMilkshake})
	})

	t.Run("if we have ingredients for banana bread and milkshake, we can make both", func(t *testing.T) {
		store := newStore()
		t.Cleanup(store.Close)

		store.Store(
			context.Background(),
			ingredients.Ingredient{Name: "Bananas", Quantity: 2},
			ingredients.Ingredient{Name: "Flour", Quantity: 1},
			ingredients.Ingredient{Name: "Eggs", Quantity: 2},
			ingredients.Ingredient{Name: "Milk", Quantity: 1},
		)
		assertAvailableRecipes(t, store, []recipe.Recipe{bananaMilkshake, bananaBread})
	})

}

func assertAvailableRecipes(t *testing.T, ingredientStore recipe.Store, expectedRecipes []recipe.Recipe) {
	t.Helper()
	suggestions, _ := recipe.NewMatcher(recipeStore, ingredientStore).SuggestRecipes(context.Background())

	// create a map to count occurrences of each recipe in the suggestions
	suggestionCounts := make(map[string]int)
	for _, suggestion := range suggestions {
		suggestionCounts[suggestion.Name]++
	}

	// check that the counts of the expected recipes match the actual counts in the suggestions
	for _, expectedRecipe := range expectedRecipes {
		actualCount, ok := suggestionCounts[expectedRecipe.Name]
		if !ok {
			t.Errorf("expected recipe %s not found in suggestions", expectedRecipe.Name)
			continue
		}
		if actualCount != 1 {
			t.Errorf("expected recipe %s to appear once in suggestions, but found %d occurrences", expectedRecipe.Name, actualCount)
		}
	}
	// check that the number of suggestions matches the expected number of recipes
	assert.Equal(t, len(suggestions), len(expectedRecipes))
}
