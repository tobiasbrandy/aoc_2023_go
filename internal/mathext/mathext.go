package mathext

import (
	"github.com/tobiasbrandy/aoc_2023_go/internal/constraintsext"
)

func IntAbs[T constraintsext.Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Sign[T constraintsext.Number](x T) int {
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}

func IntPow[T constraintsext.Integer, U constraintsext.Unsigned](x T, exp U) T {
	var ret T = 1

	for {
		if exp&1 != 0 {
			ret *= x
		}
		exp >>= 1
		if exp == 0 {
			break
		}
		x *= x
	}

	return ret
}

// Mod is like %, but always returns a positive number or zero
func Mod[T constraintsext.Integer](x, m T) T {
	return (x%m + m) % m
}
