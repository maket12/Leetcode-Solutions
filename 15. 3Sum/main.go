func threeSum(nums []int) [][]int {
    res := make([][]int, 0)

	slices.Sort(nums)

	for i := 0; i < len(nums) - 2; i++ {
        if i != 0 && nums[i] == nums[i-1] {
            continue
        }
		
		target := -1 * nums[i]
		left, right := i + 1, len(nums) - 1
		for left < right {
			if nums[left] + nums[right] == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--

                for left < right && nums[left] == nums[left-1] {
                    left++
                }

                for left < right && nums[right] == nums[right+1] {
                    right--
                }

			} else if nums[left] + nums[right] > target {
				right--
			} else {
				left++
			}
		}
	}

	return res
}