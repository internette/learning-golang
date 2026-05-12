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

func buildWordDictionary(paragraph string) map[string]int {
	words := strings.Split(paragraph, " ")
	dict := make(map[string]int)
	for _, w := range words {
		dict[w]++
	}
	return dict
}

func findMostOccurringWord(paragraph string) (string, int) {
	dict := buildWordDictionary(paragraph)
	mostOccurringWord := ""
	maxCount := 0
	for word, count := range dict {
		if count > maxCount {
			maxCount = count
			mostOccurringWord = word
		}
	}
	return mostOccurringWord, maxCount
}

func main() {
	thingToFind := flag.String("find", "word", "Thing to find")
	wordToFind := flag.String("word", "", "Word to find")
	paragraph := flag.String("paragraph", "", "Paragraph to search")
	flag.Parse()
	switch *thingToFind {
	case "word":
		wordCount := countWord(*wordToFind, *paragraph)
		fmt.Println("Frequency of '", *wordToFind, "':", wordCount)
	case "most-occurring":
		mostOccurringWord, maxCount := findMostOccurringWord(*paragraph)
		fmt.Println("Most occurring word:", mostOccurringWord, ", count:", maxCount)
	default:
		wordCount := countWord(*wordToFind, *paragraph)
		fmt.Println("Frequency of '", *wordToFind, "':", wordCount)
	}
}
