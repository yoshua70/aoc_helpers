# Advent of Code Helpers

A collection of helper function I used to solve challenges from Advent of Code (AoC).

## Helpers

### GetInput

```go
func GetInput(filePath string) (*os.File, error)
```

GetInput takes a path to a file and returns a pointer to that file.

### GetInputAsString

```go
func GetInputAsString(filePath string) (string, error)
```

GetInputAsString takes a path to a file and returns a string containing the content of the file.
