package datastructures

import (
	"testing"
)

func TestStack(t *testing.T) {
	cases := []struct {
		name string
		data []any
	}{
		{
			name: "empty stack",
			data: []any{},
		},
		{
			name: "non-empty stack",
			data: []any{1, 2, 3, 4, 5},
		},
		{
			name: "stack with nil values",
			data: []any{nil, nil, nil},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := NewStack[any]()
			if !s.IsEmpty() {
				t.Errorf("expected stack to be empty")
			}

			for _, v := range c.data {
				s.Push(v)
			}

			if s.Size() != len(c.data) {
				t.Errorf("expected stack size to be %d, got %d", len(c.data), s.Size())
			}

			for i := len(c.data) - 1; i >= 0; i-- {
				v, ok := s.Pop()
				if !ok || v != c.data[i] {
					t.Errorf("expected stack to pop %v, got %v", c.data[i], v)
				}
			}

			if !s.IsEmpty() {
				t.Errorf("expected stack to be empty")
			}
		})
	}
}
