func mySqrt(x int) int {
    if x < 2 {
        return x
    }

	var answer int
	
    left, right := 1, x / 2
	for left <= right {
		middle := left + (right - left) / 2
		if middle <= x / middle {
            answer = middle
            left = middle + 1
        } else {
            right = middle - 1
        }
	}

	return answer
}