package main

func rotate90(rows []string) []string {
	rotated := make([]string, 0, len(rows[0]))

	for x := range len(rows[0]) {
		rotatedRow := make([]byte, len(rows))

		for y := range rows {
			rotatedRow[y] = rows[len(rows) - 1 - y][x];
		}

		rotated = append(rotated, string(rotatedRow))
	}

	return rotated
}

