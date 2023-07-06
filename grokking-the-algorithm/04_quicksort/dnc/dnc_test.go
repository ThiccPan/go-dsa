package dnc

import (
	"testing"
)

func TestRecSum(t *testing.T) {
	val := RecSum([]int{1, 2, 3, 4, 5}, 4)
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

func TestRecursiveSum(t *testing.T) {
	node3 := Node{Value: 3}
	node2 := Node{Value: 2, Next: &node3}
	node1 := Node{Value: 1, Next: &node2}
	sum := RecursiveSum(&node1)
	if sum != 6 {
		t.Error("nopers")
	}
	t.Log(sum)
}

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		arr      []int
		search   int
		expected int
	}{
		{
			desc: "success with odd arr len",
			arr: []int{1,2,3,4,5,6,7,8,9},
			search: 4,
			expected: 3,
		},
		{
			desc: "success with even arr len",
			arr: []int{1,2,3,4,5,6,8,9},
			search: 8,
			expected: 6,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lastIdx := len(tC.arr)-1
			searchIdx := RecursiveBSearch(tC.arr, tC.search, 0, lastIdx, lastIdx/2)
			if searchIdx != tC.expected {
				t.Error("got:", searchIdx, ", expected:", tC.expected)
			}
		})
	}
}
