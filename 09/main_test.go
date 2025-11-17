package main

import (
	"reflect"
	"testing"
)

func TestParseDiskMap(t *testing.T) {
	diskMap, err := parseDiskMap([]byte("0123456789"))
	expected := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(diskMap, expected) {
		t.Errorf("Expected %v to be %v", diskMap, expected)
	}
}

func TestParseDiskMapIgnoreNewline(t *testing.T) {
	diskMap, err := parseDiskMap([]byte("012\n"))
	expected := []int{ 0, 1, 2 }

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(diskMap, expected) {
		t.Errorf("Expected %v to be %v", diskMap, expected)
	}
}

func TestParseDiskMapErrorOnInvalidChar(t *testing.T) {
	_, err := parseDiskMap([]byte("A"))

	if err == nil {
		t.Errorf("Expected an error")
	}
}

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

