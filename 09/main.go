package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func parseDiskMap(src []byte) ([]int, error) {
	diskMap := make([]int, 0, len(src))

	for _, char := range src {
		if char == '\n' {
			continue
		}

		if char < '0' || char > '9' {
			return nil, errors.New(fmt.Sprintf("'%c' is not a valid digit", char))
		}

		diskMap = append(diskMap, int(char - '0'))
	}

	return diskMap, nil
}

/*
Expand a disk map into the memory it describes
*/
func expandDiskMap(diskMap []int) []int {
	totalBlocks := 0
	for _, size := range diskMap {
		totalBlocks += size
	}

	blocks := make([]int, totalBlocks)
	bi := 0

	for mi, size := range diskMap {
		isFile := mi % 2 == 0

		content := -1
		if isFile {
			fileId := mi / 2
			content = fileId
		}

		for range size {
			blocks[bi] = content
			bi += 1
		}
	}

	return blocks
}

func computeChecksumForUncompactedBlocks(blocks []int) int {
	start, end := 0, len(blocks) - 1
	checksum := 0

	for start <= end {
		if blocks[start] == -1 {
			checksum += start * blocks[end]
			end -= 1

			for blocks[end] == -1 {
				end -= 1
			}
		} else {
			checksum += start * blocks[start]
		}

		start += 1
	}

	return checksum
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("You must specify a file\n")
	}

	path := os.Args[1]
	src, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not read file: %s\n%v\n", path, err)
	}

	diskMap, err := parseDiskMap(src)
	if err != nil {
		log.Fatalf("Invalid input data: %v\n", err)
	}

	blocks := expandDiskMap(diskMap)
	checksum := computeChecksumForUncompactedBlocks(blocks)

	fmt.Printf("Checksum: %d\n", checksum)
}

