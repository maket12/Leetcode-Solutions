func searchRange(nums []int, target int) []int {
    start, end := -1, -1

	left, right := 0, len(nums) - 1
	for left <= right {
		middle := left + (right - left) / 2
		if nums[middle] == target {
			start = middle
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	if start == -1 {
		return []int{-1, -1}
	}

	left, right = start, len(nums) - 1
	for left <= right {
		middle := left + (right - left) / 2
		if nums[middle] == target {
			end = middle
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		}
	}

	return []int{start, end}
}