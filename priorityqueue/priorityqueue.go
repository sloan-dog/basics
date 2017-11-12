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
		pq.MaxHeapify(i)
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

func (pq *PriorityQueue) MaxHeapify(i int) {
	l := left(i)
	r := right(i)
	var largest int
	if l < pq.heapSize && pq.arr[l] > pq.arr[i] {
		largest = l
	} else {
		largest = i
	}
	if r < pq.heapSize && pq.arr[r] > pq.arr[largest] {
		largest = r
	}
	if largest != i {
		pq.swap(i, largest)
		pq.MaxHeapify(largest)
	}
}

func (pq *PriorityQueue) swap(a, b int) {
	tmp := pq.arr[a]
	pq.arr[a] = pq.arr[b]
	pq.arr[b] = tmp
}

func (pq *PriorityQueue) HeapMax() (error, int) {
	if pq.heapSize < 1 {
		return errors.New("heap underflow"), -1
	}
	return nil, pq.arr[0]
}

func (pq *PriorityQueue) HeapExtractMax() (error, int) {
	if pq.heapSize < 1 {
		return errors.New("heap underflow"), -1
	}

	max := pq.arr[0]                  // maximum is first element
	pq.arr[0] = pq.arr[pq.heapSize-1] // swap last with first
	pq.heapSize--                     // decrement heapsize
	pq.arr = pq.arr[:pq.heapSize]     // forget that last element
	pq.MaxHeapify(0)                  // restore maxheap property
	return nil, max
}

func (pq *PriorityQueue) HeapIncreaseKey(i int, key int) error {
	if key < pq.arr[i] {
		return errors.New("new key is smaller than current key")
	}
	pq.arr[i] = key
	for i > 0 && pq.arr[parent(i)] < pq.arr[i] { // the parent is less than the key of the element were increasing
		p := parent(i) // get parent index
		pq.swap(i, p)  // swap them
		i = p          // our element is now at index == p
	}
	return nil
}

func (pq *PriorityQueue) MaxHeapInsert(key int) {
	pq.heapSize++
	pq.arr = append(pq.arr, key) // we always reduce slice size when popping element, so we always append when inserting
	pq.HeapIncreaseKey(pq.heapSize-1, key)
}

func (pq *PriorityQueue) GetNumElements() int {
	return pq.heapSize
}
