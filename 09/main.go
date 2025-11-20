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

	// There can be empty blocks at the end which we should ignore
	for blocks[end] == -1 {
		end -= 1
	}

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

func computeChecksumFromDiskMap(diskMap []int) int {
	start :=  0
	startOffset := 0
	end := len(diskMap) - 1
	endOffset := diskMap[end] - 1

	incrementStart := func() {
		startOffset += 1

		// If a segment has zero length, this will immediately be true and we
		// will move on to the next
		for startOffset >= diskMap[start] {
			start += 1

			if start >= len(diskMap) {
				break
			}

			startOffset = 0
		}
	}

	decrementEnd := func() {
		endOffset -= 1

		// Similarly, if a segment has zero length this will move past it
		for endOffset < 0 {
			end -= 1

			if end < 0 {
				break
			}

			endOffset = diskMap[end] - 1
		}
	}

	// Determining if a segment is a file or empty is now a check against the
	// segment index
	for end % 2 == 1 {
		decrementEnd()
	}

	startNotAfterEnd := func() bool {
		if start != end {
			return start < end
		}

		return startOffset <= endOffset
	}

	// Separately keeping track of the index is easier than re-calculating it
	// from the disk map each iteration
	checksum, i := 0, 0
	for startNotAfterEnd() {
		// Just like the other solution, if the start pointer is empty we take
		// from the end pointer
		if start % 2 == 1 {
			checksum += i * end / 2
			decrementEnd()

			// This is not so simple to replace. It looks like you could do it
			// in a single step, but the loop actually is to account for empty
			// file blocks after free spaces, like [1 1 0 1 1]. When pointing to
			// the end, decrementing puts you in a free block. Decrementing
			// again lands you on the first free segment, because the file
			// segment was skipped
			//
			// It is doing too much work for large free segments.
			for end % 2 == 1 {
				decrementEnd()
			}
		} else {
			checksum += i * start / 2
		}

		// We still always increment the start position, as we either used it or
		// a value from the end in its place
		incrementStart()
		i += 1
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
	checksum2 := computeChecksumFromDiskMap(blocks)

	fmt.Printf("            Checksum: %d\n", checksum)
	fmt.Printf("Checksum from blocks: %d\n", checksum2)
}

