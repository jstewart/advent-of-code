package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type waypoint struct {
	orientation string   // direction ship is facing
	directions  []string // N S E W
	x           int      // east-west
	y           int      // north-south
}

type ship struct {
	orientation string   // direction ship is facing
	directions  []string // N S E W
	x           int      // east-west
	y           int      // north-south
	waypoint    *waypoint
}

func newShip(waypoint *waypoint) *ship {
	return &ship{orientation: "E",
		directions: []string{"N", "E", "S", "W"},
		waypoint:   waypoint}
}

func newWaypoint() *waypoint {
	return &waypoint{orientation: "E",
		directions: []string{"N", "E", "S", "W"},
		x:          10,
		y:          1,
	}
}

func (s *ship) turn(direction string) {
	units, _ := strconv.Atoi(direction[1:])
	moves, idx := units/90, 0
	for i, x := range s.directions {
		if s.orientation == x {
			idx = i
		}
	}

	for moves > 0 {
		if direction[:1] == "R" {
			if idx < 3 {
				idx++
			} else {
				idx = 0
			}
		} else {
			if idx > 0 {
				idx--
			} else {
				idx = 3
			}
		}
		moves--
	}
	s.orientation = s.directions[idx]
}

func (w *waypoint) turn(direction string) {
	// 10 e 4 n
	// rotate 90R
	// 4 e 10n
	units, _ := strconv.Atoi(direction[1:])
	moves, idx := units/90, 0
	for i, x := range w.directions {
		if w.orientation == x {
			idx = i
		}
	}

	for moves > 0 {
		if direction[:1] == "R" {
			tmp := w.y
			w.y = 0 - w.x
			w.x = tmp
			if idx < 3 {
				idx++
			} else {
				idx = 0
			}
		} else {
			tmp := w.x
			w.x = 0 - w.y
			w.y = tmp
			if idx > 0 {
				idx--
			} else {
				idx = 3
			}
		}
		moves--
	}
	w.orientation = w.directions[idx]
}

func (s *ship) move(direction string) {
	units, _ := strconv.Atoi(direction[1:])
	switch direction[:1] {
	case "F":
		s.move(fmt.Sprintf("%s%d", s.orientation, units))
	case "L":
		s.turn(direction)
	case "R":
		s.turn(direction)
	case "N":
		s.y += units
	case "S":
		s.y -= units
	case "E":
		s.x += units
	case "W":
		s.x -= units
	}

}

func (s *ship) move2(direction string) {
	fmt.Println(direction)
	units, _ := strconv.Atoi(direction[1:])
	switch direction[:1] {
	case "F":
		s.x += (s.waypoint.x * units)
		s.y += (s.waypoint.y * units)
	case "L":
		s.waypoint.turn(direction)
	case "R":
		s.waypoint.turn(direction)
	case "N":
		s.waypoint.y += units
	case "S":
		s.waypoint.y -= units
	case "E":
		s.waypoint.x += units
	case "W":
		s.waypoint.x -= units
	}
	fmt.Println(s, s.waypoint)

}

func partOne(s *ship, movements []string) int {
	for _, m := range movements {
		s.move(m)
	}

	return int(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}

func partTwo(s *ship, movements []string) int {
	for _, m := range movements {
		s.move2(m)
	}

	return int(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(data), "\n")
	if err != nil {
		log.Fatal("Error handling input")
	}

	s := newShip(nil)
	s2 := newShip(newWaypoint())

	fmt.Println("part 1:", partOne(s, lines))
	fmt.Println("part 2:", partTwo(s2, lines))
}
