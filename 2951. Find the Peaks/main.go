func findPeaks(mountain []int) []int {
    answer := make([]int, 0)
	for i := 1; i < len(mountain) - 1; i++ {
		if mountain[i] > mountain[i - 1] && mountain[i] > mountain[i + 1] {
			answer = append(answer, i)
			i++
		}
	}
	return answer
}