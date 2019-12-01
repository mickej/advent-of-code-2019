package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total = 0
	for scanner.Scan() {
		mass, _ := strconv.ParseFloat(scanner.Text(), 64)
		fuel := (int(mass) / 3) - 2
		total += fuel
	}

	fmt.Printf(strconv.FormatInt(int64(total), 10))
}
