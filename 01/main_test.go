package main

import (
	"testing"
)


func TestListDistEmpty(t *testing.T) {
	left := []int{}
	right := []int{}
	expected := 0

	dist := listDist(left, right)

	if (expected != dist) {
		t.Errorf("Expected %d to be %d", expected, dist)
	}
}

func TestListDistSample(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}
	expected := 11

	dist := listDist(left, right)

	if (expected != dist) {
		t.Errorf("Expected %d to be %d", expected, dist)
	}
}

