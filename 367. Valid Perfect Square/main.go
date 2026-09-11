func isPerfectSquare(num int) bool {
    if num < 2 {
        return true
    }

    left, right := 1, num / 2

	for left <= right {
		middle := left + (right - left) / 2
		if middle * middle == num {
			return true
		} else if middle * middle < num {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return false
}