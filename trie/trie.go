package trie

import "github.com/ThomasLee94/autosuggest/node"


/* -------------------------------------------------------------------------- */
/*                                   STRUCT                                   */
/* -------------------------------------------------------------------------- */

type Trie struct {
	Root *node.Node
	Size int
}

/* -------------------------------------------------------------------------- */
/*                                 CONSTRUCTOR                                */
/* -------------------------------------------------------------------------- */


// Init trie
func NewTrie(wordsOrChars ...string) *Trie{
	var trie Trie
	trie.Root = node.NewNode("")
	trie.Size = 0
	// Insert each string, if any were given
	if len(wordsOrChars) > 0 {
		for _, wordOrChar := range wordsOrChars {
			trie.Insert(wordOrChar)
		}
	}
	return &trie
}


/* -------------------------------------------------------------------------- */
/*                             TRIE CLASS METHODS                             */
/* -------------------------------------------------------------------------- */


// IsEmpty - eturn true if this prefix tree is empty.
func (trie *Trie) IsEmpty() bool {
	if trie.Size == 0 {
		return true
	}
	return false
}

// Contains - return rrue if this prefix tree contains the given string.
func (trie *Trie) Contains(word string) bool {

	foundNode := trie.FindNode(word)

	// If node is terminal and not
	if foundNode.IsTerminal() && foundNode != nil {
		return true
	}
	return false
}

// FindNode - return the node that terminates the given string
// in this prefix tree, or if the given string is not completely
// found, return nil.Search is done iteratively with a loop
// starting from the root node.
func (trie *Trie) FindNode(word string) *node.Node {
	node := trie.Root

	// case: empty string
	if len(word) == 0 {
		return node
	}

	// iterate through word by char
	for _, char := range word {
		// iterate through children of current node
		_, found := node.Children[string(char)]
		if found {
			// traverse through children
			node = node.Children[string(char)]
		} else {
			node = nil
			break
		}
	}

	return node
}

// Insert the given string into this prefix tree.
func (trie *Trie) Insert(word string) {
	node := node.FindNode(word)

	// case: node already exists & is a terminal
	if node && node.Terminal {
		return
	}

	node = trie.Root

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
	node.Terminal = true

	trie.Size++
}

// Complete - return a list of all strings stored in this
// prefix tree that start with the given prefix string.
func (trie *Trie) Complete(wordOrPrefix string) []string {

	// slice for completions
	completions := []string{}

	node := trie.FindNode(wordOrPrefix)

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
		visit := append(completions, wordOrPrefix+childNode.character)
		trie.traverse(childNode, wordOrPrefix+child.character, visit)
	}

	return completions

}

// Strings - return a list of all strings stored in this trie.
func (trie *Trie) Strings() []string {
	// all strings list
	allStrings := []string

	for _, node := range trie.Root.Children {
		if node {
			traverse(child, child.character, func(allStrings []string, prefix string))
		}
	}

	return allStrings
}

// Traverse this prefix tree with recursive depth-first traversal.
// Start at the given node and visit each node with the given function.
func (trie *Trie) traverse(node *Node, prefix string, visit func(completions []string, prefix string)) {

	// execute visit if it is terminal
	if node.IsTerminal() {
		visit(completions, prefix)
	}

	for _, childNode := range node.Children {
		// concat chars
    trie.traverse(childNode, prefix + childNode.character, visit)
	}
}
