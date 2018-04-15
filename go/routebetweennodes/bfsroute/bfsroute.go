package bfsroute

type GraphNode struct {
	data int
	adj  []*GraphNode
}

func NewGraphNode(val int) *GraphNode {
	return &GraphNode{
		data: val,
		adj:  make([]*GraphNode, 0, 10),
	}
}

func (gn *GraphNode) addAdj(n *GraphNode) {
	gn.adj = append(gn.adj, n)
}

type SimpleQueue struct {
	arr []*GraphNode
}

func (s *SimpleQueue) enque(n *GraphNode) {
	s.arr = append(s.arr, n)
}

func (s *SimpleQueue) deque() *GraphNode {
	r, tmp := s.arr[len(s.arr)-1], s.arr[:len(s.arr)-1]
	s.arr = tmp
	return r
}

func (s SimpleQueue) isEmpty() bool {
	return len(s.arr) < 1
}

func IsGraphRoutePossible(n *GraphNode, target *GraphNode) bool {
	// 0 length, 20 cap
	q := &SimpleQueue{
		arr: make([]*GraphNode, 0, 10),
	}
	m := make(map[*GraphNode]bool)
	q.enque(n)
	for !q.isEmpty() {
		n = q.deque()
		if n == target {
			return true
		}
		for _, adj := range (*n).adj {
			// if its not in the map, add it to the queue, mark it seen in the map
			if exists, _ := m[n]; exists != true {
				// add to the queue
				q.enque(adj)
				// mark that we've been here before
				m[n] = true
			}
		}
	}
	return false
}
