package processor

import "rectangles-calculator/internal/geometry"

type Processor struct {
	Points []geometry.Point
}

func NewProcessor(points []geometry.Point) *Processor {
	return &Processor{Points: points}
}

func DeduplicateRectangles(rects []*geometry.Rectangle) []*geometry.Rectangle {
	seen := make(map[geometry.RectangleKey]bool)
	var unique []*geometry.Rectangle

	for _, rect := range rects {
		key := rect.ToKey()

		if !seen[key] {
			seen[key] = true
			unique = append(unique, rect)
		}
	}

	return unique
}
