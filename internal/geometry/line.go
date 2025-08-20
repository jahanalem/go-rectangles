package geometry

import "fmt"

type Line struct {
	Point1 Point
	Point2 Point
}

func NewLine(p1, p2 Point) Line {
	return Line{Point1: p1, Point2: p2}
}

func (l Line) Equals(other Line) bool {
	return (l.Point1.Equals(other.Point1) && (l.Point2.Equals(other.Point2))) ||
		(l.Point1.Equals(other.Point2) && (l.Point2.Equals(other.Point1)))
}

func (l Line) HashCode() int {
	return l.Point1.HashCode() ^ l.Point2.HashCode()
}

func (l Line) String() string {
	return fmt.Sprintf("%s, %s", l.Point1, l.Point2)
}
