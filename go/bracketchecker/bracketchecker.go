package main

import (
	"fmt"
	"time"
)

type SimpleQueue struct {
	arr []int
}

func NewSimpleQueue(length int) *SimpleQueue {
	// length of string we are analyzing
	// we wont need the go runtime to double and realloc/copy our underlying array
	arr := make([]int, 0, length)
	return &SimpleQueue{
		arr: arr,
	}
}

func (q *SimpleQueue) enque(idx int) {
	q.arr = append(q.arr, idx)
}

func (q *SimpleQueue) deque() int {
	tmp, r := q.arr[0], q.arr[1:]
	q.arr = r
	return tmp
}

func (q *SimpleQueue) notEmpty() bool {
	return len(q.arr) > 0
}

func getNearestBracketLeft(start int, s string) int {
	i := start
	for i > -1 && s[i] != '[' {
		i--
	}
	return i
}

func getMatchingRightBracket(start int, s string) int {
	i := start
	q := new(SimpleQueue)
	q.enque(i)
	i++
	for q.notEmpty() {
		if s[i] == '[' {
			q.enque(i)
		} else if s[i] == ']' {
			q.deque()
		}
		i++
	}
	return i - 1
}

func getSurroundingBracketIndexes(cursor int, s string) (int, int) {
	firstLeftBracketIdx := getNearestBracketLeft(cursor, s)
	closingRightBracketIdx := getMatchingRightBracket(firstLeftBracketIdx, s)
	return firstLeftBracketIdx, closingRightBracketIdx
}

func getRightBracketIndexNoQueue(start int, s string) int {
	count := 1
	i := start + 1
	for count > 0 {
		if s[i] == '[' {
			count++
		} else if s[i] == ']' {
			count--
		}
		i++
	}
	return i - 1
}

func getSurroundingBracketIndexesNoQueue(cursor int, s string) (int, int) {
	firstLeftBracketIdx := getNearestBracketLeft(cursor, s)
	rightBracket := getRightBracketIndexNoQueue(firstLeftBracketIdx, s)
	return firstLeftBracketIdx, rightBracket
}

func main() {
	s := "abc[defg[gf[gd[gf]fisuhsfifshiushdiuhfiuhfiuhidhdfghgfddfgdfggd[dgdfgdfgfdg]dfgdgfgdfgdfgdg[dffgdfgdfgdfg]dfgdfgdfgdfg[dfgdfgdfg]dfggh]gh]ijhddfgg]"
	idx := 10

	t0 := time.Now()
	r1_left, r1_right := getSurroundingBracketIndexes(idx, s)
	t1 := time.Now()
	tSub1 := t1.Sub(t0)

	fmt.Println(r1_left, r1_right, tSub1)

	t2 := time.Now()
	r2_left, r2_right := getSurroundingBracketIndexesNoQueue(idx, s)
	t3 := time.Now()
	tSub2 := t3.Sub(t2)

	fmt.Println(r2_left, r2_right, tSub2)

}
