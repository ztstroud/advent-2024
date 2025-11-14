package main

import (
	"reflect"
	"testing"
)

func TestGetPageIndices(t *testing.T) {
	pages := []int{75, 47, 61, 53, 29}
	actual := getPageIndices(pages)
	expected := map[int]int{
		75: 0,
		47: 1,
		61: 2,
		53: 3,
		29: 4,
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

