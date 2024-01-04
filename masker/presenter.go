package masker

import (
	"bufio"
	"fmt"
	"os"
)

// result handler (writing to a file)
type Presenter interface {
	present(s []string) (err error)
}

type Present struct {
	Path string
}

// Present constructor
func NewPresent() *Present {
	return &Present{}
}

func (p Present) present(s []string) error {
	f, err := os.Create(p.Path)
	if err != nil {
		return fmt.Errorf("error creating the file: %w", err)
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, line := range s {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error when writing to a file: %w", err)
		}
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("error flushing the writer: %w", err)
	}
	return nil
}
