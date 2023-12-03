package day03

import (
	"github.com/tobiasbrandy/aoc_2023_go/internal/errexit"
	"github.com/tobiasbrandy/aoc_2023_go/internal/fileline"
	"github.com/tobiasbrandy/aoc_2023_go/internal/parse"
	"github.com/tobiasbrandy/aoc_2023_go/internal/pos"
	"github.com/tobiasbrandy/aoc_2023_go/internal/set"
	"github.com/tobiasbrandy/aoc_2023_go/internal/stringsext"
)

type schematicT [][]rune

func (s schematicT) Rows() int {
	return len(s)
}

func (s schematicT) Cols() int {
	return len(s[0])
}

func (s schematicT) InBounds(x, y int) bool {
	return y >= 0 && y < s.Rows() && x >= 0 && x < s.Cols()
}

func (s schematicT) IsSymbol(x, y int) bool {
	elem := s[y][x]
	return elem != '.' && !stringsext.IsAsciiDigit(elem)
}

func (s schematicT) IsGear(x, y int) bool {
	return s[y][x] == '*'
}

func (s schematicT) IsPartNumber(x, y int) bool {
	return stringsext.IsAsciiDigit(s[y][x])
}

func (s schematicT) PartNumberStart(x, y int) (int, int) {
	var startX int
	for startX = x - 1; startX >= 0; startX-- {
		if !s.IsPartNumber(startX, y) {
			break
		}
	}
	return startX + 1, y
}

func (s schematicT) PartNumber(startX, startY int) int {
	cols := s.Cols()

	var endX int
	for endX = startX + 1; endX < cols; endX++ {
		if !s.IsPartNumber(endX, startY) {
			break
		}
	}

	return parse.Int(string(s[startY][startX:endX]))
}

func buildSchematic(filePath string) schematicT {
	var schematic schematicT

	fileline.ForEach(filePath, errexit.HandleScanError, func(line string) {
		schematic = append(schematic, []rune(line))
	})

	return schematic
}

func Part1(inputPath string) any {
	ret := 0

	schematic := buildSchematic(inputPath)
	rows := schematic.Rows()
	cols := schematic.Cols()

	// Get all part numbers
	partNumbers := set.Set[pos.D2]{}
	// For each element
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Find symbol
			if schematic.IsSymbol(col, row) {
				// Get symbol neighbours
				for _, n := range pos.Neighbours8(col, row) {
					// Find neighbouring part number
					if schematic.InBounds(n.X, n.Y) && schematic.IsPartNumber(n.X, n.Y) {
						// Add number start position
						partNumbers.Add(pos.New2D(schematic.PartNumberStart(n.X, n.Y)))
					}
				}
			}
		}
	}

	for partNumberIdx := range partNumbers {
		ret += schematic.PartNumber(partNumberIdx.X, partNumberIdx.Y)
	}

	return ret
}

func Part2(inputPath string) any {
	ret := 0

	schematic := buildSchematic(inputPath)
	rows := schematic.Rows()
	cols := schematic.Cols()

	// Reusable set
	partNumbers := set.Set[pos.D2]{}

	// For each element
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Find symbol
			if schematic.IsGear(col, row) {
				// Get symbol neighbours
				for _, n := range pos.Neighbours8(col, row) {
					// Find neighbouring part number
					if schematic.InBounds(n.X, n.Y) && schematic.IsPartNumber(n.X, n.Y) {
						// Add number start position
						partNumbers.Add(pos.New2D(schematic.PartNumberStart(n.X, n.Y)))
					}
				}

				// Add gear number if exactly touches 2 numbers
				if partNumbers.Len() == 2 {
					gearValue := 1
					for partNumberIdx := range partNumbers {
						gearValue *= schematic.PartNumber(partNumberIdx.X, partNumberIdx.Y)
					}
					ret += gearValue
				}

				// Reuse set
				partNumbers.Clear()
			}
		}
	}

	return ret
}
