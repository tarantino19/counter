package main

import (
	"fmt"
	"os"

	// "strings"
	"bytes"
)

	func main (){
		data, _ := os.ReadFile("./words.txt")
		wordCount := CountWords(data)
		fmt.Println("wordCount:", wordCount)
	}

	//use go built in library for returning strings
	func CountWords (data []byte) int{
		words := bytes.Fields(data) //package more suited for bytes (since it came from os.ReadFile())
		return len(words)
	}