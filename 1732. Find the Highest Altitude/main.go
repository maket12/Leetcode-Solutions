func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestAltitude(gain []int) int {
    var largest, current int
	for _, i := range gain {
		current += i
		largest = max(largest, current)
	}
	return largest
}