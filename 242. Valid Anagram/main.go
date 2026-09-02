func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Frequency counters
    sMap := make(map[rune]int, len(s))
	tMap := make(map[rune]int, len(t))

	for _, r := range s {
		sMap[r]++
	}

	for _, r := range t {
		tMap[r]++
	}

	for key, sVal := range sMap {
		if tVal, ok := tMap[key]; !ok || sVal != tVal {
			return false
		}
	}

	return true
}