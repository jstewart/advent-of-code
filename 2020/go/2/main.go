package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pwEnt struct {
	Letter   string
	Password string
	Lower    int
	Upper    int
}

func parsePwEnt(entry string) pwEnt {
	x := strings.Split(entry, " ")
	charRange, letter, pw := x[0], x[1], x[2]
	xRange := strings.Split(charRange, "-")
	xLower, xUpper := xRange[0], xRange[1]
	lower, _ := strconv.Atoi(xLower)
	upper, _ := strconv.Atoi(xUpper)

	return pwEnt{letter[:1], pw, lower, upper}
}

// NumValidPasswords analyzes password/policies ans returns the number of valid ones
func NumValidPasswords(passwd []string, matcher func(pwEnt) bool) int {
	var ans int

	for _, pwPolicy := range passwd {
		entry := parsePwEnt(pwPolicy)

		if matcher(entry) {
			ans++
		}

	}

	return ans
}

func rangeMatcher(entry pwEnt) bool {
	matchingChars, ans := 0, 0

	for _, c := range entry.Password {
		if string(c) == entry.Letter {
			matchingChars++
		}
	}

	if matchingChars >= entry.Lower && matchingChars <= entry.Upper {
		ans++
		return true
	}

	return false
}

func positionMatcher(entry pwEnt) bool {
	s := strings.Split(entry.Password, "")
	lower, upper := s[entry.Lower-1], s[entry.Upper-1]

	if !(lower == entry.Letter && upper == entry.Letter) &&
		(lower == entry.Letter || upper == entry.Letter) {
		return true
	}

	return false
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

	fmt.Println("Number of valid passwords (1): ", NumValidPasswords(input, rangeMatcher))
	fmt.Println("Number of valid passwords (2): ", NumValidPasswords(input, positionMatcher))
}
