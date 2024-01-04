package masker

import (
	"bufio"
	"log"
	"os"
	"sync"
)

// we read the links and send them to the channel
func getLinks(links []string) <-chan string {
	out := make(chan string, len(links))
	go func() {
		defer close(out)
		for _, n := range links {
			out <- n
		}
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
	return &Produce{}
}

func (p Produce) produce() ([]string, error) {
	file, err := os.Open(p.Path)
	if err != nil {
		log.Println("error opening the file", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineList := make([]string, 0, 100)
	for scanner.Scan() {
		lineList = append(lineList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println("error reading file", err)
		return nil, err
	}

	result := make([]string, 0, len(lineList))

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, line := range lineList {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			someCh := getLinks([]string{l})
			someRes := GetMasks(someCh, URL)
			mu.Lock()
			defer mu.Unlock()
			result = append(result, <-someRes)
		}(line)
	}

	wg.Wait()

	return result, nil
}
