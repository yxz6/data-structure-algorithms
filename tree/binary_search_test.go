package tree

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 9, 10, 20, 30}
	t.Log(BinarySearch(nums, 1))
	t.Log(BinarySearch(nums, 10))
	t.Log(BinarySearch(nums, 100))
}
