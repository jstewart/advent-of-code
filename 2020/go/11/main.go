package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func deepCopy(src [][]string) [][]string {
	dst := make([][]string, len(src))
	for i, row := range src {
		nr := make([]string, len(row))
		copy(nr, row)
		dst[i] = nr
	}
	return dst
}

func deepEqual(a, b [][]string) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func getAdjacents(loc coordinate, read *[][]string) []coordinate {
	var adjacents []coordinate
	canMoveLeft, canMoveRight, canMoveUp, canMoveDown := loc.x > 0,
		loc.x < len((*read)[0])-1,
		loc.y > 0, loc.y < len((*read))-1

	if canMoveUp {
		adjacents = append(adjacents, coordinate{y: loc.y - 1, x: loc.x})
		if canMoveLeft {
			adjacents = append(adjacents, coordinate{y: loc.y - 1, x: loc.x - 1})
		}
		if canMoveRight {
			adjacents = append(adjacents, coordinate{y: loc.y - 1, x: loc.x + 1})
		}
	}

	if canMoveDown {
		adjacents = append(adjacents, coordinate{y: loc.y + 1, x: loc.x})
		if canMoveLeft {
			adjacents = append(adjacents, coordinate{y: loc.y + 1, x: loc.x - 1})
		}
		if canMoveRight {
			adjacents = append(adjacents, coordinate{y: loc.y + 1, x: loc.x + 1})
		}
	}

	if canMoveLeft {
		adjacents = append(adjacents, coordinate{y: loc.y, x: loc.x - 1})
	}
	if canMoveRight {
		adjacents = append(adjacents, coordinate{y: loc.y, x: loc.x + 1})
	}

	return adjacents
}

func markSeat(loc coordinate, read, write *[][]string) {
	/*  Rules:
	 *	If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
	 *	If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
	 * 	Otherwise, the seat's state does not change.
	 */
	square := (*read)[loc.y][loc.x]
	if square == "." {
		return
	}

	adjacents := getAdjacents(loc, read)
	var occupied int
	for _, x := range adjacents {
		if (*read)[x.y][x.x] == "#" {
			occupied++
		}
	}

	if square == "L" && occupied == 0 {
		(*write)[loc.y][loc.x] = "#"
	}
	if square != "." && occupied >= 4 {
		(*write)[loc.y][loc.x] = "L"
	}
}

func markSeat2(loc coordinate, read, write *[][]string) {
	/*  Rules:
	 * 5 or more VISIBLE occupied seats in all 8 directions
	 */
	square := (*read)[loc.y][loc.x]
	if square == "." {
		return
	}

	var occupied int
	// starting position
	// left, right
	// up, down
	// up/left
	// up/right
	// down/left
	// down/right
	// stop if hit L/# or edge
	// occupied++ if #
	// for _, x := range adjacents {
	// 	if (*read)[x.y][x.x] == "#" {
	// 		occupied++
	// 	}
	// }

	if square == "L" && occupied == 0 {
		(*write)[loc.y][loc.x] = "#"
	}
	if square != "." && occupied >= 5 {
		(*write)[loc.y][loc.x] = "L"
	}
}

func partOne(read, write *[][]string, changed bool) int {
	var occupied int

	for changed {
		for row := 0; row < len(*read); row++ {
			for col := 0; col < len((*read)[0]); col++ {
				markSeat(coordinate{x: col, y: row}, read, write)
			}
		}
		if deepEqual(*read, *write) {
			changed = false
		} else {
			tmp := deepCopy(*write)
			read = &tmp
		}
	}

	for _, row := range *read {
		for _, col := range row {
			if col == "#" {
				occupied++
			}
		}
	}
	return occupied
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(data), "\n")
	var read [][]string

	if err != nil {
		log.Fatal("Error handling input")
	}
	for _, x := range lines {
		read = append(read, strings.Split(x, ""))
	}
	write := deepCopy(read)
	fmt.Println("part 1:", partOne(&read, &write, true))
}
