package main

func findMarkIn2dSlice[T comparable](s [][]T, mark T) (x, y int) {
	for y = range s {
		for x = range s[y] {
			if mark == s[y][x] {
				return x, y
			}
		}
	}

	return -1, -1
}

func getValFrom2dSlice[T comparable](s [][]T, x, y int) T {
	var zero T
	if y < 0 || y >= len(s) {
		return zero
	}
	if x < 0 || x >= len(s[y]) {
		return zero
	}
	return s[y][x]
}

func move2d(x, y int, dir [2]int) (int, int) {
	return x + dir[0], y + dir[1]
}

func moveToDir(move byte) [2]int {
	var idx int
	switch move {
	case '^':
		idx = 0
	case '>':
		idx = 1
	case 'v':
		idx = 2
	case '<':
		idx = 3
	default:
		panicf("moveToDir invalid move: %v", string(move))
	}

	return dirs[idx]
}

func attemptMoveA(puzzle [][]byte, x, y int, dir [2]int) (int, int, bool) {
	nextX, nextY := move2d(x, y, dir)
	nextVal := getValFrom2dSlice(puzzle, nextX, nextY)
	switch nextVal {
	case '#':
		// noop
		return x, y, false
	case '.':
		puzzle[nextY][nextX] = puzzle[y][x]
		puzzle[y][x] = '.'
		return nextX, nextY, true
	case 'O':
		_, _, moved := attemptMoveA(puzzle, nextX, nextY, dir)
		if moved {
			puzzle[nextY][nextX] = puzzle[y][x]
			puzzle[y][x] = '.'
			return nextX, nextY, true
		}
		return x, y, false
	default:
		panicf("attemptMove: invalid move %s", string(nextVal))
		return 0, 0, false
	}
}

func solve15a() {
	var puzzle [][]byte
	var s string
	for scanf("%v\n", &s) {
		puzzle = append(puzzle, []byte(s))
	}

	x, y := findMarkIn2dSlice(puzzle, '@')

	var moves []byte
	for scan(&s) {
		moves = append(moves, []byte(s)...)
	}

	for _, m := range moves {
		nextX, nextY, _ := attemptMoveA(puzzle, x, y, moveToDir(m))
		x, y = nextX, nextY
	}

	var res int64
	for y = range puzzle {
		for x = range puzzle[y] {
			if puzzle[y][x] != 'O' {
				continue
			}

			res += int64(100*y + x)
		}
	}
	printToOut(res)
}

func canMove(puzzle [][]byte, x, y int, dir [2]int) bool {
	nextX, nextY := move2d(x, y, dir)
	nextVal := getValFrom2dSlice(puzzle, nextX, nextY)
	switch nextVal {
	case '#':
		// noop
		return false
	case '.':
		return true
	}

	if nextVal != '[' && nextVal != ']' {
		panicf("canMove: invalid move %s", string(nextVal))
		return false
	}

	// horizontal move
	if dir[1] == 0 {
		return canMove(puzzle, nextX, nextY, dir)
	}

	// vertical (pushing 2 parts at once - both needs to be pushable for the move to be alright)
	if nextVal == '[' {
		return canMove(puzzle, nextX, nextY, dir) && canMove(puzzle, nextX+1, nextY, dir)
	} else {
		return canMove(puzzle, nextX, nextY, dir) && canMove(puzzle, nextX-1, nextY, dir)
	}
}

func performMoveB(puzzle [][]byte, x, y int, dir [2]int) (int, int) {
	nextX, nextY := move2d(x, y, dir)
	nextVal := getValFrom2dSlice(puzzle, nextX, nextY)

	// here we assume that the path is clear
	defer func() {
		puzzle[nextY][nextX] = puzzle[y][x]
		puzzle[y][x] = '.'
	}()

	switch nextVal {
	case '#':
		panicf("performMoveB: crashed into wall")
		return x, y
	case '.':
		return nextX, nextY
	}

	if nextVal != '[' && nextVal != ']' {
		panicf("canMove: invalid move %s", string(nextVal))
		return x, y
	}

	// horizontal move
	if dir[1] == 0 {
		performMoveB(puzzle, nextX, nextY, dir)
		return nextX, nextY
	}

	// vertical (pushing 2 parts at once - both needs to be pushable for the move to be alright)
	if nextVal == '[' {
		performMoveB(puzzle, nextX, nextY, dir)
		performMoveB(puzzle, nextX+1, nextY, dir)
		return nextX, nextY
	} else {
		performMoveB(puzzle, nextX, nextY, dir)
		performMoveB(puzzle, nextX-1, nextY, dir)
		return nextX, nextY
	}
}

func solve15b() {
	var puzzle [][]byte
	var s string
	for scanf("%v\n", &s) {
		var line = make([]byte, 0, len(s)*2)
		for i := range s {
			switch s[i] {
			case '#':
				line = append(line, '#', '#')
			case '.':
				line = append(line, '.', '.')
			case 'O':
				line = append(line, '[', ']')
			case '@':
				line = append(line, '@', '.')
			}
		}
		puzzle = append(puzzle, line)
	}

	x, y := findMarkIn2dSlice(puzzle, '@')

	var moves []byte
	for scan(&s) {
		moves = append(moves, []byte(s)...)
	}

	for _, m := range moves {
		if canMove(puzzle, x, y, moveToDir(m)) {
			x, y = performMoveB(puzzle, x, y, moveToDir(m))
		}
	}

	var res int64
	for y = range puzzle {
		for x = range puzzle[y] {
			if puzzle[y][x] != '[' {
				continue
			}
			res += int64(100*y + x)
		}
	}
	printToOut(res)
}
