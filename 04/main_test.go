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

func TestCountGridOccurrences(t *testing.T) {
	input := []string {
		".XMAS..",
		"SAMXMAS",
		".......",
	}

	expected := 3
	actual := countGridOccurances(input, "XMAS")

	if expected != actual {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestCountDiagonalOccurrences(t *testing.T) {
	input := []string {
		"..XX...",
		".X.MM..",
		"..M.AA.",
		"...A.SS",
		"....S..",
	}

	expected := 3
	actual := countDiagonalOccurances(input, "XMAS")

	if expected != actual {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

