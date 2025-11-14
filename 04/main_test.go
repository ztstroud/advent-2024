package main

import (
	"testing"
)

func TestReverseByBytes(t *testing.T) {
	input := "XMAS"
	expected := "SAMX"
	actual := revrseByBytes(input)

	if expected != actual {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

