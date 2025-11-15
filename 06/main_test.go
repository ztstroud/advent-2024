package main

import (
	"reflect"
	"testing"
)

func TestPositionAdd(t *testing.T) {
	sum := Position{1, 2}.add(Position{3, 4})
	expected := Position{4, 6}

	if expected != sum {
		t.Errorf("Expected %v to be %v", sum, expected)
	}
}

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

func TestInBounds(t *testing.T) {
	field := Field{
		{ WALL,  EMPTY, EMPTY },
		{ EMPTY, EMPTY, WALL },
		{ EMPTY, WALL,  EMPTY },
	}

	pos := Position{ x: 1, y: 1 }

	if !inBounds(field, pos) {
		t.Errorf("Expected %v to be out of bounds", pos)
	}
}

func TestInBoundsOutOfBounds(t *testing.T) {
	field := Field{
		{ WALL,  EMPTY, EMPTY },
		{ EMPTY, EMPTY, WALL },
		{ EMPTY, WALL,  EMPTY },
	}

	cases := []struct{
		name string
		pos Position
	}{
		{ name: "BeforeY", pos: Position{ y: -1 } },
		{ name: "AfterY", pos: Position{ y: 3 } },
		{ name: "BeforeX", pos: Position{ x: -1 } },
		{ name: "AfterX", pos: Position{ x: 3 } },
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if inBounds(field, testCase.pos) {
				t.Errorf("Expected %v to be out of bounds", testCase.pos)
			}
		})
	}
}

func TestWallAtWithWall(t *testing.T) {
	field := Field{
		{ WALL,  EMPTY, EMPTY },
		{ EMPTY, EMPTY, WALL },
		{ EMPTY, WALL,  EMPTY },
	}

	pos := Position{ x: 2, y: 1 }

	if !wallAt(field, pos) {
		t.Errorf("Missed wall at %v", pos)
	}
}

func TestWallAtNoWall(t *testing.T) {
	field := Field{
		{ WALL,  EMPTY, EMPTY },
		{ EMPTY, EMPTY, WALL },
		{ EMPTY, WALL,  EMPTY },
	}

	pos := Position{ x: 2, y: 2 }

	if wallAt(field, pos) {
		t.Errorf("Mistaken wall at %v", pos)
	}
}

func TestWallAtOutOfBounds(t *testing.T) {
	field := Field{
		{ WALL,  EMPTY, EMPTY },
		{ EMPTY, EMPTY, WALL },
		{ EMPTY, WALL,  EMPTY },
	}

	pos := Position{ x: 3, y: 2 }

	if wallAt(field, pos) {
		t.Errorf("Mistaken wall at %v", pos)
	}
}

func TestSimulate(t *testing.T) {
	field := Field{
		{ WALL,  EMPTY, EMPTY },
		{ EMPTY, EMPTY, WALL },
		{ EMPTY, WALL,  EMPTY },
	}

	pos := Position{0, 2}
	simulate(field, pos)

	expected := Field{
		{ WALL,    EMPTY,   EMPTY },
		{ VISITED, VISITED, WALL },
		{ VISITED, WALL,    EMPTY },
	}

	if !reflect.DeepEqual(expected, field) {
		t.Errorf("Expected %v to be %v", field, expected)
	}
}

