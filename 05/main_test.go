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

