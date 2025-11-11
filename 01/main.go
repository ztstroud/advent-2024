package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func listDist(left []int, right []int) int {
	// I don't like that this sorts the input, this operation should not modify them
	// However, for the challenge I am going to leave it because it doesn't matter
	sort.Ints(left)
	sort.Ints(right)

	dist := 0
	for i := range(left) {
		if left[i] <= right[i] {
			dist += right[i] - left[i]
		} else {
			dist += left[i] - right[i]
		}
	}

	return dist
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You must provide an input file")
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed to read from file: ", path, err)
	}

	left := make([]int, 0, 32)
	right := make([]int, 0, 32)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())

		leftValue, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatal("Failed to parse int from: ", values[0], err)
		}

		left = append(left, leftValue)

		rightValue, err := strconv.Atoi(values[1])
		if err != nil {
			log.Fatal("Failed to parse int from: ", values[1], err)
		}

		right = append(right, rightValue)
	}

	dist := listDist(left, right)
	fmt.Printf("%d\n", dist)
}

