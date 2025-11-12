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
	return reportFailsAt(report) == -1
}

func diffIsSafe(diff int, positive bool) bool {
	const maxDiff = 3

	if diff == 0 {
		return false
	}

	if diff < -maxDiff || diff > maxDiff {
		return false
	}

	if diff > 0 != positive {
		return false
	}

	return true
}

func reportFailsAt(report []int) (errorAt int) {
	if len(report) <= 1 {
		return -1
	}

	// If the first diff is zero, it will be immediately rejected and the
	// direction of the sequence will be irrelevant
	positive := report[1] - report[0] > 0

	for i := range(len(report) - 1) {
		diff := report[i + 1] - report[i]
		if !diffIsSafe(diff, positive) {
			return i + 1
		}
	}

	return -1
}

func cloneWithout(report []int, i int) []int {
	clone := make([]int, 0, len(report) - 1)

	clone = append(clone, report[:i]...)
	clone = append(clone, report[i + 1:]...)

	return clone
}

func isReportSafeDampened(report []int) bool {
	if len(report) <= 1 {
		return true
	}

	errorAt := reportFailsAt(report)
	if errorAt == -1 {
		return true
	}

	removePrevious := errorAt >= 2 && isReportSafe(cloneWithout(report, errorAt - 2))
	removeBefore := isReportSafe(cloneWithout(report, errorAt - 1))
	removeAfter := isReportSafe(cloneWithout(report, errorAt))

	return removePrevious || removeBefore || removeAfter
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
	dampenedCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report, err := parseReport(scanner.Text())
		if err != nil {
			log.Fatal("Error reading report: ", scanner.Text(), err)
		}

		if isReportSafe(report) {
			safeCount += 1
		}

		if isReportSafeDampened(report) {
			dampenedCount += 1
		}
	}

	fmt.Printf("Safe reports: %d\n", safeCount)
	fmt.Printf("Safe dampened reports: %d\n", dampenedCount)
}

