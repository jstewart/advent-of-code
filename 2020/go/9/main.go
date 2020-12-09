package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func asInt(strings []string) []int {
	var ints []int
	for _, s := range strings {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return ints
}

func isTwoSum(tgt int, window []int) bool {
	addendCache := make(map[int]int)

	for i, x := range window {
		diff := tgt - x

		if _, ok := addendCache[diff]; ok {
			return true
		}
		addendCache[x] = i
	}
	return false
}

func partOne(data []int) int {
	for i := 25; i < len(data); i++ {
		windowStart := i - 25
		if !isTwoSum(data[i], data[windowStart:i]) {
			return data[i]
		}

	}
	return 0
}

func partTwo(target int, data []int) int {
	for i := 0; i < len(data)-2; i++ {
		sum, sm, lg := data[i], data[i], data[i]
		for j := i + 1; j < len(data); j++ {
			if data[j] < sm {
				sm = data[j]
			}
			if data[j] > lg {
				lg = data[j]
			}

			sum += data[j]
			if sum == target {
				return sm + lg

			}
		}
	}
	return 0
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	lines := asInt(strings.Split(string(data), "\n"))

	if err != nil {
		log.Fatal("Error handling input")
	}

	p1 := partOne(lines)
	p2 := partTwo(p1, lines)

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
