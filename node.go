package node

import "errors"

// ******
// Struct
// ******
type Node struct {
	character string
	children  map[rune]*Node
	terminal  bool
}

// ******************
// Node class methods
// ******************

func (node *Node) isTerminal() bool {
	// Return True if this prefix tree node terminates a string
	return node.terminal
}

func (node *Node) hasChildren(character rune) bool {
	// Return True if this prefix tree node has a child node that
	// represents the given character amongst its children.

	_, found := node.children[character]

	if found {
		return true
	}

	return false

}

func (node *Node) getChildren(character rune) (*Node, error) {
	// Return this prefix tree node's child node that represents the given
	// character if it is amongst its children, or raise error if not.

	if node.hasChildren(character) {
		return node.children[character], nil
	}

	return -1, errors.New("No child exist for character: ", character)
}

func (node *Node) addChildren(character rune, childNode *Node) {
	// Add the given character and child node as a child of this node, or
	// raise ValueError if given character is amongst this node's children.
	_, found := node.children[character]

	if !found {
		node.children[character] = childNode
	}

	return -1, errors.New("Child already exists for character: ", character)

}
