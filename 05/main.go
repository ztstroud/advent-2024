package main

func getPageIndices(pages []int) map[int]int {
	indices := make(map[int]int)
	for index, page := range pages {
		indices[page] = index
	}

	return indices
}

type Ordering struct{
	before int
	after int
}

func validate(pages []int, rules []Ordering) bool {
	indices := getPageIndices(pages)

	for _, rule := range rules {
		beforeIndex, ok := indices[rule.before]
		if !ok {
			continue
		}

		afterIndex, ok := indices[rule.after]
		if !ok {
			continue
		}

		if beforeIndex >= afterIndex {
			return false
		}
	}

	return true
}

