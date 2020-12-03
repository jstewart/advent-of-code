package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func treesHit(xFactor, yFactor int, grid [][]string) int {
	trees, x, y := 0, 0, 0
	xBounds, yBounds := len(grid[0]), len(grid)-1

	for y < yBounds {
		x += xFactor
		y += yFactor

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

func calcSlopes(grid [][]string) int {
	ans := 1
	slopes := [][]int{
		[]int{1, 1},
		[]int{3, 1},
		[]int{5, 1},
		[]int{7, 1},
		[]int{1, 2},
	}

	for _, pair := range slopes {
		ans *= treesHit(pair[0], pair[1], grid)
	}

	return ans
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

	fmt.Println("Part 1:", treesHit(3, 1, grid))
	fmt.Println("Part 2:", calcSlopes(grid))
}
