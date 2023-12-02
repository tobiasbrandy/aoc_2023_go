package day02

import (
	"github.com/tobiasbrandy/aoc_2023_go/internal/errexit"
	"github.com/tobiasbrandy/aoc_2023_go/internal/fileline"
	"github.com/tobiasbrandy/aoc_2023_go/internal/parse"
	"strings"
)

func parseGame(game string, callback func(name string, count int) bool) bool {
	// Parse
	// Game %d: %d %s, ...; ...
	game = game[strings.IndexRune(game, ':')+1:]

	plays := strings.Split(game, ";")
	for _, play := range plays {
		cubes := strings.Split(play, ",")
		for _, cube := range cubes {
			cubeSplit := strings.Split(strings.TrimSpace(cube), " ")
			if !callback(cubeSplit[1], parse.Int(cubeSplit[0])) {
				return false
			}
		}
	}
	return true
}

func Part1(inputPath string) any {
	ret := 0

	availableCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	isValidGame := func(name string, count int) bool {
		available := availableCubes[name]
		if count > available {
			return false
		}
		return true
	}

	fileline.ForEach(inputPath, errexit.HandleScanError, func(game string) {
		gameId := parse.Int(game[5:strings.IndexRune(game, ':')])

		validGame := parseGame(game, isValidGame)
		if validGame {
			ret += gameId
		}
	})

	return ret
}

func Part2(inputPath string) any {
	ret := 0

	var cubesUsed map[string]int
	fillCubesUsed := func(name string, count int) bool {
		used := cubesUsed[name]
		if count > used {
			cubesUsed[name] = count
		}
		return true
	}

	fileline.ForEach(inputPath, errexit.HandleScanError, func(game string) {
		cubesUsed = map[string]int{}
		parseGame(game, fillCubesUsed)

		power := 1
		for _, v := range cubesUsed {
			power *= v
		}

		ret += power
	})

	return ret
}
