package minstack

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	stack := NewMinStack()
	fmt.Println("IFUSHIUDhuidfHUIF")
	t.Run("Push", func(t *testing.T) {
		stack.push(5)
		stack.push(4)
		stack.push(3)
		stack.push(7)
		stack.push(2)
		r, err := stack.GetMin()
		if err != nil {
			t.Error(err)
		}
		if r != 2 {
			t.Errorf("got %d expected 2", r)
		}
	})
	t.Run("Pop", func(t *testing.T) {
		r, err := stack.pop()
		if err != nil {
			t.Error(err)
		}
		if r != 2 {
			t.Errorf("got %d expected 2", r)
		}
		m, err := stack.GetMin()
		if err != nil {
			t.Error(err)
		}
		if m != 3 {
			t.Errorf("got %d expected 3", m)
		}
	})
}
