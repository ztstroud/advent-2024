package main

const EMPTY byte = 0
const VISITED byte = 1
const WALL byte = 2

type Field [][]byte
type Position struct { x, y int }

/*
Parse the string source into a field and starting position

The starting position is NOT marked as visited
*/
func parseField(src []string) (Field, Position) {
	pos := Position{}

	field := make(Field, len(src))
	for y, row := range src {
		field[y] = make([]byte, len(row))

		for x, char := range row {
			switch char {
			case '#':
				field[y][x] = 2
			case '^':
				pos = Position{ x: x, y: y }
			}
		}
	}

	return field, pos
}

func inBounds(field Field, pos Position) bool {
	if pos.y < 0 || pos.y >= len(field) {
		return false
	}

	if pos.x < 0 || pos.x >= len(field[pos.y]) {
		return false
	}

	return true
}

