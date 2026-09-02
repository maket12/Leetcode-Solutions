func findContentChildren(g []int, s []int) int {
    slices.Sort(g)
    slices.Sort(s)

    counter := 0
    lastAssignedCookieAt := 0
    for i := 0; i < len(g); i++ {
        for lastAssignedCookieAt < len(s) {
			if g[i] <= s[lastAssignedCookieAt] {
				counter++
				lastAssignedCookieAt++
				break
			}
			lastAssignedCookieAt++
		}
    }

    return counter
}