package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type space_object struct {
	name string
	orbits string
}

func build_orbits(orbits []string) map[string]space_object {
	m := make(map[string]space_object)
	m["COM"] = space_object{name: "COM"}

	for _, v := range orbits {
		split := strings.Split(v, ")")
		orbited := split[0]
		orbiting := split[1]

		o1, present := m[orbited]
		if !present {
			m[orbited] = space_object{name: orbited}
			o1 = m[orbited]
		}

		o2, present := m[orbiting]
		if !present {
			m[orbiting] = space_object{name: orbiting, orbits: o1.name}
		} else if o2.orbits == "" {
			m[orbiting] = space_object{name: orbiting, orbits: o1.name}
		}
	}

	return m
}

func part1(m map[string]space_object) int {
	cnt := 0
	for _, v := range m {
		next := v.orbits
		for next != "" {
			next = m[next].orbits
			cnt++
		}
	}

	return cnt
}

func part2(m map[string]space_object) int {
	san := m["SAN"]
	santocom := make(map[string]space_object)
	for san.orbits != "" {
		santocom[san.orbits] = san
		san = m[san.orbits]
	}

	transfers := 0
	firstCommon := m["YOU"]
	for _, present := santocom[firstCommon.orbits]; !present;  _, present = santocom[firstCommon.orbits] {
		transfers++
		firstCommon = m[firstCommon.orbits]
	}

	san = m[m["SAN"].orbits]
	for san.name != firstCommon.orbits {
		transfers++
		san = m[san.orbits]
	}

	return transfers

}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	array := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		array = append(array, line)
	}

	orbits := build_orbits(array)
	fmt.Println(part1(orbits))
	fmt.Println(part2(orbits))
}
