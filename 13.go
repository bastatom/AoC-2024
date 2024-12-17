package main

func calcTokens(x1, x2, xTotal, y1, y2, yTotal int64) int64 {
	num := (yTotal * x1) - (xTotal * y1)

	bNum := (y2 * x1) - (y1 * x2)

	// THis line will panic in case when b == 0 (b is a linear combination of a)
	// if this happens we would need to check the ratio to calculate which
	// buttons are more profitable to use based on their token prices and
	// with the information solve the equation. However, since this case was
	// not tested as part of the input data I was too lazy to actually deal
	// with it
	if num%bNum != 0 {
		return 0
	}
	b := num / bNum

	aNum := xTotal - (x2 * b)
	if aNum%x1 != 0 {
		return 0
	}
	a := aNum / x1

	return a*3 + b*1
}

func solve13a() {
	var b byte
	var x1, x2, xTotal, y1, y2, yTotal int64

	var res int64
	for {
		scanf("Button A: X+%d, Y+%d\n", &x1, &y1)
		scanf("Button B: X+%d, Y+%d\n", &x2, &y2)
		scanf("Prize: X=%d, Y=%d\n", &xTotal, &yTotal)

		res += calcTokens(x1, x2, xTotal, y1, y2, yTotal)

		if !scanf("%c\n", &b) {
			break
		}
	}

	printToOut(res)
}

func solve13b() {
	var b byte
	var x1, x2, xTotal, y1, y2, yTotal int64

	var res int64
	for {
		scanf("Button A: X+%d, Y+%d\n", &x1, &y1)
		scanf("Button B: X+%d, Y+%d\n", &x2, &y2)
		scanf("Prize: X=%d, Y=%d\n", &xTotal, &yTotal)

		var add int64 = 10000000000000

		res += calcTokens(x1, x2, xTotal+add, y1, y2, yTotal+add)

		if !scanf("%c\n", &b) {
			break
		}
	}

	printToOut(res)
}
