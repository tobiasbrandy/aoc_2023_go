package main

import (
	"flag"
	"fmt"
	"github.com/tobiasbrandy/aoc_2023_go/day01"
	"github.com/tobiasbrandy/aoc_2023_go/internal/errexit"
	"time"
)

type AoCSolver func(string, int) any

func PartsSolver(part1, part2 func(string) any) AoCSolver {
	return func(inputPath string, part int) any {
		switch part {
		case 1:
			return part1(inputPath)
		case 2:
			return part2(inputPath)
		default:
			panic("unreachable")
		}
	}
}

var DaySolvers = [...]AoCSolver{
	PartsSolver(day01.Part1, day01.Part2),
}

func main() {
	day := flag.Int("day", 0, "AoC challenge day number.")
	part := flag.Int("part", 1, "AoC challenge part number. Default: 1.")
	inputPath := flag.String("input", "", "Path to the input file. Default: `day{day}/input.txt`.")
	takeTime := flag.Bool("time", false, "Print execution time")
	test := flag.Bool("test", false, "Ignore `input` parameter and use `day{day}/test.txt` as input.")
	flag.Parse()

	if *day < 1 || *day > len(DaySolvers) {
		errexit.HandleArgsError(fmt.Errorf("day must be between 1 and %d: not %d", len(DaySolvers), *day))
	}

	if *part != 1 && *part != 2 {
		errexit.HandleArgsError(fmt.Errorf("AoC challenges only have part 1 or 2, not part %d", *part))
	}

	if *test {
		*inputPath = fmt.Sprintf("day%02d/test.txt", *day)
	} else if *inputPath == "" {
		*inputPath = fmt.Sprintf("day%02d/input.txt", *day)
	}

	t := time.Now()
	ret := DaySolvers[*day-1](*inputPath, *part)
	execTime := time.Since(t)

	fmt.Println(ret)
	if *takeTime {
		fmt.Println("Execution time:", execTime)
	}
}
