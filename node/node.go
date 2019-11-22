package node

import (
	"fmt"
)

// ******
// Struct
// ******
type Node struct {
	character string
	children  map[rune]*Node
	terminal  bool
}

// ***********
// Consturctor
// ***********

// Init node
func (node *Node) Init(character string) {
	node.character = character
	node.children = make(map[rune]*Node)
	node.terminal = false
}

// ******************
// Node class methods
// ******************

func (node *Node) IsTerminal() bool {
	// Return True if this prefix tree node terminates a string
	return node.terminal
}

func (node *Node) HasChildren(character rune) bool {
	// Return True if this prefix tree node has a child node that
	// represents the given character amongst its children.

	_, found := node.children[character]

	if found {
		return true
	}

	return false

}

func (node *Node) GetChildren(character rune) (*Node, error) {
	// Return this prefix tree node's child node that represents the given
	// character if it is amongst its children, or raise error if not.

	if node.HasChildren(character) {
		return node.children[character], nil
	}

	return nil, fmt.Errorf("No child exist for character: %g", character)

}

func (node *Node) AddChildren(character rune, childNode *Node) error {
	// Add the given character and child node as a child of this node, or
	// raise ValueError if given character is amongst this node's children.
	_, found := node.children[character]

	if !found {
		node.children[character] = childNode
	}

	return fmt.Errorf("Child already exists for character: %g", character)

}
