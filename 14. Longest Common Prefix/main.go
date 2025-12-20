import "strings"


func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
		return ""
	}

	var longestPrefix string = strs[0]
	for i := 0; i < len(longestPrefix); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || longestPrefix[i] != strs[j][i] {
				return longestPrefix[:i]
			}
		}
	}

	return longestPrefix
}
