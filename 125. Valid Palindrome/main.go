import "strings"

func isPalindrome(s string) bool {
    sLow := strings.ToLower(s)

    isLetter := func(c byte) bool {
        return c >= 'a' && c <= 'z'
    }
    isDigit := func(c byte) bool {
        return c >= '0' && c <= '9'
    }
    isRight := func(c byte) bool {
        return isLetter(c) || isDigit(c)
    }

    i, j := 0, len(s) - 1
    for i < j {
        if isRight(sLow[i]) && isRight(sLow[j]) {
            if sLow[i] != sLow[j] {
                return false
            } else {
                i++
                j--
            }
        } else {
            if !isRight(sLow[i]) {
                i++
            }
            if !isRight(sLow[j]) {
                j--
            }
        }
    }
    return true
}