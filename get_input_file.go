// Package helpers provide helper method for the
// Advent of Code challenge.
package helpers

import (
	"errors"
	"fmt"
	"os"
)

// GetInput takes a path to a file and returns a pointer to that file.
func GetInput(filePath string) (*os.File, error) {
	input, err := os.Open(filePath)

	if err != nil {
		errorMsg := fmt.Sprintf("Error while opening file at %s", filePath)
		return nil, errors.New(errorMsg)
	}

	return input, nil
}

func GetInputAsString(filePath string) (string, error) {
	buf, err := os.ReadFile(filePath)

	if err != nil {
		errorMsg := fmt.Sprintf("Error while opening file at %s", filePath)
		return "", errors.New(errorMsg)
	}

	return string(buf), nil
}
