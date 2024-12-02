package main

func sign(a int64) int64 {
	if a < 0 {
		return -1
	}
	return 1
}

func safe2a(line []int64) bool {
	sgn := sign(line[len(line)-1] - line[0])
	for i := 1; i < len(line); i++ {
		diff := line[i] - line[i-1]
		if abs(diff) < 1 || abs(diff) > 3 || sign(diff) != sgn {
			return false
		}
	}
	return true
}

func solve2a() {
	res := 0
	for {
		report, ok := scanLineToSlice[int64]()
		if !ok || len(report) == 0 {
			break
		}

		if safe2a(report) {
			res++
		}
	}
	printToOut(res)
}

func newSlice[T any](slices ...[]T) []T {
	var s []T
	for _, slice := range slices {
		s = append(s, slice...)
	}
	return s
}

func safe2b(line []int64) bool {
	sgn := sign(line[len(line)-1] - line[0])
	for i := 1; i < len(line); i++ {
		diff := line[i] - line[i-1]
		if abs(diff) < 1 || abs(diff) > 3 || sign(diff) != sgn {
			// once we find a 'faulty' index we check if omitting either of the
			// conflicting elements would solve the issue. However, we also need
			// to check for first/last element because in that case we could have
			// gotten false alarm because of the poorly calculated sgn
			// (omitting the first element should not really be necessary
			// since it is covered by the previous 2 cases, but it is kept here
			// for better clarity the check for the final element, however, needs
			// to be done to also correctly solve for cases such as [7 6 5 10])
			return safe2a(newSlice(line[:i], line[i+1:])) ||
				safe2a(newSlice(line[:i-1], line[i:])) ||
				safe2a(newSlice(line[1:])) ||
				safe2a(newSlice(line[:len(line)-1]))
		}
	}

	return true
}

func solve2b() {
	res := 0
	for {
		report, ok := scanLineToSlice[int64]()
		if !ok || len(report) == 0 {
			break
		}

		if safe2b(report) {
			res++
		}
	}
	printToOut(res)
}
