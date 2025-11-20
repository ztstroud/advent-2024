package main

import (
	"fmt"
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
	funcs := map[string]func([]int) int{
		"ExpandMemory": func(diskMap []int) int {
			blocks := expandDiskMap(diskMap)
			return computeChecksumForUncompactedBlocks(blocks)
		},
		"FromDiskMap": computeChecksumFromDiskMap,
	}

	cases := map[string]struct{
		diskMap string
		checksum int
	}{
		"GivenSmall": { "12345", 60 },
		"GivenMedium": { "2333133121414131402", 1928 },
		"Endsfree": { "1342", 10 },
	}

	for fnName, fn := range funcs {
		for name, c := range cases {
			t.Run(fmt.Sprintf("%s/%s", fnName, name), func(t *testing.T) {
				diskMap, err := parseDiskMap([]byte(c.diskMap))
				if err != nil {
					t.Errorf("Failed to parse disk map: %s (This is an error in the test)", c.diskMap)
				}

				actual := fn(diskMap)
				if actual != c.checksum {
					t.Errorf("Expected %v to be %v", actual, c.checksum)
				}
			})
		}
	}

}

