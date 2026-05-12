package main

import (
	"flag"
	"fmt"
	"strings"
	"unicode"
)

func countWord(word string, paragraph string) int {
	count := 0
	if !strings.Contains(paragraph, word) {
		fmt.Println("The paragraph does not contain the word:", word)
	}
	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSpace(r)
	})
	for _, w := range words {
		if strings.EqualFold(w, word) {
			count++
		}
	}
	return count
}

func buildWordDictionary(paragraph string) map[string]int {
	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSpace(r)
	})
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

func findLeastOccurringWord(paragraph string) (string, int) {
	dict := buildWordDictionary(paragraph)
	leastOccurringWord := ""
	minCount := int(^uint(0) >> 1) // max int value as starting point

	for word, count := range dict {
		if count < minCount {
			minCount = count
			leastOccurringWord = word
		}
	}

	return leastOccurringWord, minCount
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
	case "least-occurring":
		leastOccurringWord, minCount := findLeastOccurringWord(*paragraph)
		fmt.Println("Least occurring word:", leastOccurringWord, ", count:", minCount)
	default:
		wordCount := countWord(*wordToFind, *paragraph)
		fmt.Println("Frequency of '", *wordToFind, "':", wordCount)
	}
}
