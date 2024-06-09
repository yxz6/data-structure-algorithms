package sort

func BubbleSort(nums []int) {
	n := len(nums)
	for round := 1; round <= n-1; round++ {
		for i := 0; i < n-round; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
}
