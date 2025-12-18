func merge(nums1 []int, m int, nums2 []int, n int) {
	j, k := n-1, m+n-1
	m--
	for j >= 0 {
		if m >= 0 && nums1[m] > nums2[j] {
			nums1[k] = nums1[m]
			m--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}