package planner_test

import (
	"context"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/ashwnacharya/working-without-mocks/adapters/persistence/inmemory"
	"github.com/ashwnacharya/working-without-mocks/adapters/persistence/sqlite"
	"github.com/ashwnacharya/working-without-mocks/domain/ingredients"
	"github.com/ashwnacharya/working-without-mocks/domain/planner"
	"github.com/ashwnacharya/working-without-mocks/domain/recipe"
)

type RecipeMatcherTest struct {
	NewIngredientStore func() planner.CloseableIngredientStore
	NewRecipeBook      func() planner.CloseableRecipeBook
}

func TestRecipeMatcher(t *testing.T) {
	t.Run("with in memory store", func(t *testing.T) {
		RecipeMatcherTest {
			NewRecipeBook: func() planner.CloseableRecipeBook {
				return inmemory.NewRecipeStore()
			},
			NewIngredientStore: func() planner.CloseableIngredientStore {
				return inmemory.NewIngredientStore()
			},
		}.Test(t)
	})

	t.Run("with sqlite", func(t *testing.T) {
		if !testing.Short() {
			RecipeMatcherTest{
				NewRecipeBook: func() planner.CloseableRecipeBook {
					return inmemory.NewRecipeStore()
				},
				NewIngredientStore: func() planner.CloseableIngredientStore {
					return sqlite.NewIngredientStore()
				},
			}.Test(t)
		}
	})
}

type CloseableStore interface {
	planner.IngredientStore
	Close()
}

func (r RecipeMatcherTest) Test(t *testing.T) {	
	t.Run("if we have no ingredients we can't make anything", func(t *testing.T) {
		store := r.NewIngredientStore()
		t.Cleanup(store.Close)
		recipeBook := r.NewRecipeBook()
		assert.NoError(t, recipeBook.AddRecipes(context.Background(), bananaBread, bananaMilkshake))
		t.Cleanup(recipeBook.Close)
		assertAvailableRecipes(t, store, recipeBook, []recipe.Recipe{})
	})

	t.Run("if we have the ingredients for banana bread we can make it", func(t *testing.T) {
		store := r.NewIngredientStore()
		t.Cleanup(store.Close)

		recipeBook := r.NewRecipeBook()
		assert.NoError(t, recipeBook.AddRecipes(context.Background(), bananaBread, bananaMilkshake))
		t.Cleanup(recipeBook.Close)

		assert.NoError(t, store.Store(
			context.Background(),
			ingredients.Ingredient{Name: "Bananas", Quantity: 2},
			ingredients.Ingredient{Name: "Flour", Quantity: 1},
			ingredients.Ingredient{Name: "Eggs", Quantity: 2},
		))
		assertAvailableRecipes(t, store, recipeBook, []recipe.Recipe{bananaBread})
	})

	t.Run("if we have bananas and milk, we can make banana milkshake", func(t *testing.T) {
		store := r.NewIngredientStore()
		t.Cleanup(store.Close)

		recipeBook := r.NewRecipeBook()
		assert.NoError(t, recipeBook.AddRecipes(context.Background(), bananaBread, bananaMilkshake))
		t.Cleanup(recipeBook.Close)

		assert.NoError(t, store.Store(
			context.Background(),
			ingredients.Ingredient{Name: "Bananas", Quantity: 2},
			ingredients.Ingredient{Name: "Milk", Quantity: 1},
		))
		assertAvailableRecipes(t, store, recipeBook, []recipe.Recipe{bananaMilkshake})
	})

	t.Run("if we have ingredients for banana bread and milkshake, we can make both", func(t *testing.T) {
		store := r.NewIngredientStore()
		t.Cleanup(store.Close)

		recipeBook := r.NewRecipeBook()
		assert.NoError(t, recipeBook.AddRecipes(context.Background(), bananaBread, bananaMilkshake))
		t.Cleanup(recipeBook.Close)

		assert.NoError(t, store.Store(
			context.Background(),
			ingredients.Ingredient{Name: "Bananas", Quantity: 2},
			ingredients.Ingredient{Name: "Flour", Quantity: 1},
			ingredients.Ingredient{Name: "Eggs", Quantity: 2},
			ingredients.Ingredient{Name: "Milk", Quantity: 1},
		))
		assertAvailableRecipes(t, store, recipeBook, []recipe.Recipe{bananaMilkshake, bananaBread})
	})

}

func assertAvailableRecipes(t *testing.T, ingredientStore planner.IngredientStore, recipeStore planner.RecipeBook, expectedRecipes []recipe.Recipe) {
	t.Helper()
	suggestions, _ := planner.New(recipeStore, ingredientStore).SuggestRecipes(context.Background())

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
