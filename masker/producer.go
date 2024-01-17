package masker

import (
	"bufio"
	"context"
	"errors"
	"os"
	"sync"
	"time"

	"github.com/Igo87/project/logger"
)

var log = logger.LogStart("slog.LevelDebug")

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

var someCancel = make(chan struct{})

// Stop cancels the context if we get a signal
func Stop(ctx context.Context) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			someCancel <- struct{}{}
			log.Warn("the program was interrupted by the user")
			return
		case <-ticker.C:
			log.Info("the programm is still running")
		}
	}
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
		log.Error("error opening the file", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineList := make([]string, 0, 100)
	for scanner.Scan() {
		lineList = append(lineList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Error("error reading file", err)
		return nil, err
	}

	result := make([]string, 0, len(lineList))
	wait := &sync.WaitGroup{}
	var mu sync.Mutex
	someCh := getLinks(lineList)
	anotherCh := make(chan string)

	for {
		select {
		case l, ok := <-someCh:
			if !ok {
				return result, nil
			}
			wait.Add(2)
			go func() {
				anotherCh <- l
				defer wait.Done()
			}()
			go func() {
				defer wait.Done()
				defer mu.Unlock()
				someRes := GetMasks(anotherCh, URL)
				mu.Lock()
				result = append(result, <-someRes)
			}()
			wait.Wait()
		case <-someCancel:
			return result, errors.New("the program was interrupted by the user")
		}
	}
}
