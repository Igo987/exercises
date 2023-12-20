package masker

import (
	"bufio"
	"log"
	"os"
)

// we read the links and send them to the channel
func getLinks(links []string) <-chan string {
	out := make(chan string)
	go func() {
		for _, n := range links {
			out <- n
		}
		close(out)
	}()
	return out
}

// data provider (reading from a file)
type Producer interface {
	produce() (data []string, e error)
}

type Produce struct {
	Path string
}

// Produce constructor
func NewProduce() *Produce {
	prdc := new(Produce)
	return prdc
}

func (p Produce) produce() ([]string, error) {
	file, err := os.Open(p.Path)
	line_list := []string{}
	result := make([]string, len(line_list))
	someLink := make(chan string)
	someFlag := make(chan bool)
	if err != nil {
		log.Println("error opening the file", err)
		return line_list, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line_list = append(line_list, line)
	}

	if err := scanner.Err(); err != nil {
		log.Println("error reading file", err)
		return line_list, err
	}

	res := getLinks(line_list)
	for i := range res {
		go func(s string) {
			someLink <- s
			someFlag <- true
		}(i)

		someRes := GetMasks(someLink, URL)
		go func() {
			<-someFlag
			result = append(result, <-someRes)
		}()
	}

	return result, nil
}
