package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func earliestDeparture(ts int, busIds map[int]int) int {
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

func departuresMatch(busses map[int]int, lowest int) int64 {
	var ts int64

	for {
		allMatch := true

		ts += int64(lowest)
		for offset, busID := range busses {
			if (ts+int64(offset))%int64(busID) != 0 {
				allMatch = false
				break
			}
		}

		if allMatch {
			break
		}

	}
	return ts
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(data), "\n")
	if err != nil {
		log.Fatal("Error handling input")
	}

	earliest, _ := strconv.Atoi(lines[0])
	busIds, lowest := make(map[int]int), 99999999
	for idx, id := range strings.Split(lines[1], ",") {
		if id != "x" {
			bus, _ := strconv.Atoi(id)
			busIds[idx] = bus
			if bus < lowest {
				lowest = bus
			}
		}
	}

	fmt.Println("part 1:", earliestDeparture(earliest, busIds))
	fmt.Println("part 2:", departuresMatch(busIds, lowest))
}
