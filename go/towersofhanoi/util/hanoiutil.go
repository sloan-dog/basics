package hanoiutil

func BuildTowers(n int) [][]int {
	start := []int{}
	for i := n; i > 0; i-- {
		// e.g. { 5 , 4, 3, 2, 1}
		start = append(start, i)
	}
	towers := [][]int{
		start,
		[]int{},
		[]int{},
	}

	return towers
}
