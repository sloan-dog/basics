package maxpriorityqueue

import (
	"errors"
)

type PriorityQueue struct {
	arr      []int
	heapSize int
}

func NewPriorityQueue(arr []int) *PriorityQueue {
	// pq is pointer to PriorityQueue
	pq := new(PriorityQueue)
	pq.arr = arr
	pq.heapSize = len(arr)
	pq.buildHeap()
	return pq
}

func (pq *PriorityQueue) buildHeap() {
	for i := pq.heapSize / 2; i >= 0; i-- {
		pq.MinHeapify(i)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1 // (i + 1) * 2 - 1
}

func right(i int) int {
	return 2*i + 2 // ( i + 1 ) * 2 - 1 + 1
}

func (pq *PriorityQueue) MinHeapify(i int) {
	l := left(i)
	r := right(i)
	var smallest int
	if l < pq.heapSize && pq.arr[l] < pq.arr[i] {
		smallest = l
	} else {
		smallest = i
	}
	if r < pq.heapSize && pq.arr[r] < pq.arr[smallest] {
		smallest = r
	}
	if smallest != i {
		pq.swap(i, smallest)
		pq.MinHeapify(smallest)
	}
}

func (pq *PriorityQueue) swap(a, b int) {
	tmp := pq.arr[a]
	pq.arr[a] = pq.arr[b]
	pq.arr[b] = tmp
}

func (pq *PriorityQueue) HeapMin() (error, int) {
	if pq.heapSize < 1 {
		return errors.New("heap underflow"), -1
	}
	return nil, pq.arr[0]
}

func (pq *PriorityQueue) HeapExtractMin() (error, int) {
	if pq.heapSize < 1 {
		return errors.New("heap underflow"), -1
	}

	min := pq.arr[0]                  // minimum is first element
	pq.arr[0] = pq.arr[pq.heapSize-1] // swap last with first
	pq.heapSize--                     // decrement heapsize
	pq.arr = pq.arr[:pq.heapSize]     // forget that last element
	pq.MinHeapify(0)                  // restore minheap property
	return nil, min
}

func (pq *PriorityQueue) HeapDecreaseKey(i int, key int) error {
	if key < pq.arr[i] {
		return errors.New("new key is larger than current key")
	}
	pq.arr[i] = key
	// bubble it up
	for i > 0 && pq.arr[parent(i)] > pq.arr[i] { // the parent is greater than the key of the element were increasing
		p := parent(i) // get parent index
		pq.swap(i, p)  // swap them
		i = p          // our element is now at index == p
	}
	return nil
}

func (pq *PriorityQueue) MinHeapInsert(key int) {
	pq.heapSize++
	pq.arr = append(pq.arr, key) // we always reduce slice size when popping element, so we always append when inserting
	// adjust the key of the last element
	pq.HeapDecreaseKey(pq.heapSize-1, key)
}

func (pq *PriorityQueue) GetNumElements() int {
	return pq.heapSize
}
