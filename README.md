# Data structures in GoLang
This is just a repository for data structures that I'm usually using with GoLang
These data structures are my implementations and I'm most of the time using them for leetcode.

# Table of Content
- [Testing](#testing)
- [Usage](#usage)
   * [Queues](#queues)
      + [Array Queue](#array-queue)
         - [Initialization](#initialization)
         - [Enqueuing an Element](#enqueuing-an-element)
         - [Dequeuing an Element](#dequeuing-an-element)
         - [Peeking an Element](#peeking-an-element)
         - [Clearing the Queue](#clearing-the-queue)
         - [Checking whether the Queue is empty or not](#checking-whether-the-queue-is-empty-or-not)
         - [Checking the size of the Queue](#checking-the-size-of-the-queue)
         - [Accessing all the Elements of the Queue](#accessing-all-the-elements-of-the-queue)
   * [Stacks](#stacks)
      + [ArrayStack](#arraystack)
         - [Initialization](#initialization-1)
         - [Pushing an element to the stack](#pushing-an-element-to-the-stack)
         - [Poping an element from the stack](#poping-an-element-from-the-stack)
         - [Peeking an element from the stack](#peeking-an-element-from-the-stack)
         - [Clearing the stack](#clearing-the-stack)
         - [Checking whether the stack is empty](#checking-whether-the-stack-is-empty)
   * [Contributing](#contributing)

# Testing
For each data structure there is a test file where all the tests for that particular data structures are written.

For running all the tests please run the `test.sh` scripts from the root of the project.

```shell
./test.sh
```

# Usage
## Queues
### Array Queue

#### Initialization
So, I'm aiming for all my data structures to be generic, since I want to use them in different contexts.
While initializing the queue, you can use the utility function `NewArrayQueue` which will return a new instance of the `ArrayQueue` implementation
```go
queue := NewArrayQueue[int]()
```

#### Enqueuing an Element
```go
queue.Enqueue(1)
```

#### Dequeuing an Element
Dequeues the front Element and returns it
```go
dequeuedElement := queue.Dequeue()
```

#### Peeking an Element
Returns the front Element of the `ArrayQueue` without removing it.

```go
peekedElement := queue.Peek()
```

#### Clearing the Queue
Removes all the elements from the Queue.
```go
queue.Clear()
```

#### Checking whether the Queue is empty or not
```go
isEmpty := queue.IsEmpty()
```

#### Checking the size of the Queue
```go
size := queue.Size
```

#### Accessing all the Elements of the Queue
This will give you the slice instance that was created once initializing the `ArrayQueue`
```go
elements := queue.Elements
```

## Stacks
### ArrayStack

#### Initialization
While initializing the stack, you can use the utility function `NewArrayStack` which will return a new instance of the `ArrayStack` implementation

```go
stack := NewArrayStack[int]()
```

#### Pushing an element to the stack

```go
stack.Push(1)
```

#### Poping an element from the stack
Removes the last Element from the Stack and returns it

```go
poppedElement := stack.Pop()
```

#### Peeking an element from the stack
Returns the last Element from the Stack without removing it

```go
lastElement := stack.Peek()
```

#### Clearing the stack
Removes all the Elements from the stack

```go
stack.Clear()
```

#### Checking whether the stack is empty
```go
isEmpty := stack.IsEmpty()
```

# Contributing
I'm not looking for contributors, as this is just my hobby repository where I'm keeping all of the data structures that I need, but of course if you spot bugs, or any issues feel free to open a PR or issue, I'll be happy to review it.

