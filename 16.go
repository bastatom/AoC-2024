package main

import (
	"container/heap"
	"math"
)

type Prioriter interface {
	Priority() int
}
type IntHeap[val Prioriter] struct {
	slice []val
}

func (gh *IntHeap[val]) Len() int {
	return len(gh.slice)
}
func (gh *IntHeap[val]) Less(i, j int) bool {
	return gh.slice[i].Priority() < gh.slice[j].Priority()
}
func (gh *IntHeap[val]) Swap(i, j int) {
	gh.slice[i], gh.slice[j] = gh.slice[j], gh.slice[i]
}
func (gh *IntHeap[val]) Push(x any) {
	gh.slice = append(gh.slice, x.(val))
}
func (gh *IntHeap[val]) Pop() any {
	tmp := gh.slice[len(gh.slice)-1]
	gh.slice = gh.slice[:len(gh.slice)-1]
	return tmp
}

type mazePos struct {
	x, y  int
	dir   int
	score int
}

func (p mazePos) Priority() int {
	return p.score
}

func solve16a() {
	var maze [][]byte
	var s string
	for scan(&s) {
		maze = append(maze, []byte(s))
	}

	x, y := findMarkIn2dSlice(maze, 'S')
	endX, endY := findMarkIn2dSlice(maze, 'E')

	pq := &IntHeap[mazePos]{}
	pq.Push(mazePos{x, y, 1, 0})
	heap.Init(pq)

	visited := make(map[[2]int]struct{}, len(maze)*len(maze[0]))

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(mazePos)
		score := cur.score
		x, y = cur.x, cur.y

		_, ok := visited[[2]int{x, y}]
		if ok {
			continue
		}
		visited[[2]int{x, y}] = struct{}{}

		if x == endX && y == endY {
			printToOut(score)
			return
		}

		for _, turn := range []int{-1, 0, 1} {
			dir := safeMod(cur.dir+turn, len(dirs))
			nextX, nextY := move2d(x, y, dirs[dir])
			if get2dVal(maze, nextX, nextY) == '#' {
				continue
			}

			scoreAdd := 1
			if turn != 0 {
				scoreAdd += 1000
			}

			heap.Push(pq, mazePos{nextX, nextY, dir, cur.score + scoreAdd})
		}
	}
}

func solve16b() {
	var maze [][]byte
	var tileScore [][][4]int
	var s string
	for scan(&s) {
		maze = append(maze, []byte(s))
		scores := make([][len(dirs)]int, len(s))
		for i := range scores {
			for dir := range dirs {
				scores[i][dir] = math.MaxInt32
			}
		}
		tileScore = append(tileScore, scores)
	}

	startX, startY := findMarkIn2dSlice(maze, 'S')
	endX, endY := findMarkIn2dSlice(maze, 'E')

	pq := &IntHeap[mazePos]{}
	pq.Push(mazePos{startX, startY, 1, 0})
	heap.Init(pq)

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(mazePos)
		score := cur.score
		x, y := cur.x, cur.y
		dir := cur.dir

		if tileScore[y][x][dir] < score {
			continue
		}
		tileScore[y][x][dir] = score

		for _, turn := range []int{-1, 0, 1} {
			dir = safeMod(cur.dir+turn, len(dirs))

			nextX, nextY := x, y
			if turn == 0 {
				nextX, nextY = move2d(x, y, dirs[dir])
				if get2dVal(maze, nextX, nextY) == '#' {
					continue
				}
			}

			scoreAdd := 1
			if turn != 0 {
				scoreAdd = 1000
			}

			heap.Push(pq, mazePos{nextX, nextY, dir, score + scoreAdd})
		}
	}

	var maxScore = math.MaxInt32
	for dir := range dirs {
		maxScore = min(maxScore, tileScore[endY][endX][dir])
	}

	bestSpots := make(map[[2]int]struct{}, len(maze)*len(maze[0]))
	bestSpots[[2]int{startX, startY}] = struct{}{}
	bestSpots[[2]int{endX, endY}] = struct{}{}

	var dfs func(int, int, int, int) bool
	dfs = func(x, y, dir, rem int) bool {
		if x == startX && y == startY {
			return true
		}

		var optimal bool
		for _, turn := range []int{-1, 0, 1} {
			nextDir := safeMod(dir+turn, len(dirs))
			nextX, nextY := move2d(x, y, dirs[nextDir])
			if get2dVal(maze, nextX, nextY) == '#' {
				continue
			}

			cost := 1
			if turn != 0 {
				cost += 1000
			}

			if tileScore[nextY][nextX][safeMod(nextDir-2, len(dirs))] > rem-cost {
				continue
			}

			optimal = dfs(nextX, nextY, nextDir, rem-cost) || optimal
		}

		if optimal {
			bestSpots[[2]int{x, y}] = struct{}{}
		}
		return optimal
	}

	for i := range dirs {
		dfs(endX, endY, i, maxScore)
	}
	printToOut(len(bestSpots))
}
