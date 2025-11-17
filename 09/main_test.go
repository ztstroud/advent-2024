package main

import (
	"reflect"
	"testing"
)

func TestExpandDiskMap(t *testing.T) {
	diskMap := []int{ 1, 2, 3, 4, 5 }
	actual := expandDiskMap(diskMap)

	expected := []int{
		0,
		-1, -1,
		1, 1, 1,
		-1, -1, -1, -1,
		2, 2, 2, 2, 2,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestComputeChecksum(t *testing.T) {
	diskMap := []int{ 1, 2, 3, 4, 5 }
	actual := computeChecksumForUncompactedBlocks(expandDiskMap(diskMap))

	expected := 98

	if actual != expected {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

