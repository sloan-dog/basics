package bfsroute

import (
	"testing"
)

func TestBfsRoute(t *testing.T) {
	// construct graph
	n1 := NewGraphNode(1)
	n2 := NewGraphNode(2)
	n3 := NewGraphNode(3)
	n1.addAdj(n2)
	n1.addAdj(n3)
	n2.addAdj(n3)
	n4 := NewGraphNode(4)
	n5 := NewGraphNode(7)
	n6 := NewGraphNode(8)
	n3.addAdj(n4)
	n3.addAdj(n5)
	n4.addAdj(n5)
	n5.addAdj(n6)

	t.Run("RoutePossible", func(t *testing.T) {
		result := IsGraphRoutePossible(n1, n6)
		if result != true {
			t.Errorf("got %b expected 'true'", result)
		}
	})
}
