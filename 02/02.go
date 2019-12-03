package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleOpcode(opcode, v1, v2 int) int {
	if opcode == 1 {
		return v1 + v2
	} else if opcode == 2 {
		return v1 * v2
	}

	panic("oh no, i dont know that opcode")
}

func toIntArray(array []string) []int {
	s := make([]int, len(array))
	for k, v := range array {
		i, _ := strconv.Atoi(v)
		s[k] = i
	}

	return s
}

func batches(array []int) [][]int {
	batchSize := 4
	var batches [][]int

	for batchSize < len(array) {
		array, batches = array[batchSize:], append(batches, array[0:batchSize:batchSize])
	}
	return append(batches, array)
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		split := toIntArray(strings.Split(line, ","))
		split[1] = 12
		split[2] = 2

		part1(split)
	} else {
		panic("oh no")
	}
}

func part1(codes []int) {
	for _, batch := range batches(codes) {
		opcode := batch[0]
		if opcode == 99 {
			fmt.Println("DONE", codes[0])
			break
		} else {
			codes[batch[3]] = handleOpcode(opcode, codes[batch[1]], codes[batch[2]])
		}
	}
}
