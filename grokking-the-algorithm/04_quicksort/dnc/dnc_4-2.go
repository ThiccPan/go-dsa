package dnc

type Node struct {
	Value int
	Next *Node
}

type List struct {
	Head *Node
}

func RecursiveLen(node *Node) int {
	listLen := 1
	// general case: next value not nil - recurse
	if node.Next != nil {
		listLen += RecursiveLen(node.Next)
	}
	// base case: no next value - return len
	return listLen
}