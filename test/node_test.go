package test

import (
	"testing"

	"github.com/ThomasLee94/autosuggest/node"
	"github.com/stretchr/testify/assert"
)

func TestNodeStructAttributes(t *testing.T) {
	// init node struct
	nodeChar := "A"
	nodeStruct := node.NewNode(nodeChar)

	// node.Character == "A"
	assert.Equal(t, nodeStruct.Character, nodeChar)

	// node.childred == empty dict
	assert.Equal(t, len(nodeStruct.Children), 0)

	// node.terminal == false
	assert.Equal(t, nodeStruct.Terminal, false)

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
	assert.Equal(t, nodeStructA.HasChildren(nodeCharB), true)

	// test get children
	childNode, _ := nodeStructA.GetChildren(nodeCharB)
	assert.Equal(t, childNode.Character, nodeCharB)

}
