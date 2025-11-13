package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	input := "mul(2,3)"
	expected := 6
	actual := sumMuls(input)

	if expected != actual {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

