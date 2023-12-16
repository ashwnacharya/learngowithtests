package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("Asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("Asserting on strings", func(t *testing.T) {
		AssertEqual(t, "a", "a")
		AssertNotEqual(t, "a", "b")
	})
}

func TestStack(t *testing.T) {
	t.Run("Integer stack", func(t *testing.T) {
		
		myStackOfInts := new(Stack[int])

		AssertTrue(t, myStackOfInts.isEmpty())

		myStackOfInts.Push(1)
		AssertFalse(t, myStackOfInts.isEmpty())

		myStackOfInts.Push(2)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 2)

		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 1)

		AssertTrue(t, myStackOfInts.isEmpty())

		myStackOfInts.Push(3)
		myStackOfInts.Push(4)

		value1, _ := myStackOfInts.Pop()
		value2, _ := myStackOfInts.Pop()

		AssertEqual(t, value1+value2, 7)
	})
}