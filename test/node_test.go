package nodetest

import (
	"testing"

	"github.com/ThomasLee94/autosuggest/node"
)

func TestNodeStructAttributes(t *testing.T) {
	// init node struct
	nodeStruct := &node.Node{}
	nodeChar := "A"
	nodeStruct.Character = nodeChar

	// node.Character == "A"
	if nodeStruct.Character != "A" {
		t.Errorf("Node Init('A') is %s; want %s", nodeStruct.Character, nodeChar)
	}
	// node.childred == empty dict
	if len(nodeStruct.Children) != 0 {
		t.Errorf("Length of node is = %d; want %d", len(nodeStruct.Children), 0)
	}
	// node.terminal == false
	if nodeStruct.Terminal != false {
		t.Errorf("Node terminal is %t; want %t", nodeStruct.Terminal, false)
	}
}

func TestNodeChildMethods(t *testing.T) {
	// init node struct
	nodeCharA := "A"
	nodeCharB := "B"

	// node A
	nodeStructA := &node.Node{}
	nodeStructA.Character = nodeCharA

	// node B
	nodeStructB := &node.Node{}
	nodeStructB.Character = nodeCharB

	// test add & get children
	nodeStructA.AddChildren(nodeCharB, nodeStructB)
	if !nodeStructA.HasChildren(nodeCharB) {
		t.Errorf("Node HasChildren is %t; want %t", nodeStructA.HasChildren(nodeCharB), true)
	}

	// test get children
	childNode, _ := nodeStructA.GetChildren(nodeCharB)
	if childNode.Character != nodeCharB {
		t.Errorf("Node GChildren is %s; want %s", childNode.Character, nodeCharB)
	}
}
