package maxpriorityqueue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	a := []int{5, 3, 2, 7, 19, 29, 32, 45, 95, 43, 23}
	pq := NewPriorityQueue(a)
	t.Run("Init", func(t *testing.T) {
		targetElement := 95
		err, r := pq.HeapMax()
		if err != nil {
			t.Errorf("got error:  %v", err)
		}
		if r != targetElement {
			t.Errorf("Max returns incorrect value, got: %d, want %d", r, targetElement)
		}
	})
	t.Run("InsertNonMax", func(t *testing.T) {
		pq.MaxHeapInsert(72)
		targetElement := 95
		err, r := pq.HeapMax()
		if err != nil {
			t.Errorf("got error:  %v", err)
		}
		if r != targetElement {
			t.Errorf("Max returns incorrect value, got: %d, want %d", r, targetElement)
		}
	})
	t.Run("InsertNewMax", func(t *testing.T) {
		targetElement := 112
		pq.MaxHeapInsert(targetElement)
		err, r := pq.HeapMax()
		if err != nil {
			t.Errorf("got error:  %v", err)
		}
		if r != targetElement {
			t.Errorf("Max returns incorrect value, got: %d, want %d", r, targetElement)
		}
	})
	t.Run("GetMaxElement", func(t *testing.T) {
		err, r := pq.HeapExtractMax()
		targetElement := 112
		if err != nil {
			t.Errorf("got error:  %v", err)
		}
		if r != targetElement {
			t.Errorf("Max returns incorrect value, got: %d, want %d", r, targetElement)
		}
	})
	t.Run("MaxElementPersists", func(t *testing.T) {
		err, r := pq.HeapExtractMax()
		targetElement := 95
		targetSize := 11
		if err != nil {
			t.Errorf("got error:  %v", err)
		}
		if r != targetElement {
			t.Errorf("Max returns incorrect value, got: %d, want %d", r, targetElement)
		}
		heapSize := pq.GetNumElements()
		if heapSize != targetSize {
			t.Errorf("wrong number of elements in heap, got %d, should be %d", heapSize, targetSize)
		}
	})
}
