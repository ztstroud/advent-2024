package main

const EMPTY byte = 0
const VISITED byte = 1
const WALL byte = 2

type Field [][]byte
type Position struct { x, y int }

func (pos Position) add(other Position) Position {
	return Position{
		x: pos.x + other.x,
		y: pos.y + other.y,
	}
}

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
				field[y][x] = WALL
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

/*
Check if there is a wall at a given position

If the position is out of bounds, it is considered empty.
*/
func wallAt(field Field, pos Position) bool {
	if !inBounds(field, pos) {
		return false
	}

	return field[pos.y][pos.x] == WALL
}

/*
Simulate the walk of a robot starting at the given position facing up

This assumes that the robot will eventually walk off the map. If that is not
true, this function will loop forever. This could be fixed by keeping track of
positions visited and the direction that was being faced. If the same state is
entered, the robot is stuck in a loop.
*/
func simulate(field Field, pos Position) {
	DIRS := []Position{
		{ x: 0, y: -1 },
		{ x: 1, y: 0 },
		{ x: 0, y: 1 },
		{ x: -1, y: 0 },
	}

	dir := 0
	for inBounds(field, pos) {
		newPos := pos.add(DIRS[dir])

		if wallAt(field, newPos) {
			dir += 1
			if dir >= len(DIRS) {
				dir = 0
			}
		} else {
			field[pos.y][pos.x] = VISITED
			pos = newPos
		}
	}
}

