func pivotIndex(nums []int) int {
    prefixSums := make([]int, len(nums))

	for i, num := range nums {
		if i == 0 {
			prefixSums[i] = num
		} else {
			prefixSums[i] = num + prefixSums[i-1]
		}
	}

	prevSum := 0
    rightEdge := prefixSums[len(prefixSums) - 1]
	for i, sum := range prefixSums {
		if prevSum == rightEdge - sum {
			return i
		}
        prevSum = sum
	}

	return -1
}