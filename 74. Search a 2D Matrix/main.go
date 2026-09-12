func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])

	var foundRow int = -1

	up, down := 0, m - 1
	for up <= down {
		middle := up + (down - up) / 2
		midEl := matrix[middle][0]

		if midEl == target {
			return true
		} else if midEl > target {
			down = middle - 1
		} else {
			foundRow = middle
			up = middle + 1
		}
	}

	if foundRow == -1 {
		return false
	}

	left, right := 1, n - 1
	for left <= right {
		middle := left + (right - left) / 2
		midEl := matrix[foundRow][middle]

		if midEl == target {
			return true
		} else if midEl < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return false
}