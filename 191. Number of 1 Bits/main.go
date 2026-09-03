func hammingWeight(n int) int {
	var counter int
	for n > 0 {
		counter += n & 1
		n >>= 1
	}
	return counter
}