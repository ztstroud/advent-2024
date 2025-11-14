package main

import (
	"slices"
	"strings"
)

func revrseByBytes(text string) string {
	bytes := []byte(text)
	slices.Reverse(bytes)
	return string(bytes)
}

func extractDiagonal(grid []string, x, y, length int) string {
	bytes := make([]byte, length)
	for i := range length {
		bytes[i] = grid[y + i][x + i]
	}

	return string(bytes)
}

func countGridOccurances(grid []string, query string) int {
	reverseQuery := revrseByBytes(query)

	sum := 0
	for _, row := range grid {
		sum += strings.Count(row, query)
		sum += strings.Count(row, reverseQuery)
	}

	return sum
}

func countDiagonalOccurances(grid []string, query string) int {
	reverseQuery := revrseByBytes(query)

	count := 0
	for y := range len(grid) - len(query) + 1 {
		for x := range len(grid[y]) - len(query) + 1 {
			extracted := extractDiagonal(grid, x, y, len(query))

			if query == extracted {
				count += 1
			}

			if reverseQuery == extracted {
				count += 1
			}
		}
	}

	return count
}

func countCrosswordOccurrences(grid []string, query string) int {
	rotated := rotate90(grid)
	return countGridOccurances(grid, query) +
		countGridOccurances(rotated, query) +
		countDiagonalOccurances(grid, query) +
		countDiagonalOccurances(rotated, query)
}

