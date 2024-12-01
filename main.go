package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var reader *bufio.Reader
var writer = bufio.NewWriter(os.Stdout)

func scan(args ...any) bool {
	n, err := fmt.Fscan(reader, args...)
	if n == len(args) {
		return true
	}

	if errors.Is(err, io.EOF) {
		return false
	}

	panic(fmt.Sprintf("failed to scan: %v", err))
}

func printToOut(args ...interface{}) {
	_, err := fmt.Fprintln(writer, args...)
	if err != nil {
		panic(fmt.Sprintf("failed to print: %v", err))
	}
}

func panicf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}

var taskSolvers = map[string]func(){
	"1_a": solve1a,
	"1_b": solve1b,
}

func main() {
	defer func() {
		err := writer.Flush()
		if err != nil {
			panicf("failed to flush: %v", err)
		}
	}()

	if len(os.Args) != 2 {
		panicf("Usage: go run main.go DAY_VARIANT")
	}

	task := os.Args[1]
	taskSplits := strings.SplitN(task, "_", 2)
	if len(taskSplits) != 2 {
		panicf("Usage: go run main.go DAY_VARIANT")
	}

	filePath := filepath.Join("data", taskSplits[0])
	f, err := os.Open(filePath)
	if err != nil {
		panicf("failed to open file %s: %v", filePath, err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			panicf("failed to close file %s: %v", filePath, err)
		}
	}()

	reader = bufio.NewReader(f)

	solver, ok := taskSolvers[task]
	if !ok {
		panicf("solver for task %s not found", task)
	}

	printToOut("Running solver for task:", task)
	solver()
}
