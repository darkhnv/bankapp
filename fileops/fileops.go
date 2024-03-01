package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// GetFloatFromFile reads a float64 value from a file
func GetFloatFromFile(fileName string) (float64, error) {
	// Read file contents
	data, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) { // Check if file does not exist
			return 0, nil
		}
		return 0, errors.New("failed to read file")
	}

	// Parse file contents to float64
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 0, errors.New("failed to parse stored value")
	}

	return value, nil
}

// WriteFloatToFile writes a float64 value to a file
func WriteFloatToFile(fileName string, value float64) error {
	// Convert float64 to string
	valueText := fmt.Sprintf("%.2f", value)

	// Write string to file
	err := os.WriteFile(fileName, []byte(valueText), 0644)
	if err != nil {
		return errors.New("failed to write to file")
	}
	return nil
}
