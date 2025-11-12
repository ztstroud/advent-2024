package main

import (
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

