package main

import (
	"reflect"
	"testing"
)

func TestParseField(t *testing.T) {
	input := []string{
		"#..",
		"..#",
		"^#.",
	}

	field, pos := parseField(input)

	expectedField := Field{
		{ WALL,  EMPTY, EMPTY },
		{ EMPTY, EMPTY, WALL },
		{ EMPTY, WALL,  EMPTY },
	}
	expectedPos := Position{ x: 0, y: 2 }

	if !reflect.DeepEqual(expectedField, field) {
		t.Errorf("Expected %v to be %v", field, expectedField)
	}

	if expectedPos != pos {
		t.Errorf("Expected %v to be %v", pos, expectedPos)
	}
}

