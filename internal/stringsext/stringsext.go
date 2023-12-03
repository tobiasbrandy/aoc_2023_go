package stringsext

import "strings"

func IsAsciiDigit(char uint8) bool {
	return char >= '0' && char <= '9'
}

func Reverse(in string) string {
	var sb strings.Builder
	runes := []rune(in)
	for i := len(runes) - 1; 0 <= i; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}
