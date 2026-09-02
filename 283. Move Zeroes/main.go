// O(n^2) solution - it's a bad one
func moveZeroes(nums []int)  {
	if len(nums) == 1 {
		return
	}

    i := 0
	for i < len(nums) {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
		i++
	}
}

// This one works for O(n) - it's a better one
func moveZeroes(nums []int)  {
	if len(nums) == 1 {
		return
	}

    lastNonZeroFoundAt := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[lastNonZeroFoundAt], nums[i] = nums[i], nums[lastNonZeroFoundAt]
            lastNonZeroFoundAt++
        }
    }
}