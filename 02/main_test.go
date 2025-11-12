package main

import (
	"reflect"
	"testing"
)

func TestParseReportEmpty(t *testing.T) {
	input := ""
	expected := []int{}
	actual, err := parseReport(input)

	if err != nil {
		t.Errorf("Got an error: %v\n", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

func TestParseReportSingle(t *testing.T) {
	input := "5"
	expected := []int{5}
	actual, err := parseReport(input)

	if err != nil {
		t.Errorf("Got an error: %v\n", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

func TestParseReportMany(t *testing.T) {
	input := "5 2 4 3 1"
	expected := []int{5, 2, 4, 3, 1}
	actual, err := parseReport(input)

	if err != nil {
		t.Errorf("Got an error: %v\n", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

func TestIsReportSafeEmpty(t *testing.T) {
	input := []int{}
	expected := true
	actual := isReportSafe(input)

	if expected != actual {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

func TestIsReportSafeSingle(t *testing.T) {
	input := []int{1}
	expected := true
	actual := isReportSafe(input)

	if expected != actual {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

func TestIsReportSafeDuplicate(t *testing.T) {
	input := []int{1, 1}
	expected := false
	actual := isReportSafe(input)

	if expected != actual {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

func TestIsReportSafeLargeGapPositive(t *testing.T) {
	input := []int{1, 100}
	expected := false
	actual := isReportSafe(input)

	if expected != actual {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

func TestIsReportSafeLargeGapNegative(t *testing.T) {
	input := []int{100, 1}
	expected := false
	actual := isReportSafe(input)

	if expected != actual {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

func TestIsReportSafeAscendDescend(t *testing.T) {
	input := []int{1, 2, 1}
	expected := false
	actual := isReportSafe(input)

	if expected != actual {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

func TestIsReportSafeDescendAscend(t *testing.T) {
	input := []int{1, 2, 1}
	expected := false
	actual := isReportSafe(input)

	if expected != actual {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

func TestIsReportSafeAllAscend(t *testing.T) {
	input := []int{1, 2, 4}
	expected := true
	actual := isReportSafe(input)

	if expected != actual {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

func TestIsReportSafeAllDescend(t *testing.T) {
	input := []int{4, 2, 1}
	expected := true
	actual := isReportSafe(input)

	if expected != actual {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

func TestIsReportSafeDampenedFirst(t *testing.T) {
	input := []int{0, 3, 2, 1}
	expected := true
	actual := isReportSafeDampened(input)

	if expected != actual {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

func TestIsReportSafeDampenedSecond(t *testing.T) {
	input := []int{3, 0, 2, 1}
	expected := true
	actual := isReportSafeDampened(input)

	if expected != actual {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

func TestIsReportSafeDampenedThird(t *testing.T) {
	input := []int{3, 2, 0, 1}
	expected := true
	actual := isReportSafeDampened(input)

	if expected != actual {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

func TestIsReportSafeDampenedLast(t *testing.T) {
	input := []int{3, 2, 1, 0}
	expected := true
	actual := isReportSafeDampened(input)

	if expected != actual {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

