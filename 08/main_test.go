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

func TestGetAntennaGroups(t *testing.T) {
	input := []string{
		".0..",
		"....",
		"..A.",
		"...A",
	}

	groups := getAntennaGroups(parseCity(input))

	expectedGroups := AntennaGroups{
		'0': { { 1, 0 } },
		'A': { { 2, 2 }, { 3, 3 } },
	}

	if !reflect.DeepEqual(groups, expectedGroups) {
		t.Errorf("Expected %v to be %v", groups, expectedGroups)
	}
}

func TestCountAntinodes(t *testing.T) {
	input := []string{
		"....",
		".0..",
		".0..",
		"..A.",
		"...A",
	}

	count := countAntinodes(parseCity(input))
	expected := 3

	if count != expected {
		t.Errorf("Expected %v to be %v", count, expected)
	}
}

func TestCountAntinodesHarmonic(t *testing.T) {
	input := []string{
		".0..",
		"....",
		".0..",
		"..A.",
		"...A",
	}

	count := countAntinodesHarmonic(parseCity(input))
	expected := 8

	if count != expected {
		t.Errorf("Expected %v to be %v", count, expected)
	}
}

