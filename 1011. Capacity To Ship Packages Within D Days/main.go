func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMaxVals(weights []int) (int, int) {
	var maxValue, sumAll int
	
	for _, i := range weights {
		maxValue = max(maxValue, i)
		sumAll += i
	}

	return maxValue, sumAll
}

func canBeShippedWithin(weights []int, days, capacity int) int {
	var temp int

	for i := 0; i < len(weights); i++ {
		if temp + weights[i] > capacity {
			temp = 0
			days--
		}
		temp += weights[i]
	}

	if temp != 0 {
		days--
	}

	return days
}

func shipWithinDays(weights []int, days int) int {
	left, right := getMaxVals(weights)
	var minimum int = right

	for left <= right {
		middle := left + (right - left) / 2
		days := canBeShippedWithin(weights, days, middle)

		if days >= 0 {
			minimum = min(minimum, middle)
			right = middle - 1
		} else {
            left = middle + 1
		}
	}

	return minimum
}