package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func prepare(input []string) [][]interface{} {
	var set [][]interface{}

	for _, ins := range input {
		x := strings.Split(ins, " ")
		arg, _ := strconv.Atoi(x[1])
		set = append(set, []interface{}{x[0], arg})
	}

	return set
}

func runPart1(instructionSet [][]interface{}) (int, int) {
	isRun := make(map[int]bool)
	acc, jmp := 0, 0

	for true {
		if _, ok := isRun[jmp]; ok {
			return acc, 1
		}
		isRun[jmp] = true

		op, arg := instructionSet[jmp][0], instructionSet[jmp][1].(int)
		switch op {
		case "acc":
			acc += arg
			jmp++
		case "jmp":
			jmp += arg
		default:
			// nop
			jmp++
		}

		if jmp >= len(instructionSet) {
			break
		}
	}

	return acc, 0
}

func runPart2(instructionSet [][]interface{}) int {
	for i, ins := range instructionSet {
		orig, op, arg := ins, ins[0], ins[1]

		if op == "nop" {
			instructionSet[i] = []interface{}{"jmp", arg}
			acc, ret := runPart1(instructionSet)
			if ret == 0 {
				return acc
			}
			instructionSet[i] = orig
		}
		if op == "jmp" {
			instructionSet[i] = []interface{}{"nop", arg}
			acc, ret := runPart1(instructionSet)
			if ret == 0 {
				return acc
			}
			instructionSet[i] = orig
		}
	}
	return 0
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

	instructionSet := prepare(input)
	result, retCode := runPart1(instructionSet)
	fmt.Println("Part 1:", result, retCode)
	fmt.Println("Part 2:", runPart2(instructionSet))
}
