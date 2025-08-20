package geometry

import "fmt"

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Point) HashCode() int {
	return p.X ^ p.Y
}

func (p Point) String() string {
	return fmt.Sprintf("%d, %d", p.X, p.Y)
}
