package sort

import "testing"

func TestBubble(t *testing.T) {
	tests := [][]int{
		{},
		{1},
		{2, 2},
		{1, 2, 3, 4},
		{5, -9, 899, 0, 5, 1},
	}

	for _, test := range tests {
		BubbleSort(test)
		t.Log(test)
	}
}
