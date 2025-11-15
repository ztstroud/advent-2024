package main

import (
	"reflect"
	"testing"
)

func TestGetPageIndices(t *testing.T) {
	pages := []int{75, 47, 61, 53, 29}
	actual := getPageIndices(pages)
	expected := map[int]int{
		75: 0,
		47: 1,
		61: 2,
		53: 3,
		29: 4,
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestValidate(t *testing.T) {
	pages := []int{10, 15, 20}
	rules := []Ordering{
		{ before: 10, after: 15 },
		{ before: 10, after: 20 },
	}

	actual := validate(pages, rules)
	expected := true

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestValidateNoRulesApply(t *testing.T) {
	pages := []int{10, 15, 20}
	rules := []Ordering{
		{ before: 5, after: 15 },
		{ before: 10, after: 25 },
	}

	actual := validate(pages, rules)
	expected := true

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestValidateInvalid(t *testing.T) {
	pages := []int{10, 15, 20}
	rules := []Ordering{
		{ before: 10, after: 15 },
		{ before: 20, after: 15 },
	}

	actual := validate(pages, rules)
	expected := false

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestParseOrdering(t *testing.T) {
	actual, err := parseOrdering("25|499")
	expected := Ordering{
		before: 25,
		after: 499,
	}

	if err != nil {
		t.Errorf("Got an error: %v", err)
	}

	if actual != expected {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestParseExtraPipe(t *testing.T) {
	_, err := parseOrdering("25|499|1")

	if err == nil {
		t.Errorf("Did not received expected error")
	}
}

func TestParsePages(t *testing.T) {
	actual, err := parsePages("25,499,33,75")
	expected := []int{ 25, 499, 33, 75 }

	if err != nil {
		t.Errorf("Got an error: %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestParsePagesInvalid(t *testing.T) {
	_, err := parsePages("25,4X9,33,75")

	if err == nil {
		t.Errorf("Did not received expected error")
	}
}

func TestBuildOrderingMap(t *testing.T) {
	rules := []Ordering{
		{ before: 16, after: 13 },
		{ before: 16, after: 21 },
		{ before: 29, after: 13 },
		{ before: 61, after: 29 },
	}

	actual := buildOrderingMap(rules)
	expected := OrderingMap{
		16: { 13: struct{}{}, 21: struct{}{} },
		29: { 13: struct{}{} },
		61: { 29: struct{}{} },
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestSortPages(t *testing.T) {
	pages := []int{ 61, 13, 29 }
	orderingMap := OrderingMap{
		16: { 13: struct{}{} },
		29: { 13: struct{}{} },
		61: { 29: struct{}{} },
	}

	sortPages(pages, orderingMap)
	actual := pages
	expected := []int{ 61, 29, 13 }

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

