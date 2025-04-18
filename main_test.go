package main

import (
	"testing"
)
func TestCountWords (t *testing.T){
	//first test
	input := "one two three four five"
	wants := 5

	result := CountWords([] byte(input))

	if result != wants {
		t.Log("Expected", wants, "got", result)
		t.Fail() //Use Errorf instead of Fail
	}

	//second test
	input = ""
	wants = 0

	result = CountWords([] byte(input))

	if result != wants {
		//logf for interpolaton works too
		t.Log("Expected", wants, "got", result)
		t.Fail()
	}

	//third test  - will fail/wrong

	input = "  "
	wants = 0

	result = CountWords([] byte(input))

	if result != wants {
		//logf for interpolaton works too
		t.Log("Expected", wants, "got", result)
		t.Fail()
	}
	//     main_test.go:39: Expected 0 got 3


}

