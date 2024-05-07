package array_queue

import "testing"

func TestQueueNew(t *testing.T) {
	queue := NewArrayQueue[int]()

	if queue == nil {
		t.Error("Queue is not initialized")
	}
}

func TestQueueEnqueue(t *testing.T) {
	queue := NewArrayQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	if queue.Size != 5 {
		t.Error("Queue Size should be 5")
	}

	queue.Enqueue(6)
	queue.Enqueue(7)

	if queue.Size != 7 {
		t.Error("Queue Size should be 7")
	}
}

func TestQueueDequeue(t *testing.T) {
	queue := NewArrayQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	if queue.Size != 5 {
		t.Error("Queue hasn't enqueued all the elements properly")
	}

	dequeuedElement := queue.Dequeue()

	if queue.Size != 4 {
		t.Error("Queue hasn't dequeued all the elements properly'")
	}

	if dequeuedElement != 1 {
		t.Error("Dequeued Element should be 1")
	}
}

func TestQueuePeek(t *testing.T) {
	queue := NewArrayQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	peekedElement := queue.Peek()

	if peekedElement != 1 {
		t.Error("Peeked element should be 1")
	}

	queue.Clear()

	peekedElement = queue.Peek()

	if peekedElement != 0 {
		t.Error("Queue should be empty")
	}
}

func TestQueueClear(t *testing.T) {
	queue := NewArrayQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	if queue.Size != 5 {
		t.Error("Queue Size should be 5")
	}

	queue.Clear()

	if queue.Size != 0 {
		t.Error("Queue Size should be 0")
	}
}

func TestQueueIsEmpty(t *testing.T) {
	queue := NewArrayQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	isEmpty := queue.IsEmpty()

	if isEmpty {
		t.Error("Queue is not empty")
	}

	queue.Clear()

	isEmpty = queue.IsEmpty()

	if !isEmpty {
		t.Error("Queue is empty")
	}
}
