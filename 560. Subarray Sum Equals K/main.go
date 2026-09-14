func subarraySum(nums []int, k int) int {
    prefixCount := make(map[int]int)
	prefixCount[0] = 1

	var current, count int
	for _, num := range nums {
		current += num

		target := current - k
        if freq, exists := prefixCount[target]; exists {
            count += freq
        }

		prefixCount[current]++
	}

	return count
}