package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func mustAtoi(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func extractRanges(spec string) [][]int {
	var ranges [][]int
	re := regexp.MustCompile(`(\d+-\d+)`)
	matches := re.FindAllString(spec, -1)

	for _, match := range matches {
		rge := strings.Split(match, "-")
		ranges = append(ranges, []int{mustAtoi(rge[0]), mustAtoi(rge[1])})
	}
	return ranges
}

func invalidNumbers(ticket string, ranges [][]int) []int {
	var invalid []int
	ticketNums := strings.Split(ticket, ",")

	for _, num := range ticketNums {
		numValid := false
		subject := mustAtoi(num)

		for _, rge := range ranges {
			if subject >= rge[0] && subject <= rge[1] {
				numValid = true
			}
		}

		if !numValid {
			invalid = append(invalid, subject)
		}
	}

	return invalid
}

func partOne(tickets string, ranges [][]int) int {
	var sum int

	for _, ticket := range strings.Split(tickets, "\n") {
		for _, invalid := range invalidNumbers(ticket, ranges) {
			sum += invalid
		}
	}

	return sum
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("Error handling input")
	}
	sections := strings.Split(string(data), "\n\n")
	spec, tickets := sections[0], sections[2]
	ranges := extractRanges(spec)

	fmt.Println("part 1:", partOne(tickets, ranges))
}
