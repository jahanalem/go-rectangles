package processor

import (
	"fmt"
	"rectangles-calculator/internal/geometry"
	"sort"
	"sync"
)

// FindRectangles is the main entry point for the processing logic.
func FindRectangles(points []geometry.Point) []*geometry.Rectangle {

	// Step 1: Data Cleaning - Remove duplicate points to ensure correctness and improve performance.
	uniquePoints := deduplicatePoints(points)
	fmt.Printf("Processing %d unique points.\n", len(uniquePoints))

	// Step 2: Group points by their Y-coordinate. This is the first step to finding horizontal lines.
	pointsByY := groupPointsByY(uniquePoints)
	fmt.Printf("Found %d unique Y-levels.\n", len(pointsByY))

	// Step 3: Create horizontal lines from the grouped points, in parallel.
	lines := createLinesParallel(pointsByY)
	fmt.Printf("Created %d potential horizontal lines.\n", len(lines))

	// Step 4: Find all rectangles from the lines, using a direct translation of the C# parallel algorithm.
	rectangles := findRectanglesParallel(lines)

	return rectangles
}

// deduplicatePoints removes duplicate points from the initial dataset.
func deduplicatePoints(points []geometry.Point) []geometry.Point {
	seen := make(map[geometry.Point]bool)
	var unique []geometry.Point
	for _, p := range points {
		if !seen[p] {
			seen[p] = true
			unique = append(unique, p)
		}
	}
	return unique
}

// groupPointsByY groups points by their Y-coordinate and sorts each group by X-coordinate.
func groupPointsByY(points []geometry.Point) map[int][]geometry.Point {
	groups := make(map[int][]geometry.Point)
	for _, p := range points {
		groups[p.Y] = append(groups[p.Y], p)
	}

	// Sorting by X is crucial for creating lines consistently.
	for _, group := range groups {
		sort.Slice(group, func(i, j int) bool {
			return group[i].X < group[j].X
		})
	}
	return groups
}

// createLinesParallel generates all possible horizontal lines from the Y-grouped points.
func createLinesParallel(pointsByY map[int][]geometry.Point) []geometry.Line {
	linesChan := make(chan geometry.Line, 1000)
	var wg sync.WaitGroup

	for _, group := range pointsByY {
		if len(group) < 2 {
			continue
		}
		wg.Add(1)
		go func(g []geometry.Point) {
			defer wg.Done()
			for i := 0; i < len(g); i++ {
				for j := i + 1; j < len(g); j++ {
					linesChan <- geometry.NewLine(g[i], g[j])
				}
			}
		}(group)
	}

	// This goroutine waits for all workers to finish, then closes the channel.
	go func() {
		wg.Wait()
		close(linesChan)
	}()

	var allLines []geometry.Line
	for line := range linesChan {
		allLines = append(allLines, line)
	}
	return allLines
}

func findRectanglesParallel(lines []geometry.Line) []*geometry.Rectangle {
	rectsMap := sync.Map{} // Thread-safe map to store unique rectangles.
	var wg sync.WaitGroup

	// Group lines by their Y-coordinate to prepare for comparison.
	linesByY := make(map[int][]geometry.Line)
	for _, l := range lines {
		// Normalizing the line ensures Point1.X is always less than Point2.X
		// This simplifies the rectangle check later.
		p1, p2 := l.Point1, l.Point2
		if p1.X > p2.X {
			p1, p2 = p2, p1
		}
		normalizedLine := geometry.NewLine(p1, p2)
		linesByY[l.Point1.Y] = append(linesByY[l.Point1.Y], normalizedLine)
	}

	// Get a sorted list of Y-levels to iterate over.
	var yLevels []int
	for y := range linesByY {
		yLevels = append(yLevels, y)
	}
	sort.Ints(yLevels)

	// Iterate through each Y-level and compare it with all subsequent Y-levels.
	for i, y1 := range yLevels {
		wg.Add(1)
		// We pass 'i' and 'y1' as arguments to the goroutine to avoid the classic loop variable capture bug.
		// This makes the Go implementation correct without changing the algorithm's logic.
		go func(currentIndex int, baseY int) {
			defer wg.Done()
			baseLines := linesByY[baseY]

			// Compare the current Y-level's lines with every Y-level that comes after it.
			for j := currentIndex + 1; j < len(yLevels); j++ {
				compY := yLevels[j]
				compLines := linesByY[compY]

				// This is the brute-force comparison, identical to the C# implementation.
				for _, baseLine := range baseLines {
					for _, compLine := range compLines {
						// Since we normalized the lines earlier, we only need one simple check.
						if baseLine.Point1.X == compLine.Point1.X && baseLine.Point2.X == compLine.Point2.X {
							rect := geometry.NewRectangle(baseLine, compLine)
							rectsMap.Store(rect.ToKey(), rect)
						}
					}
				}
			}
		}(i, y1)
	}

	wg.Wait()

	// Collect the results from the sync.Map.
	var finalRects []*geometry.Rectangle
	rectsMap.Range(func(key, value interface{}) bool {
		finalRects = append(finalRects, value.(*geometry.Rectangle))
		return true
	})

	return finalRects
}
