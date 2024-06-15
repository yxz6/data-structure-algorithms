package kth_largest

func findKthLargest(nums []int, k int) int {
	pivot := partition(nums)
	if pivot == k-1 {
		return nums[pivot]
	}

	if pivot < k-1 {
		return findKthLargest(nums[pivot+1:], k-pivot-1)
	}

	return findKthLargest(nums[:pivot], k)
}

func partition(nums []int) int {
	n := len(nums)
	// nums[n-1] is pivot
	i := 0
	for j := 0; j < n-1; j++ {
		if nums[j] > nums[n-1] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[n-1], nums[i] = nums[i], nums[n-1]
	return i
}
