package sort

func MergeSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	middle := n / 2
	left := nums[0:middle]
	right := nums[middle:]
	MergeSort(left)
	MergeSort(right)

	res := make([]int, 0, n)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	copy(nums, res)
}
