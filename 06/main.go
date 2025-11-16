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

var DIRS = []Position{
	{ x: 0, y: -1 },
	{ x: 1, y: 0 },
	{ x: 0, y: 1 },
	{ x: -1, y: 0 },
}

func rotateDir(dir int) int {
	dir += 1
	if dir >= len(DIRS) {
		dir = 0
	}

	return dir
}

func getNextDir(field Field, pos Position, dir int) int {
	nextPos := pos.add(DIRS[dir])

	// We only ever have to turn twice, otherwise we must have passed through a
	// wall, which isn't possible
	for i := 0; i < 2 && wallAt(field, nextPos); i++ {
		dir = rotateDir(dir)
		nextPos = pos.add(DIRS[dir])
	}

	return dir
}

/*
Simulate the walk of a robot starting at the given position facing up

This assumes that the robot will eventually walk off the map. If that is not
true, this function will loop forever. This could be fixed by keeping track of
positions visited and the direction that was being faced. If the same state is
entered, the robot is stuck in a loop.
*/
func simulate(field Field, pos Position) int {
	// A visitDir is valid iff the corresponding field entry is visited
	visitDirs := make([][]int, len(field))
	for y, row := range field {
		visitDirs[y] = make([]int, len(row))
	}

	loopCount := 0
	dir := 0

	// Each loop represents the action of _entering_ the cell at pos
	for inBounds(field, pos) {
		newDir := getNextDir(field, pos, dir)
		newPos := pos.add(DIRS[newDir])

		newPosInBounds := inBounds(field, newPos)

		if newPosInBounds && meetsOldPath(field, pos, rotateDir(newDir), visitDirs) {
			//fmt.Printf("Can place at %v\n", newPos)
			loopCount += 1
		}

		if field[pos.y][pos.x] == EMPTY {
			field[pos.y][pos.x] = VISITED
			visitDirs[pos.y][pos.x] = dir
		}

		pos = newPos
		dir = newDir
	}

	return loopCount
}

/*
Checks if the robot would get on the same path at some point in the future

This does not just mean hitting the same cell, but in the same direction as
well. In this case, the robot loops without needing to place any object.
*/
func meetsOldPath(field Field, pos Position, dir int, visitDirs [][]int) bool {
	for inBounds(field, pos) {
		newDir := getNextDir(field, pos, dir)
		newPos := pos.add(DIRS[newDir])

		if field[pos.y][pos.x] == VISITED && visitDirs[pos.y][pos.x] == dir {
			return true
		}

		pos = newPos
		dir = newDir
	}

	return false
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

