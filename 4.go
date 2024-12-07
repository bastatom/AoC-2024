package main

func getPuzzleByte(puzzle [][]byte, x, y int) byte {
	if x < 0 || x >= len(puzzle) {
		return ' '
	}
	if y < 0 || y >= len(puzzle[x]) {
		return ' '
	}
	return puzzle[x][y]
}

const xmasString = "XMAS"

func checkXMAS(puzzle [][]byte, x, y, xAdd, yAdd int) int64 {
	for i := range xmasString {
		if getPuzzleByte(puzzle, x+(i*xAdd), y+(i*yAdd)) != xmasString[i] {
			return 0
		}
	}

	return 1
}

func solve4a() {
	var puzzle [][]byte
	var s string
	for scan(&s) {
		puzzle = append(puzzle, []byte(s))
	}

	var res int64
	for x := 0; x < len(puzzle); x++ {
		for y := 0; y < len(puzzle[x]); y++ {
			if getPuzzleByte(puzzle, x, y) != 'X' {
				continue
			}

			for _, xAdd := range []int{-1, 0, 1} {
				for _, yAdd := range []int{-1, 0, 1} {
					if xAdd == 0 && yAdd == 0 {
						continue
					}

					res += checkXMAS(puzzle, x, y, xAdd, yAdd)
				}
			}
		}
	}

	printToOut(res)
}

func checkX_MAS(puzzle [][]byte, x, y int) int64 {
	for _, add := range []int{-1, 1} {
		tmp := string(getPuzzleByte(puzzle, x-1, y+add)) + string(getPuzzleByte(puzzle, x+1, y-add))
		if tmp != "MS" && tmp != "SM" {
			return 0
		}
	}

	return 1
}

func solve4b() {
	var puzzle [][]byte
	var s string
	for scan(&s) {
		puzzle = append(puzzle, []byte(s))
	}

	var res int64
	for x := 0; x < len(puzzle); x++ {
		for y := 0; y < len(puzzle[x]); y++ {
			if getPuzzleByte(puzzle, x, y) != 'A' {
				continue
			}

			res += checkX_MAS(puzzle, x, y)
		}
	}

	printToOut(res)
}
