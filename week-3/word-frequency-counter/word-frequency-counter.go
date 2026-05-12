package main

import (
	"flag"
	"fmt"
	"strings"
)

func countWord(word string, paragraph string) int {
	count := 0
	if !strings.Contains(paragraph, word) {
		fmt.Println("The paragraph does not contain the word:", word)
	}
	words := strings.Split(paragraph, " ")
	for _, w := range words {
		if w == word {
			count++
		}
	}
	return count
}

func main() {
	wordToFind := flag.String("word", "", "Word to find")
	paragraph := flag.String("paragraph", "", "Paragraph to search")
	flag.Parse()
	wordCount := countWord(*wordToFind, *paragraph)
	fmt.Println("Frequency of '", *wordToFind, "':", wordCount)
}
