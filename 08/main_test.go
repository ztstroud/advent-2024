package main

import (
	"reflect"
	"testing"
)

func TestParseCity(t *testing.T) {
	input := []string{
		".0..",
		"....",
		"..A.",
		"...A",
	}

	city := parseCity(input)

	expectedCity := City{
		{ 0, '0', 0, 0 },
		{ 0, 0, 0, 0 },
		{ 0, 0, 'A', 0 },
		{ 0, 0, 0, 'A' },
	}

	if !reflect.DeepEqual(city, expectedCity) {
		t.Errorf("Expected %v to be %v", city, expectedCity)
	}
}

