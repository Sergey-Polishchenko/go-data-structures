package singly

// node represents an element in the list.
type node[T comparable] struct {
	value T
	next  *node[T]
}
