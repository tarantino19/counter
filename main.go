package main

import (
	"fmt"
	"os"
)

func main (){
	data, _ := os.ReadFile("./words.txt")

	wordCount := countWords(data)

	fmt.Println("wordCount:", wordCount)
}

func countWords (data []byte) int{
	wordCount := 0

	for _ , x := range data {
		if x == ' ' {    //if there's a space, add one word count //can also use const = 32 (space in ascii)
			wordCount++
		}
	}

	wordCount++

	return wordCount

}