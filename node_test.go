package g3

import (
	"testing"
)

type nodeAddTest struct {
	node1 Node
	node2 Node
	want  Node
}

var nodeAddTests = []nodeAddTest{
	{
		Node{1, 2, 3},
		Node{4, 5, 6},
		Node{5, 7, 9},
	},
}

func TestAdd(t *testing.T) {
	for testInd, tt := range nodeAddTests {
		got := tt.node1.Add(tt.node2)
		if !got.Equals(tt.want) {
			t.Errorf("Test #%d: want: %v, got %v", testInd, tt.want, got)
		}
	}
}
