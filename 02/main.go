package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseReport(report string) ([]int, error) {
	strVals := strings.Fields(report)
	vals := make([]int, 0, len(strVals))

	for _, strVal := range(strVals) {
		val, err := strconv.Atoi(strVal)
		if err != nil {
			return nil, err
		}

		vals = append(vals, val)
	}

	return vals, nil
}

func isReportSafe(report []int) bool {
	if len(report) <= 1 {
		return true
	}

	// If the first diff is zero, it will be immediately rejected and the
	// direction of the sequence will be irrelevant
	positive := report[1] - report[0] > 0

	const maxDiff = 2
	for i := range(len(report) - 1) {
		diff := report[i + 1] - report[i]
		if diff == 0 {
			return false
		}

		if diff < -maxDiff || diff > maxDiff {
			return false
		}

		if diff > 0 != positive {
			return false
		}
	}

	return true
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

	safeCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report, err := parseReport(scanner.Text())
		if err != nil {
			log.Fatal("Error reading report: ", scanner.Text(), err)
		}

		if isReportSafe(report) {
			safeCount += 1
		}
	}

	fmt.Printf("Safe reports: %d\n", safeCount)
}

