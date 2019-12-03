package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type directiondistance struct {
	direction string
	distance  int
}

type point struct {
	x int
	y int
	steps int
}

func toDirectiondistance(line string) ([]directiondistance, int) {
	l := strings.Split(line, ",")
	var ret []directiondistance
	var highest = 0
	for _, v := range l {
		i, _ := strconv.Atoi(v[1:])
		ret = append(ret, directiondistance{direction: string(v[0]), distance: i})

		if i > highest {
			highest = i
		}
	}

	return ret, highest
}

func line(row []directiondistance) []point {
	x := 0
	y := 0
	steps := 1
	points := make([]point, 0)
	for _, v := range row {
		distance := v.distance
		for distance > 0 {
			if v.direction == "R" {
				x = x + 1
			} else if v.direction == "L" {
				x = x - 1
			} else if v.direction == "U" {
				y = y + 1
			} else if v.direction == "D" {
				y = y - 1
			} else {
				panic("oh crap")
			}

			points = append(points, point{x: x, y: y, steps: steps})

			distance--
			steps++
		}
	}

	return points
}

func duplicates(l1, l2 []point) []point {
	dups := make([]point, 0)
	for _, v1 := range l1 {
		for _, v2 := range l2 {
			if v1.y == v2.y && v1.x == v2.x {
				dups = append(dups, v1)
			}
		}
	}

	return dups
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(a, b point) int {
	return Abs(a.x - b.x) + Abs(a.y - b.y)
}

func part1(row1, row2 []directiondistance) int {
	line1 := line(row1)
	line2 := line(row2)

	dups := duplicates(line1, line2)
	min_distance := math.MaxInt32
	for _, v := range dups {
		d := distance(point{x: 0, y: 0}, v)
		if d < min_distance {
			min_distance = d
		}
	}

	return min_distance
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	row1, _ := toDirectiondistance(scanner.Text())
	scanner.Scan()
	row2, _ := toDirectiondistance(scanner.Text())
	fmt.Println(part1(row1, row2))
}
