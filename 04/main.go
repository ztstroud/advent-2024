package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func xmasAt(grid []string, x, y int) bool {
	if grid[y + 1][x + 1] != 'A' {
		return false
	}

	c00 := grid[y][x]
	c01 := grid[y][x + 2]
	c10 := grid[y + 2][x]
	c11 := grid[y + 2][x + 2]

	return c00 == 'M' && c11 == 'S' && c01 == 'M' && c10 == 'S' ||
		c00 == 'S' && c11 == 'M' && c01 == 'M' && c10 == 'S' ||
		c00 == 'M' && c11 == 'S' && c01 == 'S' && c10 == 'M' ||
		c00 == 'S' && c11 == 'M' && c01 == 'S' && c10 == 'M'
}

func countXmas(grid []string) int {
	count := 0
	for y := range len(grid) - 2 {
		for x := range len(grid[y]) - 2 {
			if xmasAt(grid, x, y) {
				count += 1
			}
		}
	}

	return count
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("You must provide an input file\n")
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to read file: %s\n%v\n", path, err)
	}
	defer file.Close()

	grid := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	query := "XMAS"
	queryCount := countCrosswordOccurrences(grid, query)
	xmasCount := countXmas(grid)

	fmt.Printf("%s occurs %d times\n", query, queryCount)
	fmt.Printf("%d X-MAS\n", xmasCount)
}

