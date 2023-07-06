package dnc

func RecursiveSum(node *Node) int {
	sum := 0
	sum += node.Value
	// general case: next value not nil - recurse
	if node.Next != nil {
		sum += RecursiveSum(node.Next)
	}
	// base case: no next value - return sum
	return sum
}