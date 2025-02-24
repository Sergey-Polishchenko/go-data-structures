// Package singly implements a singly-linked list.
//
// The list is not thread-safe and supports generic comparable types.
// It provides standard operations such as Add, Remove, Get, and more,
// with error handling for edge cases.
//
// Example:
//
//	list := singly.New[int]()
//	list.Add(1, 2, 3)
//	value, err := list.Get(0) // Returns 1, nil
package singly

import (
	"github.com/Sergey-Polishchenko/go-data-structures/errors"
	"github.com/Sergey-Polishchenko/go-data-structures/linkedlist"
)

// Assert List implementation for checkout List implementation
var _ linkedlist.List[int] = (*List[int])(nil)

// List represents a singly-linked list.
// It maintains references to the first and last nodes, and the total size.
// T is a comparable type constraint.
type List[T comparable] struct {
	first *node[T]
	last  *node[T]
	size  int
}

// New creates a new empty List. Optional initial values can be provided.
// Time complexity: O(1) (or O(n) if values are provided).
func New[T comparable](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Size returns the number of elements in the list.
// Time complexity: O(1).
func (list *List[T]) Size() int {
	return list.size
}

// Add appends one or more values to the end of the list.
// Time complexity: O(n) where n is the number of values added.
func (list *List[T]) Add(values ...T) {
	for _, value := range values {
		elem := &node[T]{value: value}
		if list.size == 0 {
			list.first = elem
			list.last = elem
		} else {
			list.last.next = elem
			list.last = elem
		}
		list.size += 1
	}
}

// Append is an alias for Add (appends values to the end).
func (list *List[T]) Append(values ...T) {
	list.Add(values...)
}

// Prepend adds values to the beginning of the list.
// Time complexity: O(n) where n is the number of values prepended.
func (list *List[T]) Prepend(values ...T) {
	for i := len(values) - 1; i >= 0; i -= 1 {
		elem := &node[T]{value: values[i], next: list.first}
		list.first = elem
		if list.size == 0 {
			list.last = elem
		}
		list.size += 1
	}
}

// First returns the head node value.
// Returns ErrEmptyList if called on an empty list.
// Time complexity: O(1).
func (list *List[T]) First() (T, error) {
	if list.IsEmpty() {
		var t T
		return t, errors.ErrEmptyList
	}
	return list.first.value, nil
}

// Last returns the tail node value.
// Returns ErrEmptyList if called on an empty list.
// Time complexity: O(1).
func (list *List[T]) Last() (T, error) {
	if list.IsEmpty() {
		var t T
		return t, errors.ErrEmptyList
	}
	return list.last.value, nil
}

// Get returns the value at the specified index.
// Returns error if index is out of bounds [0, size).
// Time complexity: O(n).
func (list *List[T]) Get(index int) (T, error) {
	if !list.inBounds(index) {
		var t T
		return t, errors.ErrIndexOutOfBounds
	}

	node := list.first
	for i := 0; i != index; i, node = i+1, node.next {
	}

	return node.value, nil
}

// Values returns a slice of all values in the list.
// Time complexity: O(n).
func (list *List[T]) Values() []T {
	values := make([]T, list.size)
	for i, node := 0, list.first; node != nil; i, node = i+1, node.next {
		values[i] = node.value
	}

	return values
}

// IndexOf returns the first index of the specified value.
// Returns -1 and ErrElementNotFound if value not found.
// Returns ErrEmptyList if called on an empty list.
// Time complexity: O(n).
func (list *List[T]) IndexOf(value T) (int, error) {
	if list.IsEmpty() {
		return -1, errors.ErrEmptyList
	}

	index := 0
	for node := list.first; node != nil; node = node.next {
		if node.value == value {
			return index, nil
		}
		index += 1
	}

	return -1, errors.ErrElementNotFound
}

// IsEmpty checks if the list has no elements.
// Time complexity: O(1).
func (list *List[T]) IsEmpty() bool {
	return list.size == 0
}

// Clear removes all elements from the list.
// Time complexity: O(1).
func (list *List[T]) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

// Remove deletes the element at the specified index.
// Returns ErrIndexOutOfBounds if index is invalid.
// Returns ErrEmptyList if called on an empty list.
// Time complexity: O(n).
func (list *List[T]) Remove(index int) error {
	if list.IsEmpty() {
		return errors.ErrEmptyList
	}
	if !list.inBounds(index) {
		return errors.ErrIndexOutOfBounds
	}

	var prev *node[T]
	node := list.first
	for i := 0; i != index; i, node = i+1, node.next {
		prev = node
	}

	if node == list.first {
		list.first = node.next
	}
	if node == list.last {
		list.last = prev
	}
	if prev != nil {
		prev.next = node.next
	}

	list.size -= 1

	return nil
}

// Contains checks if all specified values exist in the list.
// Returns true if all values are present, false otherwise.
// If no values are provided, returns true (empty set is always a subset).
// If the list is empty and values are provided, returns false.
// Time complexity: O(n + m) where n is list size and m is the number of values.
func (list *List[T]) Contains(values ...T) bool {
	if len(values) == 0 {
		return true
	}
	if list.size == 0 || list.size < len(values) {
		return false
	}

	valuesToFind := make(map[T]bool)
	for _, value := range values {
		valuesToFind[value] = true
	}

	for node := list.first; node != nil; node = node.next {
		if valuesToFind[node.value] {
			delete(valuesToFind, node.value)
			if len(valuesToFind) == 0 {
				return true
			}
		}
	}

	return false
}

// Swap swaps values of list nodes by their index.
// Returns ErrIndexOutOfBounds if index is out of range [0, size].
// Time complexity: O(n).
func (list *List[T]) Swap(i, j int) error {
	if !list.inBounds(i) || !list.inBounds(j) {
		return errors.ErrIndexOutOfBounds
	}

	if i == j {
		return nil
	}

	var node1, node2 *node[T]
	current := list.first
	for n := 0; current != nil && (node1 == nil || node2 == nil); n, current = n+1, current.next {
		if i == n {
			node1 = current
		}
		if j == n {
			node2 = current
		}
	}
	node1.value, node2.value = node2.value, node1.value

	return nil
}

// Insert adds one or more values at the specified index.
// If index is 0, the values are prepended to the list.
// If index equals the list size, the values are appended.
// Returns ErrIndexOutOfBounds if index is out of range [0, size].
// Time complexity: O(n).
func (list *List[T]) Insert(index int, values ...T) error {
	if !list.inBounds(index) {
		if index == list.size {
			list.Append(values...)
			return nil
		}
		return errors.ErrIndexOutOfBounds
	}

	if index == 0 {
		list.Prepend(values...)
		return nil
	}

	list.size += len(values)

	var prev *node[T]
	current := list.first
	for i := 0; i != index; i, current = i+1, current.next {
		prev = current
	}

	oldNext := prev.next
	for _, v := range values {
		newNode := &node[T]{value: v}
		prev.next = newNode
		prev = newNode
	}
	prev.next = oldNext

	return nil
}

// Set updates the value at the specified index.
// Returns ErrIndexOutOfBounds if index is invalid.
// Time complexity: O(n).
func (list *List[T]) Set(index int, value T) error {
	if !list.inBounds(index) {
		return errors.ErrIndexOutOfBounds
	}

	node := list.first
	for i := 0; i != index; i, node = i+1, node.next {
	}

	node.value = value

	return nil
}

// Check if index is in bounds [0, size).
func (list *List[T]) inBounds(index int) bool {
	return index >= 0 && index < list.size
}
