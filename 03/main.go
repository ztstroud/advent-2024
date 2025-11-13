package main

import (
	"log"
	"regexp"
	"strconv"
)

/*
Return the sum of mul statements in the given string, ignoring invalid
characters.

mul statements must be of the format `mul(n,m)` where n and m are non-negative
integers.
*/
func sumMuls(memory string) int {
	mulRegEx := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := mulRegEx.FindAllStringSubmatch(memory, -1)

	sum := 0
	for _, match := range(matches) {
		left, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatalf("Failed to parse %s to int\n", match[1])
		}

		right, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatalf("Failed to parse %s to int\n", match[2])
		}

		sum += left * right
	}

	return sum
}

