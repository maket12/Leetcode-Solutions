func singleNumber(nums []int) int {
	var result int
	for _, i := range nums {
		result = result ^ i
	}
	return result
}