func sortArrayByParity(nums []int) []int {
    res := make([]int, len(nums))
	left, right := 0, len(nums) - 1
	for _, i := range nums {
		if i % 2 == 0 {  // even
			res[left] = i
			left++
		} else {  // odd
			res[right] = i
			right--	
		}
	}
	return res
}