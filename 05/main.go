package main

func getPageIndices(pages []int) map[int]int {
	indices := make(map[int]int)
	for index, page := range pages {
		indices[page] = index
	}

	return indices
}

