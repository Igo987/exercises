package masker

import (
	"bufio"
	"log"
	"os"
	"sync"
)

var wg sync.WaitGroup
var chanWg sync.WaitGroup

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

func getResult(linksList []string) <-chan string {
	res := make(chan string, len(linksList))
	link := make(chan string)
	allLinks := getLinks(linksList)
	for i := range allLinks {
		wg.Add(2)
		i := i
		go func(u string) {
			link <- u
			defer wg.Done()
		}(i)
		go func(s chan string) {
			res <- GetMasks(s, URL)
			defer wg.Done()
		}(link)
	}
	wg.Wait()
	return res
}

// data provider (reading from a file)
type Producer interface {
	produce() (data []string, e error)
}

type Produce struct {
	path string
}

// Produce constructor
func NewProduce() *Produce {
	prdc := new(Produce)
	return prdc
}

func (p Produce) produce() ([]string, error) {
	file, err := os.Open(p.path)
	line_list := []string{}
	result := []string{}

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
	res := getResult(line_list)
	chanWg.Add(len(line_list))
	go func() {
		for item := range res {
			result = append(result, item)
			chanWg.Done()
		}
	}()
	chanWg.Wait()
	return result, nil
}
