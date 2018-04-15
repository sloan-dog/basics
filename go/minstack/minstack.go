package minstack

import "errors"

type MSNode struct {
	value       int
	previousMin int
}

func newMSNode(value int, previousMin int) MSNode {
	n := MSNode{
		value:       value,
		previousMin: previousMin,
	}
	return n
}

type MinStack struct {
	arr    []MSNode
	cap    int
	min    int
	topIdx int
}

func NewMinStack() *MinStack {
	return &MinStack{
		arr:    make([]MSNode, 10),
		cap:    10,
		min:    -1,
		topIdx: -1,
	}
}

func (ms *MinStack) isFull() bool {
	return ms.topIdx+1 == ms.cap
}

func (ms *MinStack) isEmpty() bool {
	return ms.topIdx < 0
}

func (ms *MinStack) makeSpace() {
	l := len(ms.arr)
	c := cap(ms.arr)
	// make a doubled slice
	newSlice := make([]MSNode, l*2, c*2)
	// copy the elements
	copy(newSlice, ms.arr)
	ms.arr = newSlice
}

func (ms *MinStack) push(val int) {
	if ms.isFull() {
		ms.makeSpace()
	}

	n := newMSNode(val, ms.min)

	if val < ms.min || ms.min == -1 {
		ms.min = val
	}

	ms.topIdx++
	ms.arr[ms.topIdx] = n
}

func (ms *MinStack) pop() (int, error) {
	if ms.isEmpty() {
		return -1, errors.New("pop on empty minstack")
	}
	n := ms.arr[ms.topIdx]
	ms.arr = ms.arr[:ms.topIdx]
	ms.min = n.previousMin
	return n.value, nil
}

func (ms *MinStack) GetMin() (int, error) {
	if ms.isEmpty() {
		return -1, errors.New("get on empty minstack")
	}
	return ms.min, nil
}
