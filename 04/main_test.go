package main

import (
	"testing"
)

func TestReverseByBytes(t *testing.T) {
	input := "XMAS"
	expected := "SAMX"
	actual := revrseByBytes(input)

	if expected != actual {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestCountOccurances(t *testing.T) {
	input := "AXMASSMASAMXMASS"
	expected := 3
	actual := countOccurrences(input, "XMAS")

	if expected != actual {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestExtractDiagonal(t *testing.T) {
	input := []string {
		"......",
		"..X...",
		"...X..",
		"....X.",
		"......",
	}

	expected := "XXX"
	actual := extractDiagonal(input, 2, 1, 3)

	if expected != actual {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

