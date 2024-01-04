package masker

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var wait sync.WaitGroup

const separator = "http://"

func TestGetMasks(t *testing.T) {
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
		result := GetMasks(tc.input, separator)
		assert.Equal(t, tc.expected, <-result)
	}
	go func() {
		wait.Wait()
	}()
}
