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

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(data), "\n")

	if err != nil {
		log.Fatal("Error handling input")
	}

	fmt.Println("Part 1:", partOne(asInt(lines)))
}
