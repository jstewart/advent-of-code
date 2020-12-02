package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// TwoSum2020 accepts and expense list and returns the product of the 2 items that total 2020
func TwoSum2020(expenses []int) int {
	var ans int

	addendCache := make(map[int]int)

	for i, expense := range expenses {
		diff := 2020 - expense

		if addend, ok := addendCache[diff]; ok {
			return expense * expenses[addend]
		}
		addendCache[expense] = i
	}

	return ans
}

// ShittyTwoSum2020 is a shitty version of a better function
func ShittyTwoSum2020(expenses []int) int {
	var ans int

	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses); j++ {
			if expenses[i]+expenses[j] == 2020 {
				return expenses[i] * expenses[j]
			}
		}
	}

	return ans
}

// ShittyThreeSum2020 is a shitty version of a better function
func ShittyThreeSum2020(expenses []int) int {
	var ans int

	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses); j++ {
			for k := i + 1; k < len(expenses); k++ {
				if expenses[i]+expenses[j]+expenses[k] == 2020 {
					return expenses[i] * expenses[j] * expenses[k]
				}
			}
		}
	}

	return ans
}

// ThreeSum2020 accepts an expense list and returns the product of the 3 items that total 2020
func ThreeSum2020(expenses []int) int {
	var ans int
	sort.Ints(expenses)

	for i := 0; i < len(expenses)-2; i++ {
		l := i + 1
		r := len(expenses) - 1
		for l < r {
			sum := expenses[i] + expenses[l] + expenses[r]
			//fmt.Printf("i: %d l: %d r: %d | %d %d %d | %d\n", i, l, r, expenses[i], expenses[l], expenses[r], sum)
			if sum == 2020 {
				return expenses[i] * expenses[l] * expenses[r]
			}
			if sum < 2020 {
				l++
			} else {
				r--
			}
		}
	}

	return ans
}

func main() {
	var input []int

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		input = append(input, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Two-Sum answer is: ", TwoSum2020(input))
	fmt.Println("Three-Sum answer is: ", ThreeSum2020(input))
}
