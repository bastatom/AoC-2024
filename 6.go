package main

func simulateGuard(guardMap [][]byte, x, y, dir int) (int, bool) {
	visited := make(map[[2]int]struct{})
	visitedWDir := make(map[[3]int]struct{})
	for get2dVal(guardMap, x, y) != 0 {
		if _, ok := visitedWDir[[3]int{x, y, dir}]; ok {
			return 0, true
		}

		visited[[2]int{x, y}] = struct{}{}
		visitedWDir[[3]int{x, y, dir}] = struct{}{}

		nextX, nextY := move2d(x, y, dirs[dir])
		if get2dVal(guardMap, nextX, nextY) == '#' {
			dir = (dir + 1) % 4
			continue
		}

		x, y = nextX, nextY
	}

	return len(visited), false
}

func solve6a() {
	guardMap := read2dBytes()
	x, y := findMarkIn2dSlice(guardMap, '^')

	visited, _ := simulateGuard(guardMap, x, y, 0)
	printToOut(visited)
}

func solve6b() {
	guardMap := read2dBytes()
	x, y := findMarkIn2dSlice(guardMap, '^')
	dir := 0

	visited := make(map[[2]int]struct{})
	loopMakers := make(map[[2]int]struct{})
	for get2dVal(guardMap, x, y) != 0 {
		visited[[2]int{x, y}] = struct{}{}
		nextX, nextY := move2d(x, y, dirs[dir])
		if get2dVal(guardMap, nextX, nextY) == '#' {
			dir = (dir + 1) % 4
			continue
		}

		if _, passed := visited[[2]int{nextX, nextY}]; !passed && get2dVal(guardMap, nextX, nextY) == '.' {
			guardMap[nextY][nextX] = '#'
			_, loop := simulateGuard(guardMap, x, y, dir)
			if loop {
				loopMakers[[2]int{nextX, nextY}] = struct{}{}
			}
			guardMap[nextY][nextX] = '.'
		}

		x, y = nextX, nextY
	}

	printToOut(len(loopMakers))
}
