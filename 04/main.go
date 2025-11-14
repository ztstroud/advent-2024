package main

import (
	"slices"
	"strings"
)

func revrseByBytes(text string) string {
	bytes := []byte(text)
	slices.Reverse(bytes)
	return string(bytes)
}

/*
Counts occurrences of the query string both forwards and backwards.
*/
func countOccurrences(text, query string) int {
	reverseQuery := revrseByBytes(query)
	return strings.Count(text, query) + strings.Count(text, reverseQuery)
}

