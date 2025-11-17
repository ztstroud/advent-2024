package main

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

