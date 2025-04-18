package main

import (
	"fmt"
	"log" //std error
	"os"

	// "strings"
	"bytes"
)

	func main (){
		data, err := os.ReadFile("./words.txt")
		log.SetFlags(0) //no date and time and other shiznits

		if err != nil {
			// fmt.Fprintln(os.Stderr, "FAILED to read file:", err)
			// os.Exit(1)
			log.Fatalf("\033[31mFAILED to read file: %v\033[0m\n", err)
		}
		

		wordCount := CountWords(data)
		fmt.Printf("SUCCESS: \033[32mwordCount: %d\033[0m\n", wordCount)
	}

	//use go built in library for returning strings
	func CountWords (data []byte) int{
		words := bytes.Fields(data) //package more suited for bytes (since it came from os.ReadFile())
		return len(words)
	}