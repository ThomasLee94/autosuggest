package node

type Node struct {
	character string
	children  map[string]Node
	terminal  bool
}
