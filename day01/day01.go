package day01

import (
	"github.com/tobiasbrandy/AoC_2022_go/internal/errexit"
	"github.com/tobiasbrandy/AoC_2022_go/internal/fileline"
	"regexp"
	"strings"
)

func isAsciiDigit(char uint8) bool {
	return char >= '0' && char <= '9'
}

func Part1(inputPath string) any {
	accum := 0

	fileline.ForEach(inputPath, errexit.HandleScanError, func(line string) {
		lineLen := len(line)

		for i := 0; i < lineLen; i++ {
			digit := line[i]
			if isAsciiDigit(digit) {
				accum += int(digit-'0') * 10
				break
			}
		}

		for i := lineLen - 1; i >= 0; i-- {
			digit := line[i]
			if isAsciiDigit(digit) {
				accum += int(digit - '0')
				break
			}
		}
	})

	return accum
}

var digitNames = [...]string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func parseDigitMatch(match string) uint8 {
	if len(match) == 1 {
		// Digit
		return match[0] - '0'
	}

	for i, digitName := range digitNames {
		if digitName == match {
			return uint8(i)
		}
	}

	// Invalid match
	return 255
}

func reverse(in string) string {
	var sb strings.Builder
	runes := []rune(in)
	for i := len(runes) - 1; 0 <= i; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}

var digitNamesReg = `one|two|three|four|five|six|seven|eight|nine`
var r = regexp.MustCompile(`[1-9]|` + digitNamesReg)
var rInv = regexp.MustCompile(`[1-9]|` + reverse(digitNamesReg))

func Part2(inputPath string) any {
	accum := 0

	fileline.ForEach(inputPath, errexit.HandleScanError, func(line string) {
		accum += int(parseDigitMatch(r.FindString(line))) * 10
		accum += int(parseDigitMatch(reverse(rInv.FindString(reverse(line)))))
	})

	return accum
}
