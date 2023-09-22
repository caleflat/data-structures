package datastructures

// Stack is a data structure that follows the LIFO (Last In First Out) principle.
// It has the following methods:
// - Push: adds an element to the top of the stack
// - Pop: removes the top element of the stack and returns it
// - Peek: returns the top element of the stack without removing it
// - IsEmpty: returns true if the stack is empty, false otherwise
// - Size: returns the number of elements in the stack
type Stack[T any] struct {
	data []T
}

// NewStack creates a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

// Pop removes the top element of the stack and returns it
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		// workaround for returning a zero value of type T
		return *new(T), false
	}

	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return value, true
}

// Peek returns the top element of the stack without removing it
func (s *Stack[T]) Peek() *T {
	if s.IsEmpty() {
		return nil
	}

	value := s.data[len(s.data)-1]

	return &value
}

// IsEmpty returns true if the stack is empty, false otherwise
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.data)
}
