package nodeTest

import (
	"testing"

	"github.com/ThomasLee94/autosuggest/node"
)

func TestNodeStruct(t *testing.T) {
	// init node struct
	nodeStruct := node.Node{}
	nodeChar := "A"
	nodeStruct.Character = nodeChar

	// node.Character == "A"
	if nodeStruct.Character != "A" {
		t.Errorf("Node Init('A') = %d; want A", nodeChar)
	}
	// node.childred == empty dict
	if len(nodeStruct.Children) != 0 {
		t.Errorf("Length of node is = %d; want 0", len(nodeStruct.Children))
	}
	// node.terminal == false
	if nodeStruct.Terminal != false {
		t.Errorf("Node terminal is %d; want false", nodeStruct.Terminal)
	}
}
