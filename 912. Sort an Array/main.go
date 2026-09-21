// The worst realisation - "Bubble Sort"
func sortArray(nums []int) []int {
    for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums) - i; j++ {
			if nums[j - 1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
    return nums
}


// Insertion Sort
func sortArray(nums []int) []int {
    for i := 1; i < len(nums); i++ {
		current := nums[i]
		j := i - 1

		for j >= 0 && nums[j] > current {
			nums[j + 1] = nums[j]
			j--
		}

		nums[j + 1] = current
	}
    return nums
}

// Selection Sort
func sortArray(nums []int) []int {
	for i := 0; i < len(nums) - 1; i++ {
		currentMin := i

		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[currentMin] {
				currentMin = j
			}
		}

		if currentMin != i {
			nums[i], nums[currentMin] = nums[currentMin], nums[i]
		}
	}
	return nums
}

// Counting Sort
func sortArray(nums []int) []int {
	var min, max int
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	frequencies := make([]int, max-min+1)
	for _, num := range nums {
		frequencies[num-min]++
	}

	// like prefix sums
	for i := 1; i < len(frequencies); i++ {
		frequencies[i] = frequencies[i-1] + frequencies[i]
	}

	// right shift
	var prev, temp int
	for i := 0; i < len(frequencies); i++ {
		temp = frequencies[i]
		frequencies[i] = prev
		prev = temp
	}

	sorted := make([]int, len(nums))
	for _, num := range nums {
		sorted[frequencies[num-min]] = num
		frequencies[num]++
	}

	return sorted
}
