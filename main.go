package main

import (
	"fmt"
	"os"
)

func main (){
	data, _ := os.ReadFile("./words.txt")

	wordCount := CountWords(data)

	fmt.Println("wordCount:", wordCount)
}

func CountWords (data []byte) int{

	if len(data) == 0 {
		return 0
	} //deal with edge case first as guard

	wordCount := 0

	for _ , x := range data {
		if x == ' ' {    //if there's a space, add one word count //can also use const = 32 (space in ascii)
			wordCount++
		}
	}

	wordCount++

	return wordCount

}