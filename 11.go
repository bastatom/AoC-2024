package main

func numLen(a int64) int {
	var res int
	for a > 0 {
		a /= 10
		res++
	}
	return res
}

func splitNum(n int64, nLen int) (int64, int64) {
	var a, b int64 = n, 0
	half := nLen / 2
	var factor int64 = 1
	for range half {
		b += (a % 10) * factor
		factor *= 10
		a /= 10
	}
	return a, b
}

func blink(mem map[int64][]int64, marble int64, remains int, maxRemains int) int64 {
	if remains == 0 {
		return 1
	}
	if marble < 0 {
		panic("Negative marble!")
	}
	if mem[marble] == nil {
		mem[marble] = make([]int64, maxRemains+1)
	}
	if mem[marble][remains] != 0 {
		return mem[marble][remains]
	}

	var res int64
	if marble == 0 {
		res = blink(mem, 1, remains-1, maxRemains)
	} else if nLen := numLen(marble); nLen%2 == 0 {
		oldNum, newNum := splitNum(marble, nLen)
		res = blink(mem, oldNum, remains-1, maxRemains) + blink(mem, newNum, remains-1, maxRemains)
	} else {
		res = blink(mem, marble*2024, remains-1, maxRemains)
	}

	mem[marble][remains] = res
	return res
}

func solve11a() {
	nums, _ := scanLineToSlice[int64]()

	var res int64
	for i := 0; i < len(nums); i++ {
		res += blink(make(map[int64][]int64), nums[i], 25, 25)
	}
	printToOut(res)
}

func solve11b() {
	nums, _ := scanLineToSlice[int64]()

	var res int64
	for i := 0; i < len(nums); i++ {
		res += blink(make(map[int64][]int64), nums[i], 75, 75)
	}
	printToOut(res)
}
