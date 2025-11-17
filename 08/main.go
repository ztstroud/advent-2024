package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type City [][]byte
type Position struct{ x, y int }
type AntennaGroups map[byte][]Position

const EMPTY = 0

func parseCity(src []string) City {
	city := make([][]byte, len(src))
	for y, line := range src {
		city[y] = make([]byte, len(line))
		for x, char := range []byte(line) {
			if char != '.' {
				city[y][x] = char
			} else {
				city[y][x] = EMPTY
			}
		}
	}

	return city
}

func getAntennaGroups(city City) AntennaGroups {
	groups := make(AntennaGroups)
	for y := range city {
		for x := range city[y] {
			antenna := city[y][x]
			if antenna == EMPTY {
				continue
			}

			group, ok := groups[antenna]
			if !ok {
				group = make([]Position, 0)
			}

			group = append(group, Position{ x, y })
			groups[antenna] = group
		}
	}

	return groups
}

func inBounds(city City, pos Position) bool {
	if pos.y < 0 || pos.y >= len(city) {
		return false
	}

	if pos.x < 0 || pos.x >= len(city[pos.y]) {
		return false
	}

	return true
}

func countAntinodes(city City) int {
	antinodes := make([][]bool, len(city))
	for y, row := range city {
		antinodes[y] = make([]bool, len(row))
	}

	groups := getAntennaGroups(city)
	for _, ps := range groups {
		for i, ip := range ps {
			for j, jp := range ps {
				if i == j {
					continue
				}

				antinodePos := Position{
					jp.x + jp.x - ip.x,
					jp.y + jp.y - ip.y,
				}

				if !inBounds(city, antinodePos) {
					continue
				}

				antinodes[antinodePos.y][antinodePos.x] = true
			}
		}
	}

	count := 0
	for _, row := range antinodes {
		for _, hasAntinode := range row {
			if hasAntinode {
				count += 1
			}
		}
	}

	return count
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a % b)
}

/*
Count harmonic antinodes in a city

To get harmonic antinodes, we need to do two additional things:
1. Reduce the offset to the smallest offset pointing in the same direction
2. Iterate along the offset to catch all positions
*/
func countAntinodesHarmonic(city City) int {
	antinodes := make([][]bool, len(city))
	for y, row := range city {
		antinodes[y] = make([]bool, len(row))
	}

	groups := getAntennaGroups(city)
	for _, ps := range groups {
		for i, from := range ps {
			for j := range i {
				to := ps[j]

				dx := to.x - from.x
				dy := to.y - from.y

				// The gcd allows us to find the smallest offset
				gcd := gcd(dx, dy)

				dx /= gcd
				dy /= gcd

				pos := from
				for inBounds(city, pos) {
					antinodes[pos.y][pos.x] = true
					pos = Position{ pos.x + dx, pos.y + dy }
				}

				pos = from
				for inBounds(city, pos) {
					antinodes[pos.y][pos.x] = true
					pos = Position{ pos.x - dx, pos.y - dy }
				}
			}
		}
	}

	count := 0
	for _, row := range antinodes {
		for _, hasAntinode := range row {
			if hasAntinode {
				count += 1
			}
		}
	}

	return count
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("You must specify a file\n")
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to read file: %s\n%v\n", path, err)
	}
	defer file.Close()

	src := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		src = append(src, scanner.Text())
	}

	city := parseCity(src)
	antinodeCount := countAntinodes(city)
	harmonicAntinodeCount := countAntinodesHarmonic(city)

	fmt.Printf("%d possible antinodes\n", antinodeCount)
	fmt.Printf("%d possible harmonic antinodes\n", harmonicAntinodeCount)
}

