package main

import (
	"slices"
	"strings"
)

func pageOrderNum(m map[int64][]int64, nums []int64) int64 {
	in := make(map[int64]struct{})

	for _, num := range nums {
		for _, val := range m[num] {
			if _, ok := in[val]; ok {
				return 0
			}
		}

		in[num] = struct{}{}
	}

	return nums[len(nums)/2]
}

func solve5a() {
	m := make(map[int64][]int64)

	var s string
	var x, y int64
	for scanf("%v\n", &s) {
		if len(s) == 0 {
			break
		}
		sscanf(s, "%d|%d", &x, &y)
		if !slices.Contains(m[x], y) {
			m[x] = append(m[x], y)
		}
	}

	var res int64
	for scanf("%v\n", &s) {
		s = strings.ReplaceAll(s, ",", " ")
		nums, _ := fscanLineToSlice[int64](strings.NewReader(s))

		res += pageOrderNum(m, nums)
	}

	printToOut(res)
}

func removeNumFromDeps(deps map[int64][]int64, num int64) {
	for key, dep := range deps {
		if idx := slices.Index(dep, num); idx != -1 {
			deps[key] = append(deps[key][:idx], deps[key][idx+1:]...)
		}
		if len(deps[key]) == 0 {
			delete(deps, key)
		}
	}
}

func reorderPages(m map[int64][]int64, nums []int64) []int64 {
	reordered := make([]int64, 0, len(nums))

	deps := make(map[int64][]int64)
	for _, num := range nums {
		for _, dep := range m[num] {
			if slices.Contains(nums, dep) {
				deps[dep] = append(deps[dep], num)
			}
		}
	}

	for range len(nums) {
		for _, num := range nums {
			if slices.Contains(reordered, num) {
				continue
			}

			if _, ok := deps[num]; !ok {
				reordered = append(reordered, num)
				removeNumFromDeps(deps, num)
				break
			}
		}
	}

	return reordered
}

func solve5b() {
	m := make(map[int64][]int64)

	var s string
	var x, y int64
	for scanf("%v\n", &s) {
		if len(s) == 0 {
			break
		}
		sscanf(s, "%d|%d", &x, &y)
		if !slices.Contains(m[x], y) {
			m[x] = append(m[x], y)
		}
	}

	var res int64
	for scanf("%v\n", &s) {
		s = strings.ReplaceAll(s, ",", " ")
		nums, _ := fscanLineToSlice[int64](strings.NewReader(s))
		if pageOrderNum(m, nums) != 0 {
			continue
		}

		nums = reorderPages(m, nums)
		res += pageOrderNum(m, nums)
	}

	printToOut(res)
}
