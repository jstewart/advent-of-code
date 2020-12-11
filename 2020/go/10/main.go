package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func partOne(input []int) int {
	oneJolt, threeJolt := 1, 1
	sort.Ints(input)

	for i := 0; i < len(input)-1; i++ {
		diff := input[i+1] - input[i]
		if diff == 1 {
			oneJolt++
		}
		if diff == 3 {
			threeJolt++
		}
	}

	return (oneJolt * threeJolt)
}

func permutations(val int, set map[int]bool, cache map[int]int) int {
	if cached, ok := cache[val]; ok {
		return cached
	}
	if val == 0 {
		cache[val] = 1
		return 1
	}
	if val < 0 {
		cache[val] = 0
		return 0
	}
	if _, ok := set[val]; !ok {
		cache[val] = 0
		return 0
	}
	cache[val] = permutations(val-1, set, cache) +
		permutations(val-2, set, cache) +
		permutations(val-3, set, cache)

	return cache[val]
}

func partTwo(input []int) int {
	var max int
	set := make(map[int]bool)

	for _, x := range input {
		if x > max {
			max = x
		}
		set[x] = true
	}

	// Add terminal and device
	max += 3
	set[0] = true
	set[max] = true

	return permutations(max, set, make(map[int]int))
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(data), "\n")
	var input []int

	if err != nil {
		log.Fatal("Error handling input")
	}

	for _, x := range lines {
		i, _ := strconv.Atoi(x)
		input = append(input, i)
	}
	fmt.Println("part 1:", partOne(input))
	fmt.Println("part 2:", partTwo(input))
}
