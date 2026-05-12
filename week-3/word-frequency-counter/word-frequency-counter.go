package main

import (
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
	testParagraph := "The cat sat near the old tree, and the cat watched the birds near the old branches of the old tree. Every bird that flew near the tree made the cat crouch lower and lower, and every time a bird landed near the old roots, the cat would inch closer. The birds, unaware that the cat was near, continued to hop from branch to branch of the old tree, singing the same songs the birds had always sung near trees like this old, familiar tree where the cat had always waited."
	wordCount := countWord("cat", testParagraph)
	fmt.Println("Frequency of 'cat':", wordCount)
}
