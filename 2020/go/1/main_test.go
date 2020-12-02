package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"testing"
)

func getInputs() []int {
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

	return input
}

func TestTwoSum2020(t *testing.T) {
	inputs := getInputs()
	expected := 731731
	actual := TwoSum2020(inputs)

	if actual != expected {
		t.Errorf("Expected %d but instead got %d!", expected, actual)
	}
}

func TestThreeSum2020(t *testing.T) {
	inputs := getInputs()
	expected := 116115990
	actual := ThreeSum2020(inputs)

	if actual != expected {
		t.Errorf("Expected %d but instead got %d!", expected, actual)
	}
}

func BenchmarkTwoSum2020(b *testing.B) {
	inputs := getInputs()

	for n := 0; n < b.N; n++ {
		TwoSum2020(inputs)
	}
}

func BenchmarkShittyTwoSum2020(b *testing.B) {
	inputs := getInputs()

	for n := 0; n < b.N; n++ {
		ShittyTwoSum2020(inputs)
	}
}
func BenchmarkShittyThreeSum2020(b *testing.B) {
	inputs := getInputs()

	for n := 0; n < b.N; n++ {
		ShittyThreeSum2020(inputs)
	}
}

func BenchmarkThreeSum2020(b *testing.B) {
	inputs := getInputs()

	for n := 0; n < b.N; n++ {
		ThreeSum2020(inputs)
	}
}
