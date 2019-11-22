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

// IsTerminal - return true if this prefix
// tree node terminates a string
func (node *Node) IsTerminal() bool {
	return node.terminal
}

// HasChildren - Return True if this prefix tree node
// has a child node that represents the given character
// amongst its children.
func (node *Node) HasChildren(character rune) bool {

	_, found := node.children[character]

	if found {
		return true
	}

	return false

}

// GetChildren - Return this prefix tree node's child node that
// represents the given character if it is amongst its children,
// or raise error if not.
func (node *Node) GetChildren(character rune) (*Node, error) {

	if node.HasChildren(character) {
		return node.children[character], nil
	}

	return nil, fmt.Errorf("No child exist for character: %g", character)

}

// AddChildren - Add given character and child node as a child of
// this node, or raise error if character is already a child
func (node *Node) AddChildren(character rune, childNode *Node) error {
	_, found := node.children[character]

	if !found {
		node.children[character] = childNode
	}

	return fmt.Errorf("Child already exists for character: %g", character)

}
