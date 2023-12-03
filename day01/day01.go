package day01

import (
	"github.com/tobiasbrandy/aoc_2023_go/internal/errexit"
	"github.com/tobiasbrandy/aoc_2023_go/internal/fileline"
	"github.com/tobiasbrandy/aoc_2023_go/internal/stringsext"
	"regexp"
)

func Part1(inputPath string) any {
	accum := 0

	fileline.ForEach(inputPath, errexit.HandleScanError, func(line string) {
		lineLen := len(line)

		for _, digit := range line {
			if stringsext.IsAsciiDigit(digit) {
				accum += int(digit-'0') * 10
				break
			}
		}

		for i := lineLen - 1; i >= 0; i-- {
			digit := line[i]
			if stringsext.IsAsciiDigit(rune(digit)) {
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

var digitNamesReg = `one|two|three|four|five|six|seven|eight|nine`
var r = regexp.MustCompile(`[1-9]|` + digitNamesReg)
var rInv = regexp.MustCompile(`[1-9]|` + stringsext.Reverse(digitNamesReg))

func Part2(inputPath string) any {
	accum := 0

	fileline.ForEach(inputPath, errexit.HandleScanError, func(line string) {
		accum += int(parseDigitMatch(r.FindString(line))) * 10
		accum += int(parseDigitMatch(stringsext.Reverse(rInv.FindString(stringsext.Reverse(line)))))
	})

	return accum
}
