// package trie contains class functions that will be used to
//create autosuggests in terminal.
package trie

import (
	"github.com/ThomasLee94/autosuggest/completion"
	"github.com/ThomasLee94/autosuggest/node"
)

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

// NewTrie - init trie
func NewTrie(wordsOrChars ...string) *Trie {
	var trie *Trie = new(Trie)
	trie.Root = node.NewNode("")
	trie.Size = 0
	// Insert each string, if any were given
	for _, element := range wordsOrChars {
		trie.Insert(string(element))
	}

	return trie
}

/* -------------------------------------------------------------------------- */
/*                             TRIE CLASS METHODS                             */
/* -------------------------------------------------------------------------- */

// IsEmpty - return true if this prefix tree is empty.
func (trie *Trie) IsEmpty() bool {
	if trie.Size == 0 {
		return true
	}
	return false
}

// Contains - return true if this prefix tree contains the given string.
func (trie *Trie) Contains(word string) bool {

	foundNode, foundFunc := trie.FindNode(word)

	// If node is terminal and not
	if foundFunc {
		if foundNode.IsTerminal() && foundNode != nil {
			return true
		}
	}

	return false
}

// FindNode - return the node that terminates the given string
// in this prefix tree, or if the given string is not completely
// found, return nil.Search is done iteratively with a loop
// starting from the root node.
func (trie *Trie) FindNode(word string) (*node.Node, bool) {
	node := trie.Root
	foundFunc := true

	// case: empty string
	if len(word) == 0 {
		return node, foundFunc
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
			foundFunc = false
			break
		}
	}

	return node, foundFunc
}

// Insert the given string into this prefix tree.
func (trie *Trie) Insert(word string) {
	nodeObj, foundFunc := trie.FindNode(word)
	// case: node already exists & is a terminal
	if foundFunc {
		if nodeObj.Terminal {
			return
		}

	}

	nodeObj = trie.Root
	for _, char := range word {
		_, found := nodeObj.Children[string(char)]

		// case: if the letter does not exist as a child from current node
		if !found {
			newChildNode := node.NewNode(string(char))
			nodeObj.AddChildren(string(char), newChildNode)
			// traverse tree
		}
		nodeObj = nodeObj.Children[string(char)]

	}

	// set node terminal to true at the end of word iteration
	nodeObj.Terminal = true

	trie.Size++
}

// AppendSlice - appends prefix to completions slice
func AppendSlice(completions []string, prefix string) []string {
	completions = append(completions, prefix)
	return completions
}

// Complete - return a list of all strings stored in this
// prefix tree that start with the given prefix string using
// completions struct.
func (trie *Trie) Complete(wordOrPrefix string) []string {

	// slice for completions
	var completions completion.Completion

	node, foundFunc := trie.FindNode(wordOrPrefix)

	// case: prefix does not exist
	if !foundFunc {
		if node == nil {
			return completions.CompleteSlice
		}
	}

	// case: wordOrPrefix is already a completed word
	if node.IsTerminal() {
		completions.CompleteSlice = append(completions.CompleteSlice, wordOrPrefix)
	}
	// traverse through prefix tree & append all terminal words
	for _, childNode := range node.Children {
		trie.traverse(childNode, wordOrPrefix+childNode.Character, &completions)
	}

	return completions.CompleteSlice

}

// Strings - return a list of all strings stored in this trie.
func (trie *Trie) Strings() []string {
	// all strings list
	var allStrings completion.Completion

	for _, node := range trie.Root.Children {

		trie.traverse(node, node.Character, &allStrings)

	}

	return allStrings.CompleteSlice
}

// Traverse this prefix tree with recursive depth-first traversal.
// Start at the given node and append prefixes that are terminal.
func (trie *Trie) traverse(node *node.Node, prefix string, completions *completion.Completion) {
	tempComplete := completions.CompleteSlice
	// execute visit if it is terminal
	if node.IsTerminal() {
		tempComplete = append(tempComplete, prefix)
		completions.CompleteSlice = tempComplete
	}

	for _, childNode := range node.Children {
		// concat chars
		trie.traverse(childNode, prefix+childNode.Character, completions)
	}

}
