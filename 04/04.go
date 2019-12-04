package main

import "fmt"

func numberAsArray(number int) []int {
	first := (number % 1000000) / 100000
	second := (number % 100000) / 10000
	third := (number % 10000) / 1000
	fourth := (number % 1000) / 100
	fifth := (number % 100) / 10
	sixth := number % 10

	return []int {first, second, third, fourth, fifth, sixth}
}

func meetsCriteria(array []int) bool {
	if array[0] <= array[1] && array[1] <= array[2] && array[2] <= array[3] && array[3] <= array[4] && array[4] <= array[5] {
		two, more := adjacents(duplicates(array))
		return two || more
	}

	return false
}

func part1(from, to int) []int {
	ret := make([]int, 0)
	for from <= to {
		if meetsCriteria(numberAsArray(from)) {
			ret = append(ret, from)
		}
		from++
	}

	return ret
}

func duplicates(array []int) map[int]int {
	m := make(map[int]int)
	for _, v := range array {
		_, present := m[v]
		if !present {
			m[v] = 1
		} else {
			m[v] = m[v] + 1
		}
	}

	return m
}

func adjacents(m map[int]int) (bool, bool) {
	twoAdjacent := false
	more := false
	for _, v := range m {
		if v == 2 {
			twoAdjacent = true
		} else if v > 2 {
			more = true
		}
	}

	return twoAdjacent, more
}

func part2(from, to int) []int {
	ret := make([]int, 0)
	for from <= to {
		array := numberAsArray(from)
		if meetsCriteria(array) {
			two, _ := adjacents(duplicates(array))
			if two {
				ret = append(ret, from)
			}
		}
		from++
	}

	return ret
}

func main() {
	from := 357253
	to   := 892942

	fmt.Println("Part1", len(part1(from, to)))
	fmt.Println("Part2", len(part2(from, to)))
}
