func isIsomorphic(s string, t string) bool {
    mapSToT := make(map[byte]byte, len(s))
    mapTToS := make(map[byte]byte, len(t))

    for i := 0; i < len(s); i++ {
        charS := s[i]
        charT := t[i]

        if mappedT, ok := mapSToT[charS]; ok {
			if mappedT != charT {
				return false
			}
		} else {
			mapSToT[charS] = charT
		}

		if mappedS, ok := mapTToS[charT]; ok {
			if mappedS != charS {
				return false
			}
		} else {
			mapTToS[charT] = charS
		}
    }

    return true
}