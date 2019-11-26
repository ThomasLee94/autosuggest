package nodetest

import (
	"testing"

	"github.com/ThomasLee94/autosuggest/node"
)

func TestTableNodeStruct(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
		error    string
	}{
		{"A", "A", "NewNode('A') is {}; want {}"},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := node.NewNode(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
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
