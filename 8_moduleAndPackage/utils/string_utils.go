package utils

import "strconv"

func StringReverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func AddAndGreet(a, b int) string {
	return strconv.Itoa(a + b)
}
