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

		fmt.Println(part1(toIntArray(strings.Split(line, ",")), 12, 2)[0])
		fmt.Println(part2(toIntArray(strings.Split(line, ","))))
	} else {
		panic("oh no")
	}
}

func part1(codes []int, noun, verb int) []int {
	codes[1] = noun
	codes[2] = verb

	for _, batch := range batches(codes) {
		opcode := batch[0]
		if opcode == 99 {
			break
		} else {
			codes[batch[3]] = handleOpcode(opcode, codes[batch[1]], codes[batch[2]])
		}
	}

	return codes
}

func part2(oc []int) int {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			codes := make([]int, len(oc))
			copy(codes, oc)

			codes = part1(codes, noun, verb)

			if codes[0] == 19690720 {
				return 100 * noun + verb
			}
		}
	}

	panic("Should have found something")
}
