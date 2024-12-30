package main

import "strings"

func towelDesignPossible(mem map[string]bool, design string, patterns []string) bool {
	if len(design) == 0 {
		return true
	}

	possible, ok := mem[design]
	if ok {
		return possible
	}

	for _, pattern := range patterns {
		if len(pattern) <= len(design) && design[:len(pattern)] == pattern {
			if towelDesignPossible(mem, design[len(pattern):], patterns) {
				mem[design] = true
				return mem[design]
			}
		}
	}

	mem[design] = false
	return mem[design]
}

func towelDesignPossibilities(mem map[string]int64, design string, patterns []string) int64 {
	if len(design) == 0 {
		return 1
	}

	possibilities, ok := mem[design]
	if ok {
		return possibilities
	}

	for _, pattern := range patterns {
		if len(pattern) <= len(design) && design[:len(pattern)] == pattern {
			possibilities += towelDesignPossibilities(mem, design[len(pattern):], patterns)
		}
	}

	mem[design] = possibilities
	return mem[design]

}

func solve19a() {
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, " ", "")
	splits := strings.Split(s, ",")

	available := make([]string, 0, len(splits))
	for _, split := range splits {
		available = append(available, split)
	}

	var res int64
	memMap := make(map[string]bool)
	for scan(&s) {
		if towelDesignPossible(memMap, s, available) {
			res++
		}
	}

	printToOut(res)
}

func solve19b() {
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, " ", "")
	splits := strings.Split(s, ",")

	available := make([]string, 0, len(splits))
	for _, split := range splits {
		available = append(available, split)
	}

	var res int64
	memMap := make(map[string]int64)
	for scan(&s) {
		res += towelDesignPossibilities(memMap, s, available)
	}

	printToOut(res)
}
