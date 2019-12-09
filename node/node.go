// package node contains class functions that will be used in package
// trie to build a trie with children nodes.
package node

import (
	"fmt"
)

/* -------------------------------------------------------------------------- */
/*                                   STRUCT                                   */
/* -------------------------------------------------------------------------- */

type Node struct {
	Character string
	Children  map[string]*Node
	Terminal  bool
}

/* -------------------------------------------------------------------------- */
/*                                 CONSTRUCTOR                                */
/* -------------------------------------------------------------------------- */

func NewNode(char string) *Node {
	var node Node
	node.Character = char
	node.Children = make(map[string]*Node)
	node.Terminal = false
	return &node
}

/* -------------------------------------------------------------------------- */
/*                             NODE CLASS METHODS                             */
/* -------------------------------------------------------------------------- */

// IsTerminal - return true if this prefix
// tree node terminates a string
func (node *Node) IsTerminal() bool {
	return node.Terminal
}

// HasChildren - Return True if this prefix tree node
// has a child node that represents the given character
// amongst its children.
func (node *Node) HasChildren(character string) bool {

	_, found := node.Children[character]

	if found {
		return true
	}

	return false

}

// GetChildren - Return this prefix tree node's child node that
// represents the given character if it is amongst its children,
// or raise error if not.
func (node *Node) GetChildren(character string) (*Node, error) {

	if node.HasChildren(character) {
		return node.Children[character], nil
	}

	return nil, fmt.Errorf("No child exist for character: %s", character)

}

// AddChildren - Add given character and child node as a child of
// this node, or raise error if character is already a child
func (node *Node) AddChildren(character string, childNode *Node) error {
	_, found := node.Children[character]

	if !found {
		node.Children[character] = childNode
	}

	return fmt.Errorf("Child already exists for character: %s", character)

}
