package helpers

import (
	"errors"
	"fmt"
	"os"
)

func GetInput(file_path string) (*os.File, error) {
	input, err := os.Open(file_path)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while opening file at %s", file_path))
	}

	return input, nil
}
