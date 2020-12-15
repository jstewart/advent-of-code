package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func memAddr(line string) (*int, *uint64) {
	r, _ := regexp.Compile("^mem\\[(\\d+)\\]\\s=\\s(\\d+)$")
	match := r.FindStringSubmatch(line)
	if len(match) == 3 {
		addr, _ := strconv.Atoi(match[1])
		val, _ := strconv.ParseUint(match[2], 10, 64)

		return &addr, &val
	}
	return nil, nil
}

func getMask(line string) string {
	var m string
	r, _ := regexp.Compile("^mask\\s=\\s([\\dX]+)$")
	match := r.FindStringSubmatch(line)
	if len(match) == 2 {
		return match[1]
	}
	return m
}

func initAddressSpace(input []string) []uint64 {
	var max int
	for _, line := range input {
		addr, _ := memAddr(line)
		if addr != nil && *addr > max {
			max = *addr + 1
		}
	}
	return make([]uint64, max)
}

func processMask(s string) (set uint64, clear uint64) {
	for i := range s {
		c := s[len(s)-1-i]
		switch c {
		case '1':
			set |= (1 << i)
		case '0':
			clear |= (1 << i)
		}
	}
	return set, ^clear
}

func partOne(lines []string) uint64 {
	mask, mem := "", initAddressSpace(lines)
	var sum, set, clear uint64

	for _, line := range lines {
		addr, val := memAddr(line)
		if addr != nil {
			*val |= set
			*val &= clear
			mem[*addr] = *val
		} else {
			mask = getMask(line)
			set, clear = processMask(mask)
		}
	}

	for _, val := range mem {
		sum += val
	}

	return sum
}

func bitAt(val uint64, index int) uint64 {
	mask := uint64(1) << index
	return (val & mask) >> index
}

func setBit(val uint64, index int) uint64 {
	return val | (1 << index)
}

func partTwo(lines []string) uint64 {
	var mask string
	var sum uint64
	mem := make(map[uint64]uint64)

	for _, line := range lines {
		addr, val := memAddr(line)
		addresses := []uint64{0}

		if addr != nil {
			for i := 0; i < 36; i++ {
				bitIndex := 35 - i
				switch mask[i] {
				case '0':
					for j := range addresses {
						if bitAt(uint64(*addr), bitIndex) == 1 {
							addresses[j] = setBit(addresses[j], bitIndex)
						}
					}
				case '1':
					for j := range addresses {
						addresses[j] = setBit(addresses[j], bitIndex)
					}
				case 'X':
					for j := range addresses {
						addresses = append(addresses, addresses[j])
						addresses[j] = setBit(addresses[j], bitIndex)
					}
				}
			}

			for _, address := range addresses {
				mem[address] = *val
			}

		} else {
			mask = getMask(line)
		}
	}

	for _, val := range mem {
		sum += val
	}

	return sum
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(data), "\n")
	if err != nil {
		log.Fatal("Error handling input")
	}

	fmt.Println("part 1:", partOne(lines))
	fmt.Println("part 2:", partTwo(lines))
}
