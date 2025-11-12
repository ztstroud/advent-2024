package main

import (
	"reflect"
	"testing"
)

func TestParseReportEmpty(t *testing.T) {
	input := ""
	expected := []int{}
	actual := parseReport(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v to equal %v\n", actual, expected)
	}
}

