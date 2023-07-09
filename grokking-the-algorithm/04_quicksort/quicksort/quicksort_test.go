package quicksort

import "testing"

func TestQuicksort(t *testing.T) {
	testCases := []struct {
		desc     string
		value    []int
		expected []int
	}{
		{
			desc: "success",
			value: []int{4,2,3,1,3,5,3},
			expected: []int{1,2,3,3,3,4,5},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			sorted :=Quicksort(tC.value)
			t.Log(sorted)
			for i := 0; i < len(sorted); i++ {
				if sorted[i] != tC.expected[i] {
					t.Error("invalid sorting")
				}
			}
		})
	}
}
