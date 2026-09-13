func search(nums []int, target int) int {
    left, right := 0, len(nums) - 1

	for left <= right {
		middle := left + (right - left) / 2
		if nums[middle] == target {
			return middle
		} else if nums[left] <= nums[middle] {  // correct half
			if nums[left] <= target && target <= nums[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else {
			if nums[middle] <= target && target <= nums[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}

	return -1
}