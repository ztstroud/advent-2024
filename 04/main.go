package main

import (
	"slices"
)

func revrseByBytes(text string) string {
	bytes := []byte(text)
	slices.Reverse(bytes)
	return string(bytes)
}

