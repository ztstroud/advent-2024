package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseEquation(src string) (int, []int, error) {
	parts := strings.SplitN(src, ": ", 2)

	testValue, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, nil, err
	}

	valueSrcs := strings.Fields(parts[1])
	values := make([]int, len(valueSrcs))
	
	for i, valueSrc := range valueSrcs {
		value, err := strconv.Atoi(valueSrc)
		if err != nil {
			return 0, nil, err
		}

		values[i] = value
	}

	return testValue, values, nil
}

func digitsIn(v int) int {
	return int(math.Log10(float64(v))) + 1
}

/*
Combine two number by placing the digits side by side

For example:
	concat(2, 8) == 28
	concat(10, 37) == 1037
	concat(8, 101) == 8101
*/
func concat(a, b int) int {
	return a * int(math.Pow10(digitsIn(b))) + b
}

func isSolvableRecursive(testValue, current int, values []int) bool {
	if len(values) == 0 {
		return testValue == current
	}

	value := values[0]
	nextValues := values[1:]

	return isSolvableRecursive(testValue, current + value, nextValues) ||
		isSolvableRecursive(testValue, current * value, nextValues)
}

func isSolvable(testValue int, values []int) bool {
	if len(values) == 0 {
		return false
	}

	return isSolvableRecursive(testValue, values[0], values[1:])
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Must provide an input file")
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file: %s\n%v\n", path, err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		src := scanner.Text()

		testValue, values, err := parseEquation(src)
		if err != nil {
			log.Fatalf("Failed to parse equation: %s\n%v\n", src, err)
		}

		slv := isSolvable(testValue, values)
		if slv {
			sum += testValue
		}
	}

	fmt.Printf("Sum of solvable equations: %d\n", sum)
}

