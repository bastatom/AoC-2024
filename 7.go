package main

import (
	"strconv"
	"strings"
)

func canCalculateEquationA(acc int64, total int64, nums []int64) bool {
	if acc > total {
		return false
	}

	if len(nums) == 0 {
		return acc == total
	}

	nextTotals := []int64{acc + nums[0], acc * nums[0]}
	for _, nextTotal := range nextTotals {
		if canCalculateEquationA(nextTotal, total, nums[1:]) {
			return true
		}
	}

	return false
}

func solve7a() {
	var res int64

	for {
		line, _ := reader.ReadString('\n')
		if len(line) == 0 {
			break
		}

		idx := strings.Index(line, ":")
		total, _ := strconv.ParseInt(line[:idx], 10, 64)

		nums, _ := fscanLineToSlice[int64](strings.NewReader(line[idx+1:]))
		if canCalculateEquationA(nums[0], total, nums[1:]) {
			res += total
		}
	}

	printToOut(res)
}

func concatNum(a, b int64) int64 {
	tmp := b
	for tmp > 0 {
		a *= 10
		tmp /= 10
	}
	return a + b
}

func canCalculateEquationB(acc int64, total int64, nums []int64) bool {
	if acc > total {
		return false
	}

	if len(nums) == 0 {
		return acc == total
	}

	nextTotals := []int64{acc + nums[0], acc * nums[0], concatNum(acc, nums[0])}
	for _, nextTotal := range nextTotals {
		if canCalculateEquationB(nextTotal, total, nums[1:]) {
			return true
		}
	}

	return false
}

func solve7b() {
	var res int64

	for {
		line, _ := reader.ReadString('\n')
		if len(line) == 0 {
			break
		}

		idx := strings.Index(line, ":")
		total, _ := strconv.ParseInt(line[:idx], 10, 64)

		nums, _ := fscanLineToSlice[int64](strings.NewReader(line[idx+1:]))
		if canCalculateEquationB(nums[0], total, nums[1:]) {
			res += total
		}
	}

	printToOut(res)
}
