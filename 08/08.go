package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func toIntArray(line string) []int {
	s := make([]int, len(line))
	for k, v := range line {
		s[k] = int(v - '0')
	}

	return s
}

func digits(layer []int, digit int) int {
	cnt := 0
	for _, v := range layer {
		if v == digit {
			cnt++
		}
	}

	return cnt
}

func layerWithFewest(layers map[int][]int, digit int) (int, int) {
	fewest := 0
	fewestCnt := math.MaxInt32
	for k, v := range layers {
		cnt := digits(v, digit)
		if cnt < fewestCnt {
			fewest = k
			fewestCnt = cnt
		}
	}

	return fewest, fewestCnt
}

func part1(width, height int, array []int) int {

	layers := make(map[int][]int)
	i := 0
	for (i * width * height) < len(array) {
		start := i * width * height
		end := start + (width * height)
		layer := array[start:end]

		layers[i + 1] = layer
		i++
	}

	layer, _ := layerWithFewest(layers, 0)
	return digits(layers[layer], 1) * digits(layers[layer], 2)
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()

		fmt.Println(part1(25, 6, toIntArray(line)))
	} else {
		panic("oh no")
	}
}
