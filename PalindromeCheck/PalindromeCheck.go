package main

func PalindromeCheck(str string) bool {
	reverse_idx := len(str) - 1
	for idx := range str {
		if str[idx] != str[reverse_idx] {
			return false
		}
		reverse_idx -= 1
	}
	return true
}
