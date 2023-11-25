package masker

import (
	"bufio"
	"log"
	"os"
)

// data provider (reading from a file)
type Producer interface {
	produce() (data []string, e error)
}

type Produce struct {
	Producer
	path string
}

func (p Produce) produce() ([]string, error) {
	file, err := os.Open(p.path)
	line_list := []string{}
	if err != nil {
		log.Println("error opening the file", err)
		return line_list, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line_list = append(line_list, GetMasks(line, URL))
	}
	if err := scanner.Err(); err != nil {
		log.Println("error reading file", err)
		return line_list, err
	}
	return line_list, nil
}
