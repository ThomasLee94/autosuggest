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

	// test add children
	nodeSructA.AddChildren(nodeCharB, nodeStructB)

	# Create node 'A' and verify it does not have any children
	node_A = PrefixTreeNode('A')
	assert node_A.num_children() == 0
	assert node_A.has_child('B') is False
	# Verify getting child from node 'A' raises error
	with self.assertRaises(ValueError):
			node_A.get_child('B')
	# Create node 'B' and add it as child to node 'A'
	node_B = PrefixTreeNode('B')
	node_A.add_child('B', node_B)
	# Verify node 'A' has node 'B' as child
	assert node_A.num_children() == 1
	assert node_A.has_child('B') is True
	child_node = node_A.get_child('B')
	assert child_node is node_B
	# Verify adding node 'B' as child to node 'A' again raises error
	with self.assertRaises(ValueError):
			node_A.add_child('B', node_B)
}
