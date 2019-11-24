package nodetest

import (
	"testing"

	"github.com/ThomasLee94/autosuggest/node"
)

func TestNodeStructAttributes(t *testing.T) {
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

func TestNodeChildMethods(t *testing.T) {
	// init node struct
	nodeCharA := "A"
	nodeCharB := "B"

	// node A
	nodeStructA := node.Node{}
	nodeStructA.Character = nodeCharA

	// node B
	nodeStructB := node.Node{}
	nodeStructB.Character = nodeCharB

	// test add & get children
	nodeStructA.AddChildren(nodeCharB, nodeStructB)
	if !nodeStructA.HasChildren(nodeCharB) {
		t.Errorf("Node HasChildren is %d; want %d", nodeStructA.HasChildren(nodeCharB), true)
	}

	// test get children
	childNode, _ := nodeStructA.GetChildren(nodeCharB)
	if childNode.Character != nodeCharB {
		t.Errorf("Node GChildren is %d; want %d", childNode.Character, nodeCharB)
	}
}
