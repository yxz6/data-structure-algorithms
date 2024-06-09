package sort

import "testing"

func TestQuick1(t *testing.T) {
	tests := [][]int{
		{},
		{1},
		{2, 2},
		{1, 2, 3, 4},
		{5, -9, 899, 0, 5, 1},
	}

	for _, test := range tests {
		t.Log("start", test)
		QuickSortV1(test)
		t.Log(test)
	}
}

func TestQuick2(t *testing.T) {
	tests := [][]int{
		{},
		{1},
		{2, 2},
		{1, 2, 3, 4},
		{5, -9, 899, 0, 5, 1},
	}

	for _, test := range tests {
		t.Log("start", test)
		QuickSortV2(test)
		t.Log(test)
	}
}
