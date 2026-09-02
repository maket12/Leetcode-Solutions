func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int, len(nums))
	for i, n := range nums {
		numsMap[n] = i
	}

	for i, n := range nums {
		if ind, ok := numsMap[target-n]; ok && i != ind {
			return []int{i, ind}
		}
	}
	
	return nil
}