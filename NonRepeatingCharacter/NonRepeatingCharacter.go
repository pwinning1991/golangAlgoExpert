package main

func NonRepeatingCharacter(str string) int {
	for idx, char := range str {
		if countChars(char, str) {
			return idx
		}
	}
	return -1
}

func countChars(char rune, str string) bool {
	count := 0
	for _, char1 := range str {
		if char1 == char {
			count += 1
		}

		if count >= 2 {
			return false
		}
	}

	return true
}
