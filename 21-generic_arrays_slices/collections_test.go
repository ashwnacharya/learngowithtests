package generic_arrays_slices

import (
	"testing"
	"strings"
)

func TestReduce(t *testing.T) {

	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(acc, x int) int { 
			return acc * x 
		}

		AssertEqual(t, Reduce([]int{1, 2, 3, 4}, multiply, 1), 24)
	})

	t.Run("concatenate strings" , func(t *testing.T) {
		concatenate := func(acc, x string) string { 
			return acc + x 
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})

	type Person struct {
		Name string
	}
	
	t.Run("Find the best programmer", func(t *testing.T) {
		people := []Person{
			Person{Name: "Kent Beck"},
			Person{Name: "Martin Fowler"},
			Person{Name: "Chris James"},
		}
	
		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})
	
		AssertTrue(t, found)
		AssertEqual(t, king, Person{Name: "Chris James"})
	})
}

