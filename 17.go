package main

import (
	"math"
	"strconv"
	"strings"
)

func operandToCombo(operand int64, reg [3]int64) int64 {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return reg[0]
	case 5:
		return reg[1]
	case 6:
		return reg[2]
	default:
		panicf("invalid combo operand %d", operand)
		return -1
	}
}

func divByPowerOf2(a, pow int64) int64 {
	if pow >= 63 {
		panicf("power is too big %d, it has to be <= 63 to fit into int64", pow)
	}
	return a / (1 << pow)
}

func performProgram(ins []int64, reg [3]int64) []int64 {
	var out []int64
	var iPointer int64 = 0
	for int(iPointer) < len(ins) {
		instruction := ins[iPointer]
		operand := ins[iPointer+1]
		switch instruction {
		case 0:
			reg[0] = divByPowerOf2(reg[0], operandToCombo(operand, reg))
		case 1:
			reg[1] = reg[1] ^ operand
		case 2:
			reg[1] = operandToCombo(operand, reg) % 8
		case 3:
			if reg[0] != 0 {
				iPointer = operand
				continue
			}
		case 4:
			reg[1] = reg[1] ^ reg[2]
		case 5:
			out = append(out, operandToCombo(operand, reg)%8)
		case 6:
			reg[1] = divByPowerOf2(reg[0], operandToCombo(operand, reg))
		case 7:
			reg[2] = divByPowerOf2(reg[0], operandToCombo(operand, reg))
		}
		iPointer += 2
	}
	return out
}

func solve17a() {
	var b byte
	var s string
	var reg [3]int64
	var ins []int64
	scanf("Register A: %d\n", &reg[0])
	scanf("Register B: %d\n", &reg[1])
	scanf("Register C: %d\n", &reg[2])
	scanf("%c\n", &b)
	scanf("Program: %v\n", &s)

	splits := strings.Split(s, ",")
	for _, split := range splits {
		in, err := strconv.Atoi(split)
		if err != nil {
			panicf("Failed to parse instruction %s: %v\n", split, err)
		}
		ins = append(ins, int64(in))
	}

	out := performProgram(ins, reg)

	sb := strings.Builder{}
	sb.Grow(len(out)*2 - 1)
	for i := range out {
		if i != 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.FormatInt(out[i], 10))
	}
	printToOut(sb.String())
}

func findMatch(ins []int64, reg [3]int64, found int) int64 {
	for i := reg[0]; i < reg[0]+8; i++ {
		out := performProgram(ins, [3]int64{i, reg[1], reg[2]})
		if out[0] != ins[len(ins)-1-found] {
			continue
		}

		if found+1 == len(ins) {
			return i
		}

		if val := findMatch(ins, [3]int64{i * 8, reg[1], reg[2]}, found+1); val != math.MaxInt64 {
			return val
		}
	}

	return math.MaxInt64
}

func solve17b() {
	var b byte
	var s string
	var reg [3]int64
	var ins []int64
	scanf("Register A: %d\n", &reg[0])
	scanf("Register B: %d\n", &reg[1])
	scanf("Register C: %d\n", &reg[2])
	scanf("%c\n", &b)
	scanf("Program: %v\n", &s)

	splits := strings.Split(s, ",")
	for _, split := range splits {
		in, err := strconv.Atoi(split)
		if err != nil {
			panicf("Failed to parse instruction %s: %v\n", split, err)
		}
		ins = append(ins, int64(in))
	}

	printToOut(findMatch(ins, [3]int64{0, reg[1], reg[2]}, 0))
}
