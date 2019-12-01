package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calcFuel(mass int) int {
	return (mass / 3) - 2
}

func calcModule(mass int) int {
	if mass <= 0 {
		return 0
	}

	return mass + calcModule(calcFuel(mass))
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalPart1 = 0
	var totalPart2 = 0
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		totalPart1 += calcFuel(i)
		totalPart2 += calcModule(calcFuel(i))
	}

	fmt.Printf("Part1: %d\n", totalPart1)
	fmt.Printf("Part2: %d\n", totalPart2)
}
