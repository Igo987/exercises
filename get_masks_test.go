package main

import (
	"sync"
	"testing"

	masker "github.com/Igo87/project/masker"
)

var wait sync.WaitGroup

func TestGetMasks(t *testing.T) {
	separator := "http://"
	/* tabular testing */
	testCases := []struct {
		input    chan string
		value    string
		expected string
	}{
		{
			input:    make(chan string),
			value:    "there is not a single link here",
			expected: "there is not a single link here",
		},
		{
			input:    make(chan string),
			value:    "And here is the link to http://stackoverflow.com",
			expected: "And here is the link to http://*****************",
		},
		{
			input:    make(chan string),
			value:    "http://google.com",
			expected: "http://**********",
		},
	}

	for _, tc := range testCases {
		tc := tc
		wait.Add(3)
		go func() {
			tc.input <- tc.value
			close(tc.input)
			defer wait.Done()
		}()
		result := masker.GetMasks(tc.input, separator)
		if <-result != tc.expected {
			t.Errorf("MyFunction(%v) = %d; want %s", tc.input, result, tc.expected)
		}
	}
	go func() {
		wait.Wait()
	}()
}
