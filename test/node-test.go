package nodeTest

import (
	"testing"

	"github.com/ThomasLee94/autosuggest/node"
)

func TestNodeInit(t *testing.T) {
	// init node struct
	node := node.Node("")
	nodeChar := "A"
	nodeObj := node.Init(nodeChar)

	// node.Character == "A"
	if nodeObj.Character != "A" {
		t.Errorf("Node Init('A') = %d; want A", nodeChar)
	}
	// node.childred == empty dict
	if len(nodeObj.childred) != 0 {
		t.Errorf("Length of node is = %d; want 0", len(nodeObj.children))
	}
	// node.terminal == false
	if nodeObj.terminal != false {
		t.Errorf("Node terminal is %d; want false", nodeObj.terminal)
	}
}
