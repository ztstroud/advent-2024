package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

Only the runes '#' and '^' are reserved, corresponding to a wall and the start
position. All other runes are treated as empty space, enabling you to annotate
the map.
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
func simulate(field Field, pos Position) int {
	DIRS := []Position{
		{ x: 0, y: -1 },
		{ x: 1, y: 0 },
		{ x: 0, y: 1 },
		{ x: -1, y: 0 },
	}

	// A visitDir is valid iff the corresponding field entry is visited
	visitDirs := make([][]int, len(field))
	for y, row := range field {
		visitDirs[y] = make([]int, len(row))
	}

	loopCount := 0
	dir := 0

	// Each loop represents the action of _entering_ the cell at pos
	for inBounds(field, pos) {
		newDir := dir
		newPos := pos.add(DIRS[newDir])

		rotatedDir := dir + 1
		if rotatedDir >= len(DIRS) {
			rotatedDir = 0
		}

		if wallAt(field, newPos) {
			newDir = rotatedDir
			newPos = pos.add(DIRS[newDir])

			// We only ever have to turn twice, otherwise we must have passed
			// through a wall, which isn't possible
			if wallAt(field, newPos) {
				newDir += 1
				if newDir >= len(DIRS) {
					newDir = 0
				}

				newPos = pos.add(DIRS[newDir])
			}
		}

		if field[pos.y][pos.x] == VISITED {
			if visitDirs[pos.y][pos.x] == rotatedDir && inBounds(field, newPos) {
				loopCount += 1
			}
		} else {
			field[pos.y][pos.x] = VISITED
			visitDirs[pos.y][pos.x] = newDir
		}

		pos = newPos
		dir = newDir
	}

	return loopCount
}

/*
Count the number of occurrences of the given query in the given field
*/
func countMatching(field Field, query byte) int {
	sum := 0
	for _, row := range field {
		for _, val := range row {
			if val == query {
				sum += 1
			}
		}
	}

	return sum
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("You must provide an input file\n")
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to read file: %s\n%v\n", path, err)
	}
	defer file.Close()

	fieldSrc := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fieldSrc = append(fieldSrc, scanner.Text())
	}

	field, pos := parseField(fieldSrc)
	simulate(field, pos)

	visitedCount := countMatching(field, VISITED)
	fmt.Printf("Visited %d unique spaces\n", visitedCount)
}

