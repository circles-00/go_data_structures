package array_stack

type ArrayStackInterface[T comparable] interface {
	Push(item T)
	Pop() T
	Peek() T
	Clear()
	IsEmpty() bool
}

type ArrayStack[T comparable] struct {
	Elements []T
	Size     int
}

func (s *ArrayStack[T]) Push(item T) {
	s.Elements = append(s.Elements, item)
	s.Size += 1
}

func (s *ArrayStack[T]) Pop() T {
	lastElement := s.Peek()
	s.Elements = s.Elements[:len(s.Elements)-1]
	s.Size -= 1

	return lastElement
}

func (s *ArrayStack[T]) Peek() T {
	return s.Elements[len(s.Elements)-1]
}

func (s *ArrayStack[T]) Clear() {
	s.Elements = make([]T, 0)
	s.Size = 0
}

func (s *ArrayStack[T]) IsEmpty() bool {
	return s.Size == 0
}

func NewArrayStack[T comparable]() *ArrayStack[T] {
	return &ArrayStack[T]{Elements: make([]T, 0), Size: 0}
}
