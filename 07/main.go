package main

import (
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

