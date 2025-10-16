package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"rectangles-calculator/internal/filereader"
	"rectangles-calculator/internal/processor"
	"strings"
	"time"
)

const dataFilePath = "data_points_16.json"

func main() {
	fmt.Println("Starting rectangle calculation...")
	points, err := filereader.ReadPointsFromFile(dataFilePath)
	if err != nil {
		log.Fatalf("Error reading data file: %v", err)
	}

	start := time.Now()
	rectangles := processor.FindRectangles(points)
	duration := time.Since(start)

	fmt.Println("----------------------------------------")
	fmt.Printf("Calculation finished in %s\n", duration)
	fmt.Printf("Found %d unique rectangles.\n", len(rectangles))
	fmt.Println("----------------------------------------")

	if len(rectangles) > 0 {
		fmt.Print("Do you want to print the found rectangles to the console? (y/n): ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := strings.ToLower(strings.TrimSpace(input))

		if choice == "y" {
			fmt.Println("Rectangles found:")
			for i, rect := range rectangles {
				points := rect.GetOrderedPoints()
				fmt.Printf(
					"  %d: [(%d, %d), (%d, %d), (%d, %d), (%d, %d)]\n",
					i+1,
					points[0].X, points[0].Y,
					points[1].X, points[1].Y,
					points[2].X, points[2].Y,
					points[3].X, points[3].Y,
				)
			}
		}
	}
}
