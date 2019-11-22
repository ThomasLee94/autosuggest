package trie

import "github.com/thomaslee94/autosuggest/node"

type Trie struct {
	root *node.Node
	size int
}
