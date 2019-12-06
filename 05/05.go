package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toIntArray(array []string) []int {
	s := make([]int, len(array))
	for k, v := range array {
		i, _ := strconv.Atoi(v)
		s[k] = i
	}

	return s
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()

		fmt.Println(part1(toIntArray(strings.Split(line, ",")), 1))
		fmt.Println(part1(toIntArray(strings.Split(line, ",")), 5))
		//fmt.Println(part2(toIntArray(strings.Split(line, ","))))
	} else {
		panic("oh no")
	}
}

func value(codes []int, mode, val int) int {
	if mode == 0 {
		return codes[val]
	}

	return val
}

func part1(codes []int, input int) int {
	output := 0

	for i := 0; i <= len(codes); {
		abcde := codes[i]
		opcode := abcde % 100
		mode1 := (abcde % 1000) / 100
		mode2 := (abcde % 10000) / 1000

		if opcode == 99 {
			break
		} else if opcode == 4 {
			val1 := value(codes, mode1, codes[i + 1])
			output = val1

			i += 2
		} else if opcode == 3 {
			codes[codes[i + 1]] = input
			i += 2
		} else if opcode == 1 {
			val1 := value(codes, mode1, codes[i + 1])
			val2 := value(codes, mode2, codes[i + 2])
			codes[codes[i + 3]] = val1 + val2

			i += 4
		} else if opcode == 2 {
			val1 := value(codes, mode1, codes[i+1])
			val2 := value(codes, mode2, codes[i+2])
			codes[codes[i+3]] = val1 * val2

			i += 4
		} else if opcode == 5 {
			if value(codes, mode1, codes[i+1]) != 0 {
				i = value(codes, mode2, codes[i+2])
			} else {
				i += 3
			}
		} else if opcode == 6 {
			if value(codes, mode1, codes[i+1]) == 0 {
				i = value(codes, mode2, codes[i+2])
			} else {
				i += 3
			}
		} else if opcode == 7 {
			if value(codes, mode1, codes[i+1]) < value(codes, mode2, codes[i+2]) {
				codes[codes[i + 3]] = 1
			} else {
				codes[codes[i + 3]] = 0
			}
			i += 4
		} else if opcode == 8 {
			if value(codes, mode1, codes[i+1]) == value(codes, mode2, codes[i+2]) {
				codes[codes[i + 3]] = 1
			} else {
				codes[codes[i + 3]] = 0
			}
			i += 4
		}
	}

	return output
}

