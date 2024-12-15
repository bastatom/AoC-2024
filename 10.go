package main

func getMapNum(m [][]int, i, j int) int {
	if i < 0 || i >= len(m) {
		return -1
	}
	if j < 0 || j >= len(m[i]) {
		return -1
	}
	return m[i][j]
}

func solve10a() {
	var tMap [][]int
	var s string
	for scan(&s) {
		a := make([]int, 0, len(s))
		for i := 0; i < len(s); i++ {
			a = append(a, int(s[i]-'0'))
		}
		tMap = append(tMap, a)
	}

	tops := make(map[[2]int]struct{})
	dirs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	var rec func(int, int)
	rec = func(i, j int) {
		num := getMapNum(tMap, i, j)
		if num == 9 {
			tops[[2]int{i, j}] = struct{}{}
		}

		for _, dir := range dirs {
			if num+1 == getMapNum(tMap, i+dir[0], j+dir[1]) {
				rec(i+dir[0], j+dir[1])
			}
		}
	}

	var res int64
	for i := 0; i < len(tMap); i++ {
		for j := 0; j < len(tMap[i]); j++ {
			if tMap[i][j] != 0 {
				continue
			}

			rec(i, j)
			res += int64(len(tops))
			clear(tops)
		}
	}

	printToOut(res)
}

func solve10b() {
	var tMap [][]int
	var s string
	for scan(&s) {
		a := make([]int, 0, len(s))
		for i := 0; i < len(s); i++ {
			a = append(a, int(s[i]-'0'))
		}
		tMap = append(tMap, a)
	}

	dirs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	var rec func(int, int) int64
	rec = func(i, j int) int64 {
		num := getMapNum(tMap, i, j)
		if num == 9 {
			return 1
		}

		var res int64
		for _, dir := range dirs {
			if num+1 == getMapNum(tMap, i+dir[0], j+dir[1]) {
				res += rec(i+dir[0], j+dir[1])
			}
		}
		return res
	}

	var res int64
	for i := 0; i < len(tMap); i++ {
		for j := 0; j < len(tMap[i]); j++ {
			if tMap[i][j] != 0 {
				continue
			}

			res += rec(i, j)
		}
	}

	printToOut(res)
}
