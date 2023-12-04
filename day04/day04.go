package day04

import (
	"github.com/gammazero/deque"
	"github.com/tobiasbrandy/aoc_2023_go/internal/errexit"
	"github.com/tobiasbrandy/aoc_2023_go/internal/fileline"
	"github.com/tobiasbrandy/aoc_2023_go/internal/parse"
	"slices"
	"strings"
)

func parseCard(card string) (id int, winnings []int, currents []int) {
	headerIdx := strings.IndexRune(card, ':')
	id = parse.Int(strings.TrimSpace(card[4:headerIdx])) // Card %d:

	splitIdx := strings.IndexRune(card[headerIdx+1:], '|') + headerIdx + 1

	for _, s := range strings.Split(card[headerIdx+1:splitIdx], " ") {
		if s != "" {
			winnings = append(winnings, parse.Int(s))
		}
	}

	for _, s := range strings.Split(card[splitIdx+1:], " ") {
		if s != "" {
			currents = append(currents, parse.Int(s))
		}
	}

	return id, winnings, currents
}

func Part1(inputPath string) any {
	ret := 0

	fileline.ForEach(inputPath, errexit.HandleScanError, func(card string) {
		_, winnings, currents := parseCard(card)

		cardVal := 1
		for _, current := range currents {
			if slices.Contains(winnings, current) {
				cardVal *= 2
			}
		}
		cardVal /= 2

		ret += cardVal
	})

	return ret
}

func Part2(inputPath string) any {
	ret := 0

	cardCounts := deque.Deque[int]{} // Circular buffer

	fileline.ForEach(inputPath, errexit.HandleScanError, func(card string) {
		cardCount := 1
		if cardCounts.Len() > 0 {
			cardCount = cardCounts.PopFront()
		}

		_, winnings, currents := parseCard(card)

		cardVal := 0
		for _, current := range currents {
			if slices.Contains(winnings, current) {
				cardVal++
			}
		}

		insertCount := cardVal - min(cardVal, cardCounts.Len())
		for i := 0; i < insertCount; i++ {
			cardCounts.PushBack(1)
		}

		for i := 0; i < cardVal; i++ {
			cardCounts.Set(i, cardCounts.At(i)+cardCount)
		}

		ret += cardCount
	})

	return ret
}
