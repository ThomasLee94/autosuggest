package trie

import "github.com/ThomasLee94/autosuggest/node"

// ******
// Struct
// ******
type Trie struct {
	root *node.Node
	size int
}

// ******************
// Trie class methods
// ******************
func (trie *Trie) IsEmpty() bool {
	// Return True if this prefix tree is empty (contains no strings).
	if trie.size == 0 {
		return true
	}

	return false
}
