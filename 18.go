package main

import (
	"fmt"
	"math"
	"sort"
)

func bfs(maze [][]byte, startX, startY, endX, endY int) int64 {
	var steps int64 = 0
	var queue = [][2]int{{startX, startY}}
	var nextQueue [][2]int
	for len(queue)+len(nextQueue) > 0 {
		if len(queue) == 0 {
			steps++
			queue = nextQueue
			nextQueue = nil
			continue
		}

		pos := queue[0]
		queue = queue[1:]

		if get2dVal(maze, pos[0], pos[1]) != '.' {
			continue
		}
		maze[pos[1]][pos[0]] = 'O'

		if pos[0] == endX && pos[1] == endY {
			return steps
		}

		for _, dir := range dirs {
			nextX, nextY := move2d(pos[0], pos[1], dir)
			if get2dVal(maze, nextX, nextY) != '.' {
				continue
			}
			nextQueue = append(nextQueue, [2]int{nextX, nextY})
		}
	}
	return math.MaxInt64
}

func simulateFallingBytes(memSpace [][]byte, fallingBytes [][2]int, width, height int) int64 {
	for i := 0; i <= height; i++ {
		for j := 0; j <= width; j++ {
			memSpace[i][j] = '.'
		}
	}

	for _, fallingByte := range fallingBytes {
		memSpace[fallingByte[1]][fallingByte[0]] = '#'
	}

	return bfs(memSpace, 0, 0, width, height)
}

func solve18a() {
	var width, height = 70, 70
	var memSpace [][]byte
	memSpace = make([][]byte, height+1)
	for i := 0; i <= height; i++ {
		memSpace[i] = make([]byte, width+1)
	}

	var x, y int
	var fallingBytes [][2]int
	for scanf("%d,%d\n", &x, &y) {
		fallingBytes = append(fallingBytes, [2]int{x, y})
	}

	printToOut(simulateFallingBytes(memSpace, fallingBytes[:1024], width, height))
}

func solve18b() {
	var width, height = 70, 70
	var memSpace [][]byte
	memSpace = make([][]byte, height+1)
	for i := 0; i <= height; i++ {
		memSpace[i] = make([]byte, width+1)
	}

	var x, y int
	var fallingBytes [][2]int
	for scanf("%d,%d\n", &x, &y) {
		fallingBytes = append(fallingBytes, [2]int{x, y})
	}

	firstInvalid := sort.Search(len(fallingBytes), func(i int) bool {
		return simulateFallingBytes(memSpace, fallingBytes[:i+1], width, height) == math.MaxInt64
	})

	printToOut(fmt.Sprintf("%d,%d", fallingBytes[firstInvalid][0], fallingBytes[firstInvalid][1]))
}
