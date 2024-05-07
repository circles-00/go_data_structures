package array_stack

import "testing"

func TestStackNew(t *testing.T) {
	stack := NewArrayStack[int]()

	if stack == nil {
		t.Error("Stack is not initialized")
	}
}

func TestStackPush(t *testing.T) {
	stack := NewArrayStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Size != 3 {
		t.Error("Stack Size should be 3")
	}

	stack.Push(4)
	stack.Push(5)

	if stack.Size != 5 {
		t.Error("Stack Size should be 5")
	}
}

func TestStackPop(t *testing.T) {
	stack := NewArrayStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Size != 3 {
		t.Error("Stack Size should be 3")
	}

	lastElement := stack.Pop()

	if lastElement != 3 {
		t.Error("Last element should be 3")
	}

	if stack.Size != 2 {
		t.Error("Stack Size should be 2")
	}
}

func TestStackPeek(t *testing.T) {
	stack := NewArrayStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Size != 3 {
		t.Error("Stack Size should be 3")
	}

	lastElement := stack.Peek()

	if lastElement != 3 {
		t.Error("Last element should be 3")
	}

	if stack.Size != 3 {
		t.Error("Stack Size should be 3")
	}
}

func TestStackClear(t *testing.T) {
	stack := NewArrayStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Size != 3 {
		t.Error("Stack Size should be 3")
	}

	stack.Clear()

	if stack.Size != 0 {
		t.Error("Stack Size should be 0")
	}
}

func TestStackIsEmpty(t *testing.T) {
	stack := NewArrayStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Size != 3 {
		t.Error("Stack Size should be 3")
	}

	isEmpty := stack.IsEmpty()

	if isEmpty {
		t.Error("Stack should not be empty")
	}

	stack.Clear()

	isEmpty = stack.IsEmpty()

	if !isEmpty {
		t.Error("Stack should be empty")
	}
}
