package filereader

import (
	"encoding/json"
	"os"
	"rectangles-calculator/internal/geometry"
)

func ReadPointsFromFile(filePath string) ([]geometry.Point, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	
	defer jsonFile.Close()

	var points []geometry.Point
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&points)
	if err != nil {
		return nil, err
	}

	return points, nil
}
