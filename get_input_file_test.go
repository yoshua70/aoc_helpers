package helpers

import (
	"os"
	"testing"
)

func TestOpeningExistingFile(t *testing.T) {
	// Arrange: create a new file named input.txt
	// 0644: File owner can read and write (6), users in the same group
	// and other can read
	filePath := "/tmp/input.txt"
	fileData := []byte("Hello, World!\n")
	err := os.WriteFile(filePath, fileData, 0644)

	if err != nil {
		panic("Error creating file.")
	}

	// Act: try to access the previously created file
	input, err := GetInput(filePath)

	// Assert: we should have a pointer to that file
	if input == nil || err != nil {
		t.Fatalf(`Cannot get specified file: %q`, filePath)
	}

	// Cleanup: delete the created file
	os.Remove(filePath)
}

func TestOpeningNonExistingFile(t *testing.T) {
	// Arrange
	filePath := "/tmp/input.txt"

	// Act: try to access the previously created file
	input, err := GetInput(filePath)

	// Assert: we should have a pointer to that file
	if input != nil || err == nil {
		t.Fatalf(`Should not be able to get specified file: %q`, filePath)
	}

	// Cleanup: delete the created file
	os.Remove(filePath)
}

func TestGetInputAsString(t *testing.T) {
	// Arrange: create a new file named input.txt
	// 0644: File owner can read and write (6), users in the same group
	// and other can read
	filePath := "/tmp/input.txt"
	dfileContent := "Hello, World!\n"
	fileData := []byte(dfileContent)
	err := os.WriteFile(filePath, fileData, 0644)

	if err != nil {
		panic("Error creating file.")
	}

	// Act: try to access the previously created file
	input, err := GetInputAsString(filePath)

	// Assert: we should have a pointer to that file
	if input == "" || input != dfileContent || err != nil {
		t.Fatalf(`Cannot get specified file: %q`, filePath)
	}

	// Cleanup: delete the created file
	os.Remove(filePath)
}
