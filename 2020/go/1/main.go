package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// cache of 2020 - expense
// key is 2020 - expense
// val is index

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

func shittyTwoSum2020(expenses []int) int {
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

// ThreeSum2020 accepts and expense list and returns the product of the 2 items that total 2020
func ThreeSum2020(expenses []int) int {
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

func main() {
	var input, differences []int

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		input = append(input, i)
		differences = append(differences, 2020-i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Two-Sum answer is: ", TwoSum2020(input))
	fmt.Println("Three-Sum answer is: ", ThreeSum2020(input))
}
