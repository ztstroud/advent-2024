package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseEquation(t *testing.T) {
	testValue, values, err := parseEquation("3267: 81 40 27")

	expectedTestValue := 3267
	expectedValues := []int{81, 40, 27}

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if testValue != expectedTestValue {
		t.Errorf("Expected %v to be %v", testValue, expectedTestValue)
	}

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected %v to be %v", values, expectedValues)
	}
}

func TestDigitsIn(t *testing.T) {
	cases := []struct{ val, digits int }{
		{ 1, 1 },
		{ 9, 1 },
		{ 10, 2 },
		{ 99, 2 },
		{ 100, 3 },
		{ 999, 3 },
		{ 1000, 4 },
		{ 9999, 4 },
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d", c.val), func(t *testing.T) {
			digits := digitsIn(c.val)
			if digits != c.digits {
				t.Errorf("Expected %v to be %v", digits, c.digits)
			}
		})
	}
}

func TestSolveable(t *testing.T) {
	cases := []struct{
		src string
		solvable bool
	}{
		{ "190: 10 19", true },
		{ "3267: 81 40 27", true },
		{ "83: 17 5", false },
		{ "156: 15 6", false },
		{ "7290: 6 8 6 15", false },
		{ "161011: 16 10 13", false },
		{ "192: 17 8 14", false },
		{ "21037: 9 7 18 13", false },
		{ "292: 11 6 16 20", true },
	}

	for _, c := range cases {
		t.Run(c.src, func(t *testing.T) {
			testValue, values, err := parseEquation(c.src)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			solvable := isSolvable(testValue, values)
			if solvable != c.solvable {
				t.Errorf("Expected %v to be %v", solvable, c.solvable)
			}
		})
	}
}

