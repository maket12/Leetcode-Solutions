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
