package trie

import "github.com/ThomasLee94/autosuggest/node"

// ******
// Struct
// ******
type Trie struct {
	root *node.Node
	size int
}

// ***********
// consturctor
// ***********

// Init trie
func (trie *Trie) Init(charOrString string) {
	trie.root = node.Node("")
	trie.size = 0
	// Insert each string, if any were given
	if len(charOrString) > 0 {
		for _, char := range charOrString {
			trie.insert(charOrString)
		}
	}
}

// ******************
// Trie class methods
// ******************

// IsEmpty - eturn true if this prefix tree is empty.
func (trie *Trie) IsEmpty() bool {
	if trie.size == 0 {
		return true
	}
	return false
}

// Contains - return rrue if this prefix tree contains the given string.
func (trie *Trie) Contains(word string) bool {

	foundNode = trie.findNode(word)

	// If node is terminal and not
	if foundNode.IsTerminal() && foundNode != nil {
		return true
	}
	return false
}

// findNode - return the node that terminates the given string
// in this prefix tree, or if the given string is not completely
// found, return nil.Search is done iteratively with a loop
// starting from the root node.
func (trie *Trie) findNode(word string) *node.Node {
	node := trie.root

	// case: empty string
	if len(word) == 0 {
		return node
	}

	// iterate through word by char
	for _, char := range word {
		// iterate through children of current node
		_, found := node.Children[char]
		if found {
			// traverse through children
			node = node.Children[char]
		} else {
			node = nil
			break
		}
	}

	return node
}

// Insert the given string into this prefix tree.
func (trie *Trie) Insert(word string) {
	node := node.findNode(word)

	// case: node already exists & is a terminal
	if node && node.terminal {
		return
	}

	node = trie.root

	for _, char := range word {
		_, found := node.Children[char]

		// case: if the letter does not exist as a child from current node
		if !found {
			newChildNode := node.Node(char)
			node.AddChildren(char, newChildNode)
			// traverse tree
		} else {
			node = node.children[char]
		}
	}

	// set node terminal to true at the end of word iteration
	node.terminal = true

	trie.size++
}

// Complete - return a list of all strings stored in this
// prefix tree that start with the given prefix string.
func (trie *Trie) Complete(wordOrPrefix string) []string {

	// slice for completions
	completions := []string{}

	node := trie.findNode(wordOrPrefix)

	// case: prefix does not exist
	if node != nil {
		return completions
	}

	// case: wordOrPrefix is already a completed word
	if node.isTerminal() {
		completions = append(completions, wordOrPrefix)
	}

	// traverse through prefix tree & append all terminal words
	for _, childNode := range node.Children {
		trie.traverse(childNode, wordOrPrefix+child.character, append(completions, wordOrPrefix+child.character))
	}

}
