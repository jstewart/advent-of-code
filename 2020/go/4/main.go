package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func parsePassPort(entry string) map[string]string {
	pp := make(map[string]string)

	splitter := func(c rune) bool {
		return c == '\n' || c == ' '
	}

	for _, field := range strings.FieldsFunc(entry, splitter) {
		entry := strings.Split(field, ":")
		pp[entry[0]] = entry[1]
	}
	return pp
}

func parseRecords(blob string) []map[string]string {
	var passports []map[string]string

	records := strings.Split(blob, "\n\n")
	for _, record := range records {
		passports = append(passports, parsePassPort(record))
	}

	return passports
}

func presenceValidator(passport map[string]string) bool {
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, rq := range required {
		if passport[rq] == "" {
			return false
		}
	}
	return true
}

func formatValidator(passport map[string]string) bool {

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	byr, err := strconv.Atoi(passport["byr"])
	if err != nil {
		return false
	}
	if byr < 1920 || byr > 2002 {
		return false
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	iyr, err := strconv.Atoi(passport["iyr"])
	if err != nil {
		return false
	}
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	eyr, err := strconv.Atoi(passport["eyr"])
	if err != nil {
		return false
	}
	if eyr < 2010 || eyr > 2030 {
		return false
	}

	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	r, _ := regexp.Compile("^(\\d+)(cm|in)$")
	match := r.FindStringSubmatch(passport["hgt"])
	if len(match) != 3 {
		return false
	}

	unit := match[2]
	measure, err := strconv.Atoi(match[1])
	if err != nil {
		return false
	}

	if unit == "cm" {
		if measure < 150 || measure > 193 {
			return false
		}
	} else {
		if measure < 59 || measure > 76 {
			return false
		}

	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	matched, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, passport["hcl"])
	if !matched {
		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	matched, _ = regexp.MatchString(`^amb|blu|brn|gry|grn|hzl|oth$`, passport["ecl"])
	if !matched {
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	matched, _ = regexp.MatchString(`^\d{9}$`, passport["pid"])
	if !matched {
		return false
	}

	return true
}

func numValidPassPorts(passports []map[string]string, validators []func(map[string]string) bool) int {
	var ans int

	for _, pp := range passports {
		valid := true
		for _, validator := range validators {
			// TODO pass both validators
			if !validator(pp) {
				valid = false
				break
			}
		}
		if valid {
			ans++
		}
	}

	return ans
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		log.Fatal("Error handling input")
	}
	p1Validator := []func(map[string]string) bool{presenceValidator}
	p2Validator := []func(map[string]string) bool{presenceValidator, formatValidator}

	fmt.Println("Part 1:", numValidPassPorts(parseRecords(string(data)), p1Validator))
	fmt.Println("Part 2:", numValidPassPorts(parseRecords(string(data)), p2Validator))
}
