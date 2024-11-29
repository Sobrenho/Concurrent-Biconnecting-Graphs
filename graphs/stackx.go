package graphs

type StackX[T any] struct {
	array []T
}

func NewStack[T any]() *StackX[T] {
	stack := new(StackX[T])
	stack.array = make([]T, 0)
	return stack
}

func (stack *StackX[T]) Push(item T) {
	stack.array = append(stack.array, item)
}

func (stack *StackX[T]) Size() int {
	return len(stack.array)
}

func (stack *StackX[T]) Pop() T {
	element := stack.array[stack.Size()]
	stack.array = stack.array[:stack.Size() - 1]
	return element
}