package main

func vecDiff(a, b [2]int) [2]int {
	return [2]int{a[0] - b[0], a[1] - b[1]}
}

func vecAdd(a, b [2]int) [2]int {
	return [2]int{a[0] + b[0], a[1] + b[1]}
}

func vecValid(a [2]int, maxX, maxY int) bool {
	if a[0] < 0 || a[0] >= maxX {
		return false
	}
	if a[1] < 0 || a[1] >= maxY {
		return false
	}
	return true
}

func solve8a() {
	m := make(map[byte][][2]int)

	var x, y int

	var s string
	for scan(&s) {
		y = 0
		for y = 0; y < len(s); y++ {
			if s[y] == '.' {
				continue
			}

			m[s[y]] = append(m[s[y]], [2]int{x, y})
		}
		x++
	}

	res := make(map[[2]int]struct{})

	for _, ants := range m {
		for i := 0; i < len(ants); i++ {
			for j := i + 1; j < len(ants); j++ {
				antI := vecAdd(ants[i], vecDiff(ants[i], ants[j]))
				if vecValid(antI, x, y) {
					res[antI] = struct{}{}
				}
				antJ := vecAdd(ants[j], vecDiff(ants[j], ants[i]))
				if vecValid(antJ, x, y) {
					res[antJ] = struct{}{}
				}
			}
		}
	}

	printToOut(len(res))
}

func solve8b() {
	m := make(map[byte][][2]int)

	var x, y int

	var s string
	for scan(&s) {
		y = 0
		for y = 0; y < len(s); y++ {
			if s[y] == '.' {
				continue
			}

			m[s[y]] = append(m[s[y]], [2]int{x, y})
		}
		x++
	}

	res := make(map[[2]int]struct{})

	for _, ants := range m {
		for i := 0; i < len(ants); i++ {
			for j := i + 1; j < len(ants); j++ {
				antI := ants[i]
				res[antI] = struct{}{}
				for {
					nextAntI := vecAdd(antI, vecDiff(ants[i], ants[j]))
					if !vecValid(nextAntI, x, y) {
						break
					}

					res[nextAntI] = struct{}{}
					antI = nextAntI
				}

				antJ := ants[j]
				res[antJ] = struct{}{}
				for {
					nextAntJ := vecAdd(antJ, vecDiff(ants[j], ants[i]))
					if !vecValid(nextAntJ, x, y) {
						break
					}

					res[nextAntJ] = struct{}{}
					antJ = nextAntJ
				}
			}
		}
	}

	printToOut(len(res))
}
