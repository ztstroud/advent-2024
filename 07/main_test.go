package main

import (
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

