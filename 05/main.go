package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func buildOrderingMap(rules []Ordering) map[int]map[int]struct{} {
	orderingMap := make(map[int]map[int]struct{})
	for _, rule := range rules {
		afters, ok := orderingMap[rule.before]
		if !ok {
			afters = make(map[int]struct{})
			orderingMap[rule.before] = afters
		}

		afters[rule.after] = struct{}{}
	}

	return orderingMap
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("You must provide an input file\n")
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to read from file: %s\n%v\n", path, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := make([]Ordering, 0)
	for scanner.Scan() {
		orderingSrc := scanner.Text()

		if orderingSrc == "" {
			break
		}

		ordering, err := parseOrdering(orderingSrc)
		if err != nil {
			log.Fatalf("Failed to parse ordering: %s\n%v\n", orderingSrc, err)
		}

		rules = append(rules, ordering)
	}

	sum := 0
	for scanner.Scan() {
		pagesSrc := scanner.Text()

		pages, err := parsePages(pagesSrc)
		if err != nil {
			log.Fatalf("Failed to parse pages: %s\n%v\n", pagesSrc, err)
		}

		if validate(pages, rules) {
			sum += pages[len(pages) / 2]
		}
	}

	fmt.Printf("Sum of middle pages of valid updates: %d\n", sum)
}

