package geometry

import (
	"fmt"
	"reflect"
	"sort"
)

const (
	hashSeed   = 19
	hashFactor = 31
)

type Rectangle struct {
	Line1 Line
	Line2 Line

	cachedOrderedPoints []Point
	cachedHashCode      *int
	cachedString        *string
}

func NewRectangle(l1, l2 Line) *Rectangle {
	return &Rectangle{Line1: l1, Line2: l2}
}

func (r *Rectangle) GetOrderedPoints() []Point {
	if r.cachedOrderedPoints != nil {
		return r.cachedOrderedPoints
	}

	points := []Point{
		r.Line1.Point1,
		r.Line1.Point2,
		r.Line2.Point1,
		r.Line2.Point2,
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].X != points[j].X {
			return points[i].X < points[j].X
		}
		return points[i].Y < points[j].Y
	})

	r.cachedOrderedPoints = points

	return r.cachedOrderedPoints
}

func (r *Rectangle) Equals(other *Rectangle) bool {
	if other == nil {
		return false
	}

	thisPoint := r.GetOrderedPoints()
	otherPoint := other.GetOrderedPoints()

	return reflect.DeepEqual(thisPoint, otherPoint)
}

func (r *Rectangle) HashCode() int {
	if r.cachedHashCode != nil {
		return *r.cachedHashCode
	}

	hash := hashSeed
	for _, point := range r.GetOrderedPoints() {
		hash = hash*hashFactor + point.HashCode()
	}

	r.cachedHashCode = &hash

	return hash
}

func (r *Rectangle) String() string {
	if r.cachedString != nil {
		return *r.cachedString
	}

	str := fmt.Sprintf("[%s,%s]", r.Line1, r.Line2)

	r.cachedString = &str

	return str
}
