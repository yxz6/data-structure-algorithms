package sort

func QuickSortV1(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	i := partitionV1(nums)
	QuickSortV1(nums[:i])
	QuickSortV1(nums[i+1:])
}

func partitionV1(nums []int) int {
	n := len(nums)
	pivot := nums[n-1]
	i := 0
	for j := 0; j < n-1; j++ {
		if nums[j] < pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[i], nums[n-1] = nums[n-1], nums[i]
	return i
}

func QuickSortV2(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	i := partitionV2(nums)
	QuickSortV2(nums[:i+1]) // pivot is included
	QuickSortV2(nums[i+1:])
}

func partitionV2(nums []int) int {
	n := len(nums)
	pivot := nums[0]
	left, right := -1, n
	for {
		left++
		for nums[left] < pivot {
			left++
		}

		right--
		for nums[right] > pivot {
			right--
		}

		if left >= right {
			return right
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
}
