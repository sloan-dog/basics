package main

import "fmt"

type Point struct {
	// type to represent a coord
	row int
	col int
}

type RobotPath struct {
	pointMap PointMap
	grid     [][]int
	path     []Point
}

type PointMap map[string]int

func (pm PointMap) Insert(row int, col int) {
	key := genKey(row, col)
	_, exists := pm[key]
	if exists {
		pm[key]++
	} else {
		pm[key] = 1
	}
}

func (pm PointMap) Get(row int, col int) (int, bool) {
	// because our return value is tuple
	// and go will not automatically return two values
	// from underlying map call
	// manually create two vars
	val, exists := pm[genKey(row, col)]
	return val, exists
}

func genKey(row int, col int) string {
	return fmt.Sprintf("%d,%d", row, col)
}

func (rp *RobotPath) _getPath(row int, col int) bool {

	// mark that this row,col pair has been checked
	// need to understand how adding a map for some dyn prog
	// does not enhance runtime
	// my instinct is that this algorithm is returning the shortest path
	// and thus returns the shortest path before other more costly routes are explored
	rp.pointMap.Insert(row, col)

	if row < 0 || col < 0 || rp.grid[row][col] == -1 {
		return false
	}

	var isOrigin bool = (row == 0 && col == 0)

	if isOrigin || rp._getPath(row-1, col) || rp._getPath(row, col-1) {
		rp.path = append(rp.path, Point{row: row, col: col})
		return true
	}

	return false
}

func buildGrid() [][]int {
	return [][]int{
		{0, 0, 0, -1, 0},
		{0, 0, 0, 0, 0},
		{0, -1, -1, -1, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, -1, 0},
	}
}

func (rp *RobotPath) getPath() *[]Point {
	if rp._getPath(4, 4) {
		return &rp.path
	}
	return nil
}

func main() {
	rp := RobotPath{
		grid:     buildGrid(),
		pointMap: PointMap{},
		path:     []Point{},
	}
	r := rp.getPath()
	fmt.Println(*r)
	// fmt.Println(rp.pointMap)
}
