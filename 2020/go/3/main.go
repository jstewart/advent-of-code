package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func treesHit(grid [][]string) int {
	trees, x, y := 0, 0, 0
	xBounds, yBounds := len(grid[0]), len(grid)-1

	for y < yBounds {
		x += 3
		y++

		// Simulates a never ending edge
		if x >= xBounds {
			x = x - xBounds
		}

		if grid[y][x] == "#" {
			trees++
		}
	}

	return trees
}

func main() {
	var grid [][]string

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Trees hit: ", treesHit(grid))
	fmt.Println("grid test", treesHit(grid))
}
