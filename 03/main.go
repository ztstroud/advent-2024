package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func sumMulsConditional(memory string) int {
	sum := 0

	const on = "do()"
	const off = "don't()"

	for len(memory) > 0 {
		offIndex := strings.Index(memory, off)

		if offIndex == -1 {
			sum += sumMuls(memory)
			break
		}

		sum += sumMuls(memory[:offIndex])

		afterOffIndex := offIndex + len(off)
		onOffset := strings.Index(memory[afterOffIndex:], on)

		if onOffset == -1 {
			break
		} else {
			afterOnIndex := afterOffIndex + onOffset + len(on)
			memory = memory[afterOnIndex:]
		}
	}

	return sum
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("You must provide an input file")
	}

	path := os.Args[1]
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read from file '%s': %v", path, err)
	}

	fmt.Printf("Sum of muls: %d\n", sumMuls(string(bytes)))
}

