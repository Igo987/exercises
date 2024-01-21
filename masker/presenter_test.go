package masker

import "testing"

func TestPresent(t *testing.T) {
	p := Present{Path: "../src/test-file.txt"}

	// Testing for successful file creation and writing
	err := p.present([]string{"http://example1.com", "http://example1.com", "http://example1.com"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Testing for error when creating the file
	p.Path = "" // empty path should trigger an error
	err = p.present([]string{"http://example1.com", "http://example1.com", "http://example1.com"})
	if err == nil {
		t.Error("Expected error when creating the file, but got nil")
	}

	// Testing for error when writing to the file
	p.Path = "../src/test-file.txt"
	err = p.present([]string{"http://example1.com", "http://example1.com", "http://example1.com"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
