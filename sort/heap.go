package sort

func HeapSort(nums []int) {
	buildHeap(nums)
	for i := len(nums) - 1; i > 0; i-- {
		// move the largest to the end
		nums[i], nums[0] = nums[0], nums[i]

		// root is modified, fix heap again
		heapify(nums[:i], 0)
	}
}

func buildHeap(nums []int) {
	// start from the last non-leaf node
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, i)
	}
}

// heapify left tree is heap, right tree is heap, make root tree is also heap
// this is a max heap
func heapify(nums []int, root int) {
	largest := root
	left := 2*root + 1
	right := 2*root + 2
	if len(nums) > left && nums[left] > nums[largest] {
		largest = left
	}

	if len(nums) > right && nums[right] > nums[largest] {
		largest = right
	}

	if largest == root {
		return
	}

	nums[largest], nums[root] = nums[root], nums[largest]
	// adjust the impacted subtree
	heapify(nums, largest)
}
