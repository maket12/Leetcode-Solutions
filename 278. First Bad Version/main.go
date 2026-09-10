/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    var lastGoodVersion int
	for lastGoodVersion < n {
		middle := (n + lastGoodVersion) / 2
		if !isBadVersion(middle) {
			lastGoodVersion = middle + 1
		} else {
			n = middle
		}
	}
	return lastGoodVersion
}