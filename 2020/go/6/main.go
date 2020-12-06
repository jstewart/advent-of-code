package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func processGroup(blob string) (int, map[string]int) {
	m := make(map[string]int)
	answerLists := strings.Split(blob, "\n")
	answerListTotal := len(answerLists)

	for _, answers := range answerLists {
		for _, answer := range answers {
			m[string(answer)]++
		}
	}

	return answerListTotal, m
}

func totalSum(blob string, all bool) int {
	var total int
	groups := strings.Split(blob, "\n\n")

	for _, group := range groups {
		answerListTotal, m := processGroup(group)

		if all {
			for _, count := range m {
				if count == answerListTotal {
					total++
				}
			}
		} else {
			total += len(m)

		}
	}

	return total
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		log.Fatal("Error handling input")
	}

	fmt.Println("Part 1:", totalSum(string(data), false))
	fmt.Println("Part 2:", totalSum(string(data), true))
}
