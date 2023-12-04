package main

import (
	"fmt"
	"strings"
)

func wordFrequency(sentence string) map[string]int {
	// Split the sentence into words
	words := strings.Fields(sentence)

	// Create a map to store word frequencies
	frequencies := make(map[string]int)

	// Count the occurrences of each word
	for i := 0; i < (len(words)); i++ {
		frequencies[words[i]]++
	}

	return frequencies
}

func main() {
	text := "This is a sample sentence. This sentence is for word frequency count in Golang."
	frequencyMap := wordFrequency(text)

	// Print the word frequencies
	fmt.Println("Word frequencies:")
	for word, count := range frequencyMap {
		fmt.Printf("%s: %d\n", word, count)
	}

	fmt.Println(strings.Fields("Aksndfgbvh"))
}
