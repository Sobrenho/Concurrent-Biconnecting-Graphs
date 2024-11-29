package graphs


type stack[T any] []T

func (s stack[T]) Push(item T) stack[T] {
	return append(s, item)
}

func (s stack[T]) Pop() (T, stack[T]) {
	l := len(s)

	return s[l-1], s[:l-1]
}