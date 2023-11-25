package main

import (
	"testing"

	masker "github.com/Igo87/project/pkg"
	"github.com/stretchr/testify/assert"
)

func TestGetMasks(t *testing.T) {
	text := "http://hehee see"
	separator := "http://"
	expected := "http://***** see"
	actual := masker.GetMasks(text, separator)

	assert.Equal(t, expected, actual)

	/* tabular testing */
	testCases := []struct {
		input    string
		expected string
	}{
		{"http://google.com", "http://**********"},
		{"there is not a single link here", "there is not a single link here"},
		{"And here is the link to http://stackoverflow.com", "And here is the link to http://*****************"},
	}

	for _, tc := range testCases {
		result := masker.GetMasks(tc.input, separator)
		if result != tc.expected {
			t.Errorf("MyFunction(%s) = %s; want %s", tc.input, result, tc.expected)
		}
	}
}
