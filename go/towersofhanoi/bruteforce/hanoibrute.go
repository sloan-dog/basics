package hanoibrute

import (
	"fmt"
	"os"
	"os/signal"
)

func _moveDisc(start int, target int, t [][]int) {
	// fmt.Println(t)
	l := len(t[start])
	// pop off disk
	tmp, r := t[start][l-1], t[start][:l-1]
	// set the nested slice a to r
	t[start] = r
	//append our popped element to target
	t[target] = append(t[target], tmp)
}

type HanoiState struct {
	towers [][]int
}

func (s HanoiState) GetTowers() [][]int {
	return s.towers
}

type StupidQueue struct {
	queue []HanoiState
}

func (q *StupidQueue) Enque(state HanoiState) {
	q.queue = append(q.queue, state)
}
func (q *StupidQueue) Deque() HanoiState {
	// grab the first el, re-ref the slice from 2nd index to end
	tmp, r := q.queue[0], q.queue[1:]
	q.queue = r
	return tmp
}

func (q *StupidQueue) Empty() bool {
	return len(q.queue) < 1
}

func SolveHanoiBrutal(nD int, t [][]int) HanoiState {

	memo := make(map[string]bool)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	counter := 0
	q := StupidQueue{
		queue: []HanoiState{},
	}
	q.Enque(HanoiState{towers: t})
	for q.Empty() != true {
		counter++
		// fmt.Println("itr: ", counter)
		s := q.Deque()
		fmt.Println(counter, " : ", s.towers)
		if s.isSolved(nD) {
			// we won!
			// return the solution
			return s
		}
		// oh snap...we were wrong...lets try something else
		// there are 6 options with every piece of state
		// a - b, a - c, b - a, b - c, c - a, c - b
		// LETS TRY EM ALL
		for _, state := range genHanoiActions(s, memo) {
			// PUT EM UP
			q.Enque(state)
		}
		select {
		case <-interrupt:
			return HanoiState{}
		default:
		}
	}
	return HanoiState{}
}

func genHanoiActions(s HanoiState, badStates map[string]bool) []HanoiState {
	// a - b , a - c
	l0 := len(s.towers[0])
	l1 := len(s.towers[1])
	l2 := len(s.towers[2])

	var l0HasElement bool = l0 > 0
	var l1HasElement bool = l1 > 0
	var l2HasElement bool = l2 > 0

	// ints are initialized to 0,
	// if we pass a zero to stackingorderpossible..we know the stack it came from is empty
	var l0El, l1El, l2El int

	if l0HasElement {
		l0El = s.towers[0][l0-1]
	}
	if l1HasElement {
		l1El = s.towers[1][l1-1]
	}
	if l2HasElement {
		l2El = s.towers[2][l2-1]
	}

	actions := []HanoiState{}
	if l0HasElement {
		if stackingOrderPossible(l0El, l1El) {
			actions = append(actions, moveAndCreateNewState(0, 1, s.towers))
		}
		if stackingOrderPossible(l0El, l2El) {
			actions = append(actions, moveAndCreateNewState(0, 2, s.towers))
		}
	}
	if l1HasElement {
		if stackingOrderPossible(l1El, l0El) {
			actions = append(actions, moveAndCreateNewState(1, 0, s.towers))
		}
		if stackingOrderPossible(l1El, l2El) {
			actions = append(actions, moveAndCreateNewState(1, 2, s.towers))
		}
	}
	if l2HasElement {
		if stackingOrderPossible(l2El, l0El) {
			actions = append(actions, moveAndCreateNewState(2, 0, s.towers))
		}
		if stackingOrderPossible(l2El, l1El) {
			actions = append(actions, moveAndCreateNewState(2, 1, s.towers))
		}
	}
	return actions
}

func moveAndCreateNewState(a int, b int, towers [][]int) HanoiState {
	// towers is copied from calling scope into this function scope, so its a new value
	// go things
	tmpOuter := make([][]int, len(towers))
	// deep copy the 2d array because of how slices are pointers to arrays
	for idx, _ := range towers {
		tmpInner := make([]int, len(towers[idx]))
		for i, val := range towers[idx] {
			tmpInner[i] = val
		}
		tmpOuter[idx] = tmpInner
	}
	_moveDisc(a, b, tmpOuter)
	return HanoiState{towers: tmpOuter}
}
func (s HanoiState) isSolved(n int) bool {
	if len(s.towers[0]) == 0 && len(s.towers[1]) == 0 && len(s.towers[2]) == n {
		return true
	}
	return false

}
func stackingOrderPossible(el int, topOfTower int) bool {
	if topOfTower == 0 {
		// other tower has no elements
		return true
	}
	if el > topOfTower {
		return false
	}
	return true
}
