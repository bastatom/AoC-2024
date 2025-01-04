package main

import (
	"math"
)

func read2dBytes() [][]byte {
	var s string
	var res [][]byte
	for scan(&s) {
		res = append(res, []byte(s))
	}
	return res
}

func bfsDistances(graph [][]byte, endX, endY int) [][]int64 {
	distances := make([][]int64, len(graph))
	for i := range distances {
		distances[i] = make([]int64, len(graph[i]))
		for j := range distances[i] {
			distances[i][j] = math.MaxInt64
		}
	}

	distances[endY][endX] = 0
	var queue [][2]int
	var nextQueue [][2]int
	var distance int64 = 0
	queue = append(queue, [2]int{endX, endY})

	for len(queue)+len(nextQueue) > 0 {
		if len(queue) == 0 {
			distance++
			queue = nextQueue
			nextQueue = nil
		}

		field := queue[0]
		queue = queue[1:]

		x, y := field[0], field[1]
		if get2dVal(distances, x, y) < distance {
			continue
		}

		distances[y][x] = distance
		for _, dir := range dirs {
			nextX, nextY := move2d(x, y, dir)
			if get2dVal(graph, nextX, nextY) == '.' {
				nextQueue = append(nextQueue, [2]int{nextX, nextY})
			}
		}
	}

	return distances
}

func absDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff
}

func countValidRacetrackCheats(graph [][]byte, distances [][]int64, fromX, fromY, maxCheatLen int, saveToCount int64) int64 {
	var res int64
	for y := -maxCheatLen; y <= maxCheatLen; y++ {
		for x := -maxCheatLen; x <= maxCheatLen; x++ {
			nextX, nextY := fromX+x, fromY+y
			if get2dVal(graph, nextX, nextY) != '.' {
				continue
			}

			manDist := absDiff(fromX, nextX) + absDiff(fromY, nextY)
			if manDist > maxCheatLen {
				continue
			}

			saved := get2dVal(distances, fromX, fromY) - get2dVal(distances, nextX, nextY) - int64(manDist)
			if saved >= saveToCount {
				res++
			}
		}
	}
	return res
}

func solve20a() {
	raceTrack := read2dBytes()

	startX, startY := findMarkIn2dSlice(raceTrack, 'S')
	endX, endY := findMarkIn2dSlice(raceTrack, 'E')

	raceTrack[startY][startX] = '.'
	raceTrack[endY][endX] = '.'

	distancesToEnd := bfsDistances(raceTrack, endX, endY)

	var res int64 = 0
	for y := range raceTrack {
		for x := range raceTrack[y] {
			if get2dVal(raceTrack, x, y) != '.' {
				continue
			}

			res += countValidRacetrackCheats(raceTrack, distancesToEnd, x, y, 2, 100)
		}
	}

	printToOut(res)
}

func solve20b() {
	raceTrack := read2dBytes()

	startX, startY := findMarkIn2dSlice(raceTrack, 'S')
	endX, endY := findMarkIn2dSlice(raceTrack, 'E')

	raceTrack[startY][startX] = '.'
	raceTrack[endY][endX] = '.'

	distancesToEnd := bfsDistances(raceTrack, endX, endY)

	var res int64 = 0
	for y := range raceTrack {
		for x := range raceTrack[y] {
			if get2dVal(raceTrack, x, y) != '.' {
				continue
			}

			res += countValidRacetrackCheats(raceTrack, distancesToEnd, x, y, 20, 100)
		}
	}
	printToOut(res)
}
