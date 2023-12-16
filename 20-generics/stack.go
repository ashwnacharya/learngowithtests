package generics

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Pop() (T, bool) {

	if s.isEmpty() {
		var zero T
		return zero, false
	}

	lastIndex := len(s.elements) - 1
	lastElement := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return lastElement, true
}
