package main_test

import (
	"testing"

	counter "github.com/tarantino19/counter"
)

// ANSI color codes
const (
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorReset = "\033[0m"
)

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{name: "Five Words", input: "one two three four five", wants: 5},
		{name: "Empty String", input: "", wants: 0},
		{name: "Leading and Trailing Spaces", input: "   ", wants: 0},
		{name: "New Lines", input: "one two three\nfour five", wants: 5},
		{name: "Multiple Spaces", input: "This is a sentence.  This is another", wants: 7},
		{name: "Prefix multiple spaces", input: "  Hello", wants: 1},
		{name: "Tab char in code", input: "Hello\tWord\n", wants: 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := counter.CountWords([]byte(tc.input))
			if result != tc.wants {
				t.Errorf("%s%s Test Failed: Expected %d, got %d%s", colorRed, tc.name, tc.wants, result, colorReset)
			}
		})
	}
}