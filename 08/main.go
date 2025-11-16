package main

type City [][]byte

func parseCity(src []string) City {
	city := make([][]byte, len(src))
	for y, line := range src {
		city[y] = make([]byte, len(line))
		for x, char := range []byte(line) {
			if char != '.' {
				city[y][x] = char
			}
		}
	}

	return city
}

