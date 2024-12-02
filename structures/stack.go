package structures

type Stack[T any] struct {
	array []T
}

func NewStack[T any]() *Stack[T] {
	stack := new(Stack[T])
	stack.array = make([]T, 0)
	return stack
}

func (stack *Stack[T]) Push(item T) {
	stack.array = append(stack.array, item)
}

func (stack *Stack[T]) Size() int {
	return len(stack.array)
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *Stack[T]) Pop() T {
	element := stack.array[stack.Size()-1]
	stack.array = stack.array[:stack.Size() - 1]
	return element
}