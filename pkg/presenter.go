package masker

import (
	"log"
	"os"
)

type Service struct {
	prod Produce
	pres Present
}

func (s *Service) Run(path_read, path_write string) error {
	s.prod.path = path_read
	s.pres.path = path_write
	data, err := s.prod.produce()
	if err != nil {
		log.Println("error opening the file", err)
		return err
	}
	err = s.pres.presenter(data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Service constructor
func NewService(r Produce, w Present) *Service {
	svc := new(Service)
	return svc
}

// result handler (writing to a file)
type Presenter interface {
	present(s []string) (err error)
}

type Present struct {
	path string
}

// Present constructor
func NewPresent() *Present {
	prznt := new(Present)
	return prznt
}

func (p Present) presenter(s []string) error {
	f, err := os.Create(p.path)
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
