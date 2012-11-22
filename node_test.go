package g3

import (
	"testing"
)

type nodeIsZeroTest struct {
	node Node
	want bool
}

var nodeIsZeroTests = []nodeIsZeroTest{
	{
		Node{0, 0, 0},
		true,
	},
	{
		Node{1, 0, 0},
		false,
	},
	{
		Node{0, 1, 0},
		false,
	},
	{
		Node{0, 0, 1},
		false,
	},
	{
		Node{1, 1, 1},
		false,
	},
}

func TestNodeIsZero(t *testing.T) {
	for testInd, tt := range nodeIsZeroTests {
		got := tt.node.IsZero()
		if got != tt.want {
			t.Errorf("Test #%d: want: %v, got %v", testInd, tt.want, got)
		}
	}
}

type nodeBinaryOpTest struct {
	node1 Node
	node2 Node
	want  Node
}

var nodeAddTests = []nodeBinaryOpTest{
	{
		Node{1, 2, 3},
		Node{4, 5, 6},
		Node{5, 7, 9},
	},
}

func TestAdd(t *testing.T) {
	for testInd, tt := range nodeAddTests {
		got := tt.node1.Add(tt.node2)
		if got != tt.want {
			t.Errorf("Test #%d: want: %v, got %v", testInd, tt.want, got)
		}
	}
}

var nodeSubTests = []nodeBinaryOpTest{
	{
		Node{1, 2, 3},
		Node{4, 5, 6},
		Node{-3, -3, -3},
	},
}

func TestSub(t *testing.T) {
	for testInd, tt := range nodeSubTests {
		got := tt.node1.Sub(tt.node2)
		if got != tt.want {
			t.Errorf("Test #%d: want: %v, got %v", testInd, tt.want, got)
		}
	}
}
