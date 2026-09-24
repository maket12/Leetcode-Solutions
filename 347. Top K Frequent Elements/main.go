func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	buckets := make([][]int, len(nums))
	for key, val := range freq {
		buckets[val-1] = append(buckets[val-1], key)
	}

	res := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0; i-- {
		if k <= 0 {
			break
		}
		for _, j := range buckets[i] {
			if k <= 0 {
				break
			}
			res = append(res, j)
			k--
		}
	}

	return res
}