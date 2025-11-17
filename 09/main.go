package main

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
				end += 1
			}
		} else {
			checksum += start * blocks[start]
		}

		start += 1
	}

	return checksum
}

