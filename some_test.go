package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	svc "pr.ru/pkg"
)

func TestGetMasks(t *testing.T) {
	text := "http://hehee see"
	separator := "http://"
	expected := "http://***** see"
	// actual := svc.getMasks(text, separator)
	actual := svc.GetMasks(text,separator)

	assert.Equal(t, expected, actual)
}
