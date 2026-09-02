func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
    map1 := make(map[int]bool, len(nums1))
	map2 := make(map[int]bool, len(nums2))

	for _, i := range nums1 {
		map1[i] = true
	}
	for _, i := range nums2 {
		map2[i] = true
	}

	for key, _ := range map1 {
		if _, ok := map2[key]; ok {
			res = append(res, key)
		}
	}
    
	return res
}

// The second variation using "two-pointers" pattern
func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
    
	slices.Sort(nums1)
	slices.Sort(nums2)

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			if len(res) == 0 || nums1[i] != res[len(res)-1] {
				res = append(res, nums1[i])
			}
			i++
			j++
		}
	}
    
	return res
}
