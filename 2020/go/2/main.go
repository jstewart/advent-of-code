package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// NumValidPasswords analyzes password/policies ans returns the number of valid ones
func NumValidPasswords(pwent []string) int {
	var ans int

	for _, pwPolicy := range pwent {
		matchingChars := 0
		x := strings.Split(pwPolicy, " ")
		charRange, letter, pw := x[0], x[1], x[2]
		xRange := strings.Split(charRange, "-")
		xLower, xUpper := xRange[0], xRange[1]
		lower, _ := strconv.Atoi(xLower)
		upper, _ := strconv.Atoi(xUpper)

		for _, c := range pw {
			if string(c) == letter[:1] {
				matchingChars++
			}

		}

		if matchingChars >= lower && matchingChars <= upper {
			ans++
		}
	}

	return ans
}

func main() {
	var input []string

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of valid passwords: ", NumValidPasswords(input))
}
