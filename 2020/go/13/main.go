package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func earliestDeparture(ts int, busIds []int) int {
	counter := ts

	for {
		for _, id := range busIds {
			if counter%id == 0 {
				return id * (counter - ts)
			}
		}
		counter++
	}
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(data), "\n")
	if err != nil {
		log.Fatal("Error handling input")
	}

	earliest, _ := strconv.Atoi(lines[0])
	var busIds []int
	for _, id := range strings.Split(lines[1], ",") {
		if id != "x" {
			bus, _ := strconv.Atoi(id)
			busIds = append(busIds, bus)
		}
	}

	fmt.Println("part 1:", earliestDeparture(earliest, busIds))
}
