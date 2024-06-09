package sort

import "testing"

func TestMerge(t *testing.T) {
	tests := [][]int{
		{},
		{1},
		{2, 2},
		{1, 2, 3, 4},
		{5, -9, 899, 0, 5, 1},
	}

	for _, test := range tests {
		t.Log("start", test)
		MergeSort(test)
		t.Log(test)
	}
}
