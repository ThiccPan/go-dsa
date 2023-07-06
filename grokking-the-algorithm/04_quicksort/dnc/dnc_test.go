package dnc

import "testing"

func TestRecSum(t *testing.T) {
	val := RecSum([]int{1,2,3,4,5}, 4)
	expected := 15
	if val != expected {
		t.Error("nopers")
	}
	t.Log("yeppers")
}

func TestRecursiveLen(t *testing.T) {
	node3 := Node{Value: 3}
	node2 := Node{Value: 2, Next: &node3}
	node1 := Node{Value: 1, Next: &node2}
	listLen := RecursiveLen(&node1)
	if listLen != 3 {
		t.Error("nopers")
	}
	t.Log(listLen)
}