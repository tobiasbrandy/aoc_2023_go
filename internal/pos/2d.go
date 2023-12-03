package pos

import (
	"fmt"
	"io"

	"github.com/tobiasbrandy/aoc_2023_go/internal/hashext"
	"github.com/tobiasbrandy/aoc_2023_go/internal/mathext"
)

func Neighbours4(x, y int) [4]D2 {
	return [...]D2{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	}
}

func Neighbours5(x, y int) [5]D2 {
	return [...]D2{
		{x - 1, y},
		{x + 1, y},
		{x, y},
		{x, y - 1},
		{x, y + 1},
	}
}

func Neighbours8(x, y int) [8]D2 {
	return [...]D2{
		{x - 1, y - 1},
		{x - 1, y},
		{x - 1, y + 1},

		{x, y - 1},
		// {x, y}, Me
		{x, y + 1},

		{x + 1, y - 1},
		{x + 1, y},
		{x + 1, y + 1},
	}
}

func Neighbours9(x, y int) [9]D2 {
	return [...]D2{
		{x - 1, y - 1},
		{x - 1, y},
		{x - 1, y + 1},

		{x, y - 1},
		{x, y},
		{x, y + 1},

		{x + 1, y - 1},
		{x + 1, y},
		{x + 1, y + 1},
	}
}

func Distance1_2D(x1, y1, x2, y2 int) int {
	return mathext.IntAbs(x1-x2) + mathext.IntAbs(y1-y2)
}

type D2 struct {
	X, Y int
}

func New2D(x, y int) D2 {
	return D2{x, y}
}

func (p D2) Neighbours4() [4]D2 {
	return Neighbours4(p.X, p.Y)
}

func (p D2) Neighbours5() [5]D2 {
	return Neighbours5(p.X, p.Y)
}

func (p D2) Neighbours8() [8]D2 {
	return Neighbours8(p.X, p.Y)
}

func (p D2) Neighbours9() [9]D2 {
	return Neighbours9(p.X, p.Y)
}

func (p D2) Distance1(o D2) int {
	return Distance1_2D(p.X, p.Y, o.X, o.Y)
}

func (p D2) String() string {
	return fmt.Sprint("(", p.X, ", ", p.Y, ")")
}

func (p D2) Hash(h io.Writer) {
	hashext.HashNum(h, p.X)
	hashext.HashNum(h, p.Y)
}
