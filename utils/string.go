package utils

func GetUniqueCharCounts(s string) map[string]int {
	present := make(map[string]int)
	for _, c := range []rune(s) {
		char := string(c)
		present[char]++
	}

	return present
}
