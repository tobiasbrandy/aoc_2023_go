package pos

import (
	"fmt"
	"github.com/tobiasbrandy/aoc_2023_go/internal/hashext"
	"github.com/tobiasbrandy/aoc_2023_go/internal/mathext"
	"io"
)

func Neighbours6(x, y, z int) [6]D3 {
	return [...]D3{
		{x - 1, y, z},
		{x + 1, y, z},
		{x, y - 1, z},
		{x, y + 1, z},
		{x, y, z + 1},
		{x, y, z - 1},
	}
}

func Neighbours7(x, y, z int) [7]D3 {
	return [...]D3{
		{x - 1, y, z},
		{x + 1, y, z},
		{x, y - 1, z},
		{x, y, z},
		{x, y + 1, z},
		{x, y, z + 1},
		{x, y, z - 1},
	}
}

func Neighbours26(x, y, z int) [26]D3 {
	return [...]D3{
		{x - 1, y - 1, z - 1},
		{x - 1, y - 1, z},
		{x - 1, y - 1, z + 1},
		{x - 1, y, z - 1},
		{x - 1, y, z},
		{x - 1, y, z + 1},
		{x - 1, y + 1, z - 1},
		{x - 1, y + 1, z},
		{x - 1, y + 1, z + 1},

		{x, y - 1, z - 1},
		{x, y - 1, z},
		{x, y - 1, z + 1},
		{x, y, z - 1},
		// {x, y, z}, Me
		{x, y, z + 1},
		{x, y + 1, z - 1},
		{x, y + 1, z},
		{x, y + 1, z + 1},

		{x + 1, y - 1, z - 1},
		{x + 1, y - 1, z},
		{x + 1, y - 1, z + 1},
		{x + 1, y, z - 1},
		{x + 1, y, z},
		{x + 1, y, z + 1},
		{x + 1, y + 1, z - 1},
		{x + 1, y + 1, z},
		{x + 1, y + 1, z + 1},
	}
}

func Neighbours27(x, y, z int) [27]D3 {
	return [...]D3{
		{x - 1, y - 1, z - 1},
		{x - 1, y - 1, z},
		{x - 1, y - 1, z + 1},
		{x - 1, y, z - 1},
		{x - 1, y, z},
		{x - 1, y, z + 1},
		{x - 1, y + 1, z - 1},
		{x - 1, y + 1, z},
		{x - 1, y + 1, z + 1},

		{x, y - 1, z - 1},
		{x, y - 1, z},
		{x, y - 1, z + 1},
		{x, y, z - 1},
		{x, y, z},
		{x, y, z + 1},
		{x, y + 1, z - 1},
		{x, y + 1, z},
		{x, y + 1, z + 1},

		{x + 1, y - 1, z - 1},
		{x + 1, y - 1, z},
		{x + 1, y - 1, z + 1},
		{x + 1, y, z - 1},
		{x + 1, y, z},
		{x + 1, y, z + 1},
		{x + 1, y + 1, z - 1},
		{x + 1, y + 1, z},
		{x + 1, y + 1, z + 1},
	}
}

func Distance1_3D(x1, y1, z1, x2, y2, z2 int) int {
	return mathext.IntAbs(x1-x2) + mathext.IntAbs(y1-y2) + mathext.IntAbs(z1-z2)
}

type D3 struct {
	X, Y, Z int
}

func New3D(x, y, z int) D3 {
	return D3{x, y, z}
}

func (p D3) Neighbours6() [6]D3 {
	return Neighbours6(p.X, p.Y, p.Z)
}

func (p D3) Neighbours7() [7]D3 {
	return Neighbours7(p.X, p.Y, p.Z)
}

func (p D3) Neighbours26() [26]D3 {
	return Neighbours26(p.X, p.Y, p.Z)
}

func (p D3) Neighbours27() [27]D3 {
	return Neighbours27(p.X, p.Y, p.Z)
}

func (p D3) Distance1(o D3) int {
	return Distance1_3D(p.X, p.Y, p.Z, o.X, o.Y, o.Z)
}

func (p D3) String() string {
	return fmt.Sprint("(", p.X, ", ", p.Y, ", ", p.Z, ")")
}

func (p D3) Hash(h io.Writer) {
	hashext.HashNum(h, p.X)
	hashext.HashNum(h, p.Y)
	hashext.HashNum(h, p.Z)
}
