package structures

type Stack[T any] struct {
	array []T
}

func MakeStack[T any]() Stack[T] {
	return Stack[T]{make([]T, 0)}
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