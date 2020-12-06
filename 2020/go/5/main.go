package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strings"
)

func rowSeatCalc(designator string, l, h int) int {
	m, res := 0, 0

	for _, d := range designator {
		if string(d) == "F" || string(d) == "L" {
			m = int(math.Floor(float64(h+l) / 2.0))
			res, h = l, m
		} else {
			m = int(math.Ceil(float64(h+l) / 2.0))
			l, res = m, h
		}
	}

	return res
}

func seatIds(designators []string) []int {
	var ids []int
	for _, designator := range designators {
		ids = append(ids, rowSeatCalc(designator[:7], 0, 127)*8+
			rowSeatCalc(designator[7:], 0, 7))
	}
	return ids
}

func intMax(x []int) int {
	var max int
	for _, i := range x {
		if i > max {
			max = i
		}
	}
	return max
}

func missingSeat(ids []int) int {
	var missing int
	sort.Ints(ids)
	for i := 0; i < len(ids)-2; i++ {
		if ids[i]+1 != ids[i+1] {
			return ids[i] + 1
		}
	}
	return missing
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("Error handling input")
	}

	passes := strings.Split(string(data), "\n")
	ids := seatIds(passes)

	fmt.Println("Part 1:", intMax(ids))
	fmt.Println("Part 2:", missingSeat(ids))

}
