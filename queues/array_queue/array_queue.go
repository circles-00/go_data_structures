package array_queue

type ArrayQueueInterface[T comparable] interface {
	Enqueue(item T)
	Dequeue() T
	Peek() T
	Clear()
	IsEmpty() bool
}

type ArrayQueue[T comparable] struct {
	Elements []T
	Size     int
}

func (q *ArrayQueue[T]) Enqueue(item T) {
	q.Elements = append(q.Elements, item)
	q.Size += 1
}

func (q *ArrayQueue[T]) Dequeue() T {
	frontElement := q.Peek()

	q.Elements = q.Elements[1:]
	q.Size -= 1

	return frontElement
}

func (q *ArrayQueue[T]) Peek() T {
	if len(q.Elements) > 0 {
		return q.Elements[0]
	}

	var empty_result T

	return empty_result
}

func (q *ArrayQueue[T]) Clear() {
	q.Elements = make([]T, 0)
	q.Size = 0
}

func (q *ArrayQueue[T]) IsEmpty() bool {
	return q.Size == 0
}

func NewArrayQueue[T comparable]() *ArrayQueue[T] {
	return &ArrayQueue[T]{Elements: make([]T, 0), Size: 0}
}
