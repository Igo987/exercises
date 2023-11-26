package masker

import (
	"log"
	"os"
)

type Service struct {
	prod Produce
	pres presenty
}

func Run(path_read, path_write string) error {
	svc := NewService(path_read, path_write)
	data, err := svc.prod.produce()
	if err != nil {
		log.Println("error opening the file", err)
		return err
	}
	err = svc.pres.presenter(data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Service constructor
func NewService(path_read, path_write string) *Service {
	svc := new(Service)
	svc.prod.path = path_read
	svc.pres.path = path_write
	return svc
}

// result handler (writing to a file)
type Present interface {
	presenter(s []string) (err error)
}

type presenty struct {
	Present
	path string
}

func (p presenty) presenter(s []string) error {
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
