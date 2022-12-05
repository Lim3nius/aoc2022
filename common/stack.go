package common

type Stack[T any] []T

func (s *Stack[T]) Pop() T {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack[T]) PopN(n int) []T {
	v := (*s)[len(*s)-n:]
	(*s) = (*s)[:len(*s)-n]
	return v
}

func (s *Stack[T]) Push(t ...T) {
	*s = append(*s, t...)
}
