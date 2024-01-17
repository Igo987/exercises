package config_test

import (
	"testing"

	"github.com/Igo87/project/config"
	"github.com/stretchr/testify/assert"
)

func TestNewCongig(t *testing.T) {
	cfg := config.NewConfig()
	assert.NotNil(t, cfg)
}

func TestGetPathToLogFile(t *testing.T) {
	cfg := config.NewConfig()
	cfg.PathToLogFile = "./src/links.txt"
	assert.Equal(t, "./src/links.txt", cfg.GetPathToLogFile())
}

func TestReadConfig(t *testing.T) {
	/* tabular testing */
	testCases := []struct {
		input    string
		expected error
	}{
		{
			input:    "./config.yaml",
			expected: nil,
		},
	}

	for _, tc := range testCases {
		_, err := config.ReadConfig(tc.input)
		assert.Equal(t, tc.expected, err)

	}

}
