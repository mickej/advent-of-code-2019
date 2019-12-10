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

		//fmt.Println(part1(toIntArray(strings.Split(line, ",")), 5))
		//fmt.Println(part2(toIntArray(strings.Split(line, ","))))
	} else {
		panic("oh no")
	}
}

func read(codes []int, mode, val, relativeBase int) int {
	if mode == 0 {
		return codes[val]
	} else if mode == 2 {
		return codes[relativeBase + val]
	}

	return val
}

func write(codes []int, idx, val int) []int {
	ret := codes
	if idx >= len(codes) {
		fmt.Println("make larger", idx)
		ret = make([]int, idx + 1)
		copy(ret, codes)
	}

	ret[idx] = val
	fmt.Println("wrote", idx, val)

	return ret
}

func getIdx(codes []int, mode, idx, relativeBase int) int {
	if mode == 0 {
		return codes[idx]
	} else if mode == 2 {
		return codes[idx] + relativeBase
	} else {
		panic("oh no")
	}
}

func part1(codes []int, input int) []int {
	output := make([]int, 0)
	relativeBase := 0

	for i := 0; i <= len(codes); {
		abcde := codes[i]
		opcode := abcde % 100
		mode1 := (abcde % 1000) / 100
		mode2 := (abcde % 10000) / 1000
		mode3 := (abcde % 100000) / 10000

		if opcode == 99 {
			fmt.Println("99")
			break
		} else if opcode == 1 {
			val1 := read(codes, mode1, codes[i + 1], relativeBase)
			val2 := read(codes, mode2, codes[i + 2], relativeBase)
			idx := getIdx(codes, mode3, i + 3, relativeBase)
			codes = write(codes, idx, val1 + val2)

			i += 4
		} else if opcode == 2 {
			val1 := read(codes, mode1, codes[i+1], relativeBase)
			val2 := read(codes, mode2, codes[i+2], relativeBase)
			idx := getIdx(codes, mode3, i + 3, relativeBase)
			codes = write(codes, idx, val1 * val2)

			i += 4
		} else if opcode == 3 {
			idx := getIdx(codes, mode1, i + 1, relativeBase)
			codes = write(codes, idx, input)
			i += 2
		} else if opcode == 4 {
			val1 := read(codes, mode1, codes[i + 1], relativeBase)
			output = append(output, val1)

			i += 2
		} else if opcode == 5 {
			if read(codes, mode1, codes[i+1], relativeBase) != 0 {
				i = read(codes, mode2, codes[i+2], relativeBase)
			} else {
				i += 3
			}
		} else if opcode == 6 {
			if read(codes, mode1, codes[i+1], relativeBase) == 0 {
				i = read(codes, mode2, codes[i+2], relativeBase)
			} else {
				i += 3
			}
		} else if opcode == 7 {
			val1 := read(codes, mode1, codes[i+1], relativeBase)
			val2 := read(codes, mode2, codes[i+2], relativeBase)
			idx := getIdx(codes, mode3, i + 3, relativeBase)
			if val1 < val2 {
				codes = write(codes, idx, 1)
			} else {
				codes = write(codes, idx, 0)
			}
			i += 4
		} else if opcode == 8 {
			val1 := read(codes, mode1, codes[i+1], relativeBase)
			val2 := read(codes, mode2, codes[i+2], relativeBase)
			idx := getIdx(codes, mode3, i + 3, relativeBase)
			if val1 == val2 {
				codes = write(codes, idx, 1)
			} else {
				codes = write(codes, idx, 0)
			}
			i += 4
		} else if opcode == 9 {
			relativeBase += read(codes, mode1, codes[i + 1], relativeBase)
			i += 2
		} else {
			fmt.Println("i don't know that")
		}
	}

	return output
}

