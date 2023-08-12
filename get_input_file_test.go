package helpers

import (
	"os"
	"testing"
)

func TestOpeningExistingFile(t *testing.T) {
	// Arrange: create a new file named input.txt
	// 0644: File owner can read and write (6), users in the same group
	// and other can read
	file_path := "/tmp/input.txt"
	file_data := []byte("Hello, World!\n")
	err := os.WriteFile(file_path, file_data, 0644)

	if err != nil {
		panic("Error creating file.")
	}

	// Act: try to access the previously created file
	input, err := GetInput(file_path)

	// Assert: we should have a pointer to that file
	if input == nil || err != nil {
		t.Fatalf(`Cannot get specified file: %q`, file_path)
	}

	// Cleanup: delete the created file
	os.Remove(file_path)
}

func TestOpeningNonExistingFile(t *testing.T) {
	// Arrange
	file_path := "/tmp/input.txt"

	// Act: try to access the previously created file
	input, err := GetInput(file_path)

	// Assert: we should have a pointer to that file
	if input != nil || err == nil {
		t.Fatalf(`Should not be able to get specified file: %q`, file_path)
	}

	// Cleanup: delete the created file
	os.Remove(file_path)
}
