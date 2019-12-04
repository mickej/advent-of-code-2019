package main

import "fmt"

func meetsCriteria(number int) bool {
	if number >= 100000 && number <= 999999 {
		first := (number % 1000000) / 100000
		second := (number % 100000) / 10000
		third := (number % 10000) / 1000
		fourth := (number % 1000) / 100
		fifth := (number % 100) / 10
		sixth := number % 10
		if first <= second && second <= third && third <= fourth && fourth <= fifth && fifth <= sixth {
			return first == second || second == third || third == fourth || fourth == fifth || fifth == sixth
		}
	}

	return false
}

func part1(from, to int) []int {
	ret := make([]int, 0)
	for from <= to {
		if meetsCriteria(from) {
			ret = append(ret, from)
		}
		from++
	}

	return ret
}

func main() {
	from := 357253
	to   := 892942

	fmt.Println(len(part1(from, to)))
}
