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
			trie.insert(char)
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
func (trie *Trie) Contains(character string) bool {

	foundNode = trie.findNode(word)

	// If node is terminal and not
	if foundNode.IsTerminal() && foundNode != nil {
		return true
	}
	return false
}
