func sortedSquares(nums []int) []int {
    res := make([]int, len(nums))
	k := len(res) - 1

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	
	i, j := 0, len(nums) - 1
	for i <= j {
		if abs(nums[i]) > abs(nums[j]) {
			res[k] = nums[i]*nums[i]
			i++
		} else {
			res[k] = nums[j]*nums[j]
			j--
		}
		k--
	}

    return res
}