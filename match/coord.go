package match

import (
	"fmt"
	"math"
)

type Coord struct {
	x int32
	y int32
}

func NewCoord(x, y int32) *Coord {
	return &Coord{
		x: x,
		y: y,
	}
}

func (c *Coord) GetCoord() (int32, int32) {
	return c.x, c.y
}
func (c *Coord) String() string {
	return fmt.Sprintf("x: %d, y: %d", c.x, c.y)
}
func (c *Coord) Distance(target *Coord) float64 {
	return math.Sqrt(math.Pow(float64(target.x-c.x), 2) + math.Pow(float64(target.y-c.y), 2))
}
