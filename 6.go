package main

func getMapStart(m [][]byte) (x, y int) {
	for i := range len(m) {
		for j := range m[i] {
			if m[i][j] == '^' {
				m[i][j] = '.'
				return i, j
			}
		}
	}
	return -1, -1
}

func solve6a() {
	var m [][]byte
	var s string
	for scan(&s) {
		m = append(m, []byte(s))
	}

	var res int64
	x, y := getMapStart(m)

	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	dirIdx := 0
	for {
		if m[x][y] == '.' {
			res++
			m[x][y] = 'X'
		}

		nextX, nextY := x+dirs[dirIdx][0], y+dirs[dirIdx][1]
		if nextX < 0 || nextX >= len(m) ||
			nextY < 0 || nextY >= len(m) {
			break
		}

		if m[nextX][nextY] == '#' {
			dirIdx = (dirIdx + 1) % 4
			continue
		}

		x, y = nextX, nextY
	}

	printToOut(res)
}
