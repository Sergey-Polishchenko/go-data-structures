package singly

import (
	"testing"

	"github.com/Sergey-Polishchenko/go-data-structures/errors"
)

type listTest struct {
	name        string
	operation   func(*List[int]) (interface{}, error)
	expected    interface{}
	expectedErr error
}

var listTests = []listTest{
	// Add/Append Test
	{
		name: "Add to empty list",
		operation: func(l *List[int]) (interface{}, error) {
			l.Add(1)
			return l.Values(), nil
		},
		expected:    []int{1},
		expectedErr: nil,
	},
	{
		name: "Append multiple elements",
		operation: func(l *List[int]) (interface{}, error) {
			l.Append(1, 2, 3)
			return l.Values(), nil
		},
		expected:    []int{1, 2, 3},
		expectedErr: nil,
	},

	// Prepend Test
	{
		name: "Prepend to empty list",
		operation: func(l *List[int]) (interface{}, error) {
			l.Prepend(1)
			return l.Values(), nil
		},
		expected:    []int{1},
		expectedErr: nil,
	},

	// Get Test
	{
		name: "Get from empty list",
		operation: func(l *List[int]) (interface{}, error) {
			return l.Get(0)
		},
		expected:    0,
		expectedErr: errors.ErrIndexOutOfBounds,
	},
	{
		name: "Get valid index",
		operation: func(l *List[int]) (interface{}, error) {
			l.Add(1, 2, 3)
			return l.Get(1)
		},
		expected:    2,
		expectedErr: nil,
	},

	// Remove Test
	{
		name: "Remove from empty list",
		operation: func(l *List[int]) (interface{}, error) {
			return nil, l.Remove(0)
		},
		expected:    nil,
		expectedErr: errors.ErrEmptyList,
	},
	{
		name: "Remove first element",
		operation: func(l *List[int]) (interface{}, error) {
			l.Add(1, 2, 3)
			return nil, l.Remove(0)
		},
		expected:    []int{2, 3},
		expectedErr: nil,
	},

	// Contains Test
	{
		name: "Contains all elements",
		operation: func(l *List[int]) (interface{}, error) {
			l.Add(1, 2, 3)
			return l.Contains(1, 2), nil
		},
		expected:    true,
		expectedErr: nil,
	},
}

func TestList(t *testing.T) {
	for _, tt := range listTests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				list := New[int]()
				result, err := tt.operation(list)

				if err != tt.expectedErr {
					t.Errorf("Error: got %v, want %v", err, tt.expectedErr)
				}

				switch v := result.(type) {
				case []int:
					if !sliceEqual(v, tt.expected.([]int)) {
						t.Errorf("Values: got %v, want %v", v, tt.expected)
					}
				case int:
					if v != tt.expected.(int) {
						t.Errorf("Value: got %d, want %d", v, tt.expected)
					}
				case bool:
					if v != tt.expected.(bool) {
						t.Errorf("Bool: got %t, want %t", v, tt.expected)
					}
				}

				if tt.expectedErr == nil && tt.expected != nil {
					if values, ok := tt.expected.([]int); ok {
						if !sliceEqual(list.Values(), values) {
							t.Errorf("After Remove: got %v, want %v", list.Values(), values)
						}
					}
				}
			},
		)
	}
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
