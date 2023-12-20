package masker

import (
	"log"
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
	prznt := new(Present)
	return prznt
}

func (p Present) present(s []string) error {
	f, err := os.Create(p.Path)
	if err != nil {
		log.Println("error creating the file", err)
		return err
	}
	defer f.Close()
	for _, line := range s {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Println("error when writing to a file", err)
			return err
		}
	}
	return nil
}
