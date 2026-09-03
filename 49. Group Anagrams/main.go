func groupAnagrams(strs []string) [][]string {
	const mapSize = 26  // since all strigns include only english lowercase letters

    frequenciesMap := make(map[[mapSize]int][]string, len(strs))
	for _, str := range strs {
		var counters [mapSize]int
		for i := range str {
			counters[str[i] - 'a']++
		}
		frequenciesMap[counters] = append(frequenciesMap[counters], str)
	}
	
	res := make([][]string, 0, len(frequenciesMap))
    for _, val := range frequenciesMap {
        res = append(res, val)
    }

	return res
}