package main

import (
	"regexp"
)

// regex to match the strings and capture the numbers
var mulRegexp = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func multiplyMemory(s string) int64 {
	matches := mulRegexp.FindAllStringSubmatch(s, -1)

	var res int64
	var x, y int64
	for _, m := range matches {
		if len(m) != 3 {
			panicf("Invalid number of captured groups for match %v", m)
		}

		// scan the number values (0th idx contains the full match)
		sscan(m[1], &x)
		sscan(m[2], &y)

		res += x * y
	}

	return res
}

func solve3a() {
	var s string
	var res int64

	for scan(&s) {
		res += multiplyMemory(s)
	}

	printToOut(res)
}

func solve3b() {
	var s string
	var res int64

	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	var enabled = true
	for scan(&s) {
		var from = 0
		var match []int
		for from < len(s) {
			if enabled {
				match = dontRe.FindStringIndex(s[from:])
			} else {
				match = doRe.FindStringIndex(s[from:])
			}

			var end = len(s)
			if match != nil {
				end = from + match[0]
			}

			if enabled {
				res += multiplyMemory(s[from:end])
			}

			if match == nil {
				break
			}

			from += match[1]
			enabled = !enabled
		}
	}

	printToOut(res)
}
