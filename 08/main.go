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

