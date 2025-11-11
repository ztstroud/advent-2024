package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

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
}

