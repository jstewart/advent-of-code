package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseBags(bags string) []string {
	var x []string
	return x
}

func prepareBagsGraph(bags []string) map[string][]string {
	m := make(map[string][]string)

	for _, bag := range bags {
		all := strings.Split(bag[:len(bag)-1], " bags contain ")
		parent, children := all[0], strings.Split(all[1], ",")
		for _, child := range children {
			if child == "no other bags" {
				continue
			}
			parts := strings.Split(strings.TrimSpace(child), " ")
			color := strings.Join(parts[1:len(parts)-1], " ")
			m[color] = append(m[color], parent)
		}
	}
	return m
}

func countParents(color string, m map[string][]string, acc map[string]bool) map[string]bool {
	parents := m[color]

	if len(parents) == 0 {
		return acc
	}

	for _, c := range parents {
		acc[c] = true
		acc = countParents(c, m, acc)
	}

	return acc
}

func partOne(graph map[string][]string) int {
	return len(countParents("shiny gold", graph, map[string]bool{}))
}

func main() {
	var input []string

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	graph := prepareBagsGraph(input)

	fmt.Println("Part 1:", partOne(graph))
}
