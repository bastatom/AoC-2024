package main

func moveRobot(pX, pY, vX, vY int64, dimX, dimY int64, rounds int64) (int64, int64) {
	x := pX + (vX * rounds)
	y := pY + (vY * rounds)

	x = safeMod(x, dimX)
	y = safeMod(y, dimY)

	return x, y
}

func getQuadrant(x, y, dimX, dimY int64) (int64, int64, int64, int64) {
	var tl, tr, bl, br int64 = 1, 1, 1, 1
	if x <= dimX/2 {
		tr, br = 0, 0
	}
	if x >= dimX/2 {
		tl, bl = 0, 0
	}
	if y <= dimY/2 {
		bl, br = 0, 0
	}
	if y >= dimY/2 {
		tl, tr = 0, 0
	}
	return tl, tr, bl, br
}

func solve14a() {
	var pX, pY, vX, vY int64
	var tl, tr, bl, br int64
	for scanf("p=%d,%d v=%d,%d\n", &pX, &pY, &vX, &vY) {
		x, y := moveRobot(pX, pY, vX, vY, 101, 103, 100)
		q1, q2, q3, q4 := getQuadrant(x, y, 101, 103)
		tl += q1
		tr += q2
		bl += q3
		br += q4
	}
	printToOut(tl * tr * bl * br)
}

func clear2dSlice(s [][]byte) {
	for i := range s {
		for j := range s[i] {
			s[i][j] = '.'
		}
	}
}

func fill2dSlice(s [][]byte, toFill [][2]int64, mark byte) {
	for i := range toFill {
		s[toFill[i][1]][toFill[i][0]] = mark
	}
}

func print2dSlice(s [][]byte) {
	for i := range s {
		printToOut(string(s[i]))
	}
	printToOut("")
	writer.Flush()
}

func moveRobots(positions, velocities [][2]int64, dimX, dimY int64) {
	for i := range positions {
		newX, newY := moveRobot(positions[i][0], positions[i][1], velocities[i][0], velocities[i][1], dimX, dimY, 1)
		positions[i][0] = newX
		positions[i][1] = newY
	}
}

var dirs = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func get2dVal[T any](s [][]T, x, y int) T {
	var zero T
	if y < 0 || y >= len(s) {
		return zero
	}
	if x < 0 || x >= len(s[y]) {
		return zero
	}
	return s[y][x]
}

func markPartition(s [][]byte, x, y int, searchFor byte, mark byte) {
	if get2dVal(s, x, y) != searchFor {
		return
	}
	s[y][x] = mark
	for i := range dirs {
		markPartition(s, x+dirs[i][0], y+dirs[i][1], searchFor, mark)
	}
}

func countPartitions(s [][]byte) int {
	var partitions int
	for y := range s {
		for x := range s[y] {
			if s[y][x] == '*' {
				partitions++
				markPartition(s, x, y, '*', '#')
			}
		}
	}

	return partitions
}

func solve14b() {
	var pX, pY, vX, vY int64

	var robotPositions [][2]int64
	var robotVelocities [][2]int64

	for scanf("p=%d,%d v=%d,%d\n", &pX, &pY, &vX, &vY) {
		robotPositions = append(robotPositions, [2]int64{pX, pY})
		robotVelocities = append(robotVelocities, [2]int64{vX, vY})
	}

	var width, height int64 = 101, 103
	plane := make([][]byte, height)
	for i := range plane {
		plane[i] = make([]byte, width)
	}

	for cnt := 0; cnt <= 1e5; cnt++ {
		printToOut("************ Move ", cnt, "************")
		clear2dSlice(plane)
		fill2dSlice(plane, robotPositions, '*')

		partitions := countPartitions(plane)

		if partitions < 200 {
			printToOut("Partitions: ", partitions)
			print2dSlice(plane)
			printToOut(cnt)
			break
		}

		moveRobots(robotPositions, robotVelocities, width, height)
	}
}
