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
}
