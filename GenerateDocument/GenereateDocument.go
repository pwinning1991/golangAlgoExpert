package main

func GenerateDocument(characters string, document string) bool {
	for _, char := range document {
		documentFrequency := countCharFrequency(char, document)
		charFrequency := countCharFrequency(char, characters)
		if documentFrequency > charFrequency {
			return false
		}
	}

	return true

}

func countCharFrequency(char rune, target string) int {
	frequnency := 0
	for _, char1 := range target {
		if char1 == char {
			frequnency += 1
		}
	}
	return frequnency

}
