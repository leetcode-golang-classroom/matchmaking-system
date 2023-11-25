package match

import "fmt"

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
