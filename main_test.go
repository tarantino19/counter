package main

import (
	"testing"
)
func TestCountWords (t *testing.T){
	input := "one two three four five"
	wants := 6

	result := countWords([] byte(input))

	if result != wants {
		t.Fail() //Use Errorf instead of Fail
	}
}