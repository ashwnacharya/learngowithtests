package recipe

import "github.com/ashwnacharya/working-without-mocks/domain/ingredients"


type MealType int

const (
	Breakfast MealType = iota
	Lunch
	Dinner
)

func (m MealType) String() string {
	return [...]string{"Breakfast", "Lunch", "Dinner"}[m]
}


type Recipe struct {
	Name string
	MealType MealType
	Ingredients []ingredients.Ingredient
}

