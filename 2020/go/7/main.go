package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

type node struct {
	count int
	name  string
}

func preparePart2Graph(bags []string) map[string][]node {
	m := make(map[string][]node)

	for _, bag := range bags {
		all := strings.Split(bag[:len(bag)-1], " bags contain ")
		parent, children := all[0], strings.Split(all[1], ",")
		var nodes []node

		for _, child := range children {
			if child == "no other bags" {
				break
			}

			parts := strings.Split(strings.TrimSpace(child), " ")
			count, _ := strconv.Atoi(parts[0])
			color := strings.Join(parts[1:len(parts)-1], " ")
			nodes = append(nodes, node{count: count, name: color})
		}
		m[parent] = nodes
	}
	return m
}

func bagCountTraversal(color string, graph map[string][]node) int {
	count := 1
	edges := graph[color]

	if len(edges) == 0 {
		return count
	}

	for _, edge := range edges {
		count += (edge.count * bagCountTraversal(edge.name, graph))
	}

	return count
}

func partTwo(graph map[string][]node) int {
	return bagCountTraversal("shiny gold", graph) - 1
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
	graph2 := preparePart2Graph(input)
	fmt.Println("Part 2:", partTwo(graph2))
	fmt.Println("Part 1:", partOne(graph))
}
