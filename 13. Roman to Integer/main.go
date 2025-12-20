func romanToInt(s string) int {
	var (
		result int
		prev string
	)

	frequency := make(map[string]int)
	for i := len(s) - 1; i >= 0; i-- {
		char := string(s[i])
		if char == "I" && prev == "V" {
			frequency["IV"]++
			frequency["V"]--
		} else if char == "I" && prev == "X" {
			frequency["IX"]++
			frequency["X"]--
		} else if char == "X" && prev == "L" {
			frequency["XL"]++
			frequency["L"]--
		} else if char == "X" && prev == "C" {
			frequency["XC"]++
			frequency["C"]--
		} else if char == "C" && prev == "D" {
			frequency["CD"]++
			frequency["D"]--
		} else if char == "C" && prev == "M" {
			frequency["CM"]++
			frequency["M"]--
		} else {
			frequency[char]++
		}
		prev = char
	}

	result += frequency["M"] * 1000
	result += frequency["CM"] * 900
	result += frequency["D"] * 500
	result += frequency["CD"] * 400
	result += frequency["C"] * 100
	result += frequency["XC"] * 90
	result += frequency["L"] * 50
	result += frequency["XL"] * 40
	result += frequency["X"] * 10
	result += frequency["IX"] * 9
	result += frequency["V"] * 5
	result += frequency["IV"] * 4
	result += frequency["I"] * 1

	return result
}