package main

import "slices"

func getGardenCrop(garden [][]byte, i, j int) byte {
	if i < 0 || i >= len(garden) {
		return 0
	}
	if j < 0 || j >= len(garden[i]) {
		return 0
	}
	return garden[i][j]
}

func solve12a() {
	var garden [][]byte
	var s string
	for scan(&s) {
		garden = append(garden, []byte(s))
	}

	dirs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	var dfs func(int, int, byte, byte) (int64, int64)
	dfs = func(i, j int, crop byte, mark byte) (area, perimeter int64) {
		area = 1
		garden[i][j] = mark
		for _, dir := range dirs {
			switch getGardenCrop(garden, i+dir[0], j+dir[1]) {
			case crop:
				subArea, subPerimeter := dfs(i+dir[0], j+dir[1], crop, mark)
				area += subArea
				perimeter += subPerimeter
			case mark:
				// noop - already visited with crop: do not add perimeter but do not start another dfs
			default:
				perimeter++
			}
		}

		return area, perimeter
	}

	var res int64
	for i := 0; i < len(garden); i++ {
		for j := 0; j < len(garden[i]); j++ {
			if garden[i][j] == '.' {
				continue
			}

			// first run the counting round
			area, perimeter := dfs(i, j, garden[i][j], ':')

			res += area * perimeter

			// next run the clearing round
			dfs(i, j, ':', '.')
		}
	}

	printToOut(res)
}

func safeMod[T int | int64](val T, mod T) T {
	return (val%mod + mod) % mod
}

func solve12b() {
	var garden [][]byte
	var s string
	var wallCheckDirsDone [][][]int
	for scan(&s) {
		garden = append(garden, []byte(s))
		wallCheckDirsDone = append(wallCheckDirsDone, make([][]int, len(s)))
	}

	dirs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	var markWall func(int, int, int, byte, byte) int64
	markWall = func(i, j int, dir int, crop byte, mark byte) int64 {
		if getGardenCrop(garden, i, j) != crop {
			return 0
		}

		if slices.Contains(wallCheckDirsDone[i][j], dir) {
			return 0
		}

		wallCheckDirsDone[i][j] = append(wallCheckDirsDone[i][j], dir)
		if next := getGardenCrop(garden, i+dirs[dir][0], j+dirs[dir][1]); next == crop || next == mark {
			return 0
		}

		for _, nextDir := range []int{safeMod(dir-1, len(dirs)), safeMod(dir+1, len(dirs))} {
			markWall(i+dirs[nextDir][0], j+dirs[nextDir][1], dir, crop, mark)
		}
		return 1
	}

	var dfs func(int, int, byte, byte) (int64, int64)
	dfs = func(i, j int, crop byte, mark byte) (area int64, perimeter int64) {
		if getGardenCrop(garden, i, j) != crop {
			return 0, 0
		}

		for k := range dirs {
			perimeter += markWall(i, j, k, crop, mark)
		}

		garden[i][j] = mark
		for _, dir := range dirs {
			subArea, subPerimeter := dfs(i+dir[0], j+dir[1], crop, mark)
			area += subArea
			perimeter += subPerimeter
		}
		return 1 + area, perimeter
	}

	var res int64
	for i := 0; i < len(garden); i++ {
		for j := 0; j < len(garden[i]); j++ {
			if garden[i][j] == '.' {
				continue
			}

			area, perimeter := dfs(i, j, garden[i][j], ':')
			res += area * perimeter
			dfs(i, j, ':', '.')
		}
	}

	printToOut(res)
}
