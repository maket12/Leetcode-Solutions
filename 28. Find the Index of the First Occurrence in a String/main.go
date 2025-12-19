func strStr(haystack string, needle string) int {
	var haystackLen, needleLen int = len(haystack), len(needle)

	if needleLen == 0 {
		return 0
	}

	for i := 0; i <= haystackLen - needleLen; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}