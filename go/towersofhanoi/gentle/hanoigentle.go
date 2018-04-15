package hanoigentle

func MoveDisc(start int, target int, tP *[][]int) {
	t := *tP
	// fmt.Println(t)
	l := len(t[start])
	// pop off disk
	tmp, r := t[start][l-1], t[start][:l-1]
	// set the nested slice a to r
	t[start] = r
	//append our popped element to target
	t[target] = append(t[target], tmp)
}

func SolveHanoi(nD int, a int, b int, c int, tP *[][]int) {
	if nD == 1 {
		MoveDisc(a, c, tP)
	} else if nD == 2 {
		// move from start to free
		MoveDisc(a, b, tP)
		// move from start to target
		MoveDisc(a, c, tP)
		// move from buffer to target
		MoveDisc(b, c, tP)
	} else {
		// move from n - 1 disks from start to buffer
		SolveHanoi(nD-1, a, c, b, tP)
		// move 1 disk from start to target
		SolveHanoi(1, a, b, c, tP)
		// move n - 1 disks from buffer to target
		SolveHanoi(nD-1, b, a, c, tP)
	}
}
