package test

import (
	"testing"

	"github.com/ThomasLee94/autosuggest/node"
)

func TestNodeStructAttributes(t *testing.T) {
	// init node struct
	nodeChar := "A"
	nodeStruct := node.NewNode(nodeChar)

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
	// chars
	nodeCharA := "A"
	nodeCharB := "B"

	// init node structs
	nodeStructA := node.NewNode(nodeCharA)
	nodeStructB := node.NewNode(nodeCharB)

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
