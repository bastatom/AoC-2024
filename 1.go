package main

import "slices"

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func solve1a() {
	var a, b []int64
	var x, y int64
	for scan(&x, &y) {
		a = append(a, x)
		b = append(b, y)
	}

	slices.Sort(a)
	slices.Sort(b)

	var res int64
	for i := 0; i < len(a); i++ {
		res += abs(a[i] - b[i])
	}

	printToOut(res)
}
func solve1b() {
	var a, b []int64
	var x, y int64
	for scan(&x, &y) {
		a = append(a, x)
		b = append(b, y)
	}

	bNums := make(map[int64]int64)
	for _, bNum := range b {
		bNums[bNum]++
	}

	var res int64
	for _, aNum := range a {
		res += aNum * bNums[aNum]
	}

	printToOut(res)
}
