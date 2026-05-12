package main

import (
	"slices"
	"testing"
)

const testParagraph = "The cat sat near the old tree, and the cat watched the birds near the old branches of the old tree. Every bird that flew near the tree made the cat crouch lower and lower, and every time a bird landed near the old roots, the cat would inch closer. The birds, unaware that the cat was near, continued to hop from branch to branch of the old tree, singing the same songs the birds had always sung near trees like this old, familiar tree where the cat had always waited."

func TestCountWord(t *testing.T) {
	testWord := "cat"
	expectedCount := 6
	actualCount := countWord(testWord, testParagraph)
	if actualCount != expectedCount {
		t.Errorf("Expected count for '%s': %d, but got %d", testWord, expectedCount, actualCount)
	}
}

func TestBuildWordDictionary(t *testing.T) {
	expectedDict := map[string]int{
		"the":       16,
		"cat":       6,
		"old":       6,
		"near":      6,
		"tree":      5,
		"and":       3,
		"birds":     3,
		"bird":      2,
		"lower":     2,
		"every":     2,
		"that":      2,
		"branch":    2,
		"always":    2,
		"had":       2,
		"to":        2,
		"of":        2,
		"sat":       1,
		"watched":   1,
		"branches":  1,
		"flew":      1,
		"made":      1,
		"crouch":    1,
		"time":      1,
		"a":         1,
		"landed":    1,
		"roots":     1,
		"would":     1,
		"inch":      1,
		"closer":    1,
		"unaware":   1,
		"was":       1,
		"continued": 1,
		"hop":       1,
		"from":      1,
		"singing":   1,
		"same":      1,
		"songs":     1,
		"sung":      1,
		"trees":     1,
		"like":      1,
		"this":      1,
		"familiar":  1,
		"where":     1,
		"waited":    1,
	}
	actualDict := buildWordDictionary(testParagraph)
	for word, expectedCount := range expectedDict {
		if actualCount := actualDict[word]; actualCount != expectedCount {
			t.Errorf("Expected count for '%s': %d, but got %d", word, expectedCount, actualCount)
		}
	}
}

func TestFindMostOccurringWord(t *testing.T) {
	expectedWord := "the"
	expectedCount := 16
	actualWord, actualCount := findMostOccurringWord(testParagraph)
	if actualWord != expectedWord || actualCount != expectedCount {
		t.Errorf("Expected most occurring word: '%s' with count %d, but got '%s' with count %d", expectedWord, expectedCount, actualWord, actualCount)
	}
}

func TestFindLeastOccurringWord(t *testing.T) {
	potentialExpectedWords := []string{"sat", "watched", "branches", "flew", "made", "crouch", "time", "a", "landed", "roots", "would", "inch", "closer", "unaware", "was", "continued", "hop", "from", "singing", "same", "songs", "sung", "trees", "like", "this", "familiar", "where", "waited"}
	expectedCount := 1
	actualWord, actualCount := findLeastOccurringWord(testParagraph)
	if !slices.Contains(potentialExpectedWords, actualWord) || actualCount != expectedCount {
		t.Errorf("Expected least occurring word to be in array: %v, but got '%s' with count %d", potentialExpectedWords, actualWord, actualCount)
	}
}
