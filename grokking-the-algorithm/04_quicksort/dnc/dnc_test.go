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
	
}