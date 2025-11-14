package main

import (
	"reflect"
	"testing"
)

func TestRotate90(t *testing.T) {
	input := []string {
		"abc",
		"def",
		"ghi",
	}

	expected := []string {
		"gda",
		"heb",
		"ifc",
	}

	actual := rotate90(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

