package stringsext

import (
	"slices"
)

func IsAsciiDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func Reverse(in string) string {
	inRunes := []rune(in)
	slices.Reverse(inRunes)
	return string(inRunes)
}
