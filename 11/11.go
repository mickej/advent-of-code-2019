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

type position struct {
	x int
	y int
	direction string
	color int
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()

		painted := outputs(toIntArray(strings.Split(line, ",")), 1)
		fmt.Println("part1", len(painted))

		largestX := 0
		smallestX := 0
		largestY := 0
		smallestY := 0
		for k, _ := range painted {
			split := strings.Split(k, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			if x > largestX {
				largestX = x
			}

			if x < smallestX {
				smallestX = x
			}

			if y > largestY {
				largestY = y
			}

			if y < smallestY {
				smallestY = y
			}
		}

		for j := smallestY; j <= largestY; j++ {
			for i := smallestX; i <= largestX; i++ {
				k := strconv.Itoa(i) + "," + strconv.Itoa(j)
				color, present := painted[k]
				if !present {
					color = 0
				}

				if color == 0 {
					fmt.Print(" ")
				} else if color == 1 {
					fmt.Print("█")
				} else {
					panic("crap...")
				}
			}

			fmt.Println()
		}
	} else {
		panic("oh no")
	}
}

func paint(p position, color int, painted map[string]int) position {
	paintedKey := strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)

	_, present := painted[paintedKey]
	if !present {
		painted[paintedKey] = 0
	}

	if color == 0 {
		painted[paintedKey] = 0
	} else if color == 1 {
		painted[paintedKey] = 1
	}

	return position{p.x, p.y, p.direction, color}
}

func move(turn int, p position) position {
	direction := p.direction
	if turn == 0 {
		if direction == "^" {
			direction = "<"
		} else if direction == "<" {
			direction = "v"
		} else if direction == "v" {
			direction = ">"
		} else if direction == ">" {
			direction = "^"
		}
	} else if turn == 1 {
		if direction == "^" {
			direction = ">"
		} else if direction == ">" {
			direction = "v"
		} else if direction == "v" {
			direction = "<"
		} else if direction == "<" {
			direction = "^"
		}
	}

	x := p.x
	y := p.y
	if direction == "<" {
		x += 1
	} else if direction == "^" {
		y += 1
	} else if direction == ">" {
		x -= 1
	} else if direction == "v" {
		y -= 1
	}

	return position{x, y, direction, p.color}
}

func makeLargerIfNeeded(codes []int, idx int) []int {
	ret := codes
	if idx >= len(codes) {
		ret = make([]int, idx + 1)
		copy(ret, codes)
	}

	return ret;
}

func read(codes []int, mode, val, relativeBase int) int {
	if mode == 0 {
		ret := makeLargerIfNeeded(codes, val)
		return ret[val]
	} else if mode == 2 {
		ret := makeLargerIfNeeded(codes, relativeBase + val)
		return ret[relativeBase + val]
	}

	return val
}

func write(codes []int, idx, val int) []int {
	ret := makeLargerIfNeeded(codes, idx)
	ret[idx] = val
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

func outputs(codes []int, input int) map[string]int {
	painted := make(map[string]int)
	currentPosition := position{0, 0, "^", input}

	output := make([]int, 0)
	relativeBase := 0

	for i := 0; i <= len(codes); {
		abcde := codes[i]
		opcode := abcde % 100
		mode1 := (abcde % 1000) / 100
		mode2 := (abcde % 10000) / 1000
		mode3 := (abcde % 100000) / 10000

		if opcode == 99 {
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
			if len(output) == 2 {
				currentPosition = paint(currentPosition, output[0], painted)
				currentPosition = move(output[1], currentPosition)

				paintedKey := strconv.Itoa(currentPosition.x) + "," + strconv.Itoa(currentPosition.y)
				v, present := painted[paintedKey]
				if !present {
					input = 0
				} else {
					input = v
				}

				output = make([]int, 0)
			}

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
		}
	}

	return painted
}

