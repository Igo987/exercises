package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMasks(t *testing.T) {
	text := "http://hehee see"
	separator := "http://"
	expected := "http://***** see"
	actual := getMasks(text, separator)

	assert.Equal(t, expected, actual)
}
