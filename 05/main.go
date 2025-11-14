package main

import (
	"strconv"
	"strings"
)

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

func parseOrdering(src string) (Ordering, error) {
	pages := strings.SplitN(src, "|", 2)

	before, err := strconv.Atoi(pages[0])
	if err != nil {
		return Ordering{}, err
	}

	after, err := strconv.Atoi(pages[1])
	if err != nil {
		return Ordering{}, err
	}

	return Ordering{
		before: before,
		after: after,
	}, nil
}

func parsePages(src string) ([]int, error) {
	pageSrcs := strings.Split(src, ",")
	pages := make([]int, len(pageSrcs))

	for i, pageSrc := range pageSrcs {
		page, err := strconv.Atoi(pageSrc)
		if err != nil {
			return nil, err
		}

		pages[i] = page
	}

	return pages, nil
}

