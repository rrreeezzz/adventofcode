package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) ([]map[string]string, error) {
	var err error
	var passports []map[string]string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	passports = append(passports, map[string]string{})
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, map[string]string{})
			i++
			continue
		}

		for _, elt := range strings.Split(line, " ") {
			key := strings.Split(elt, ":")[0]
			value := strings.Split(elt, ":")[1]
			passports[i][key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return passports, nil
}

func validateRange(v, low, high int) bool {
	if v >= low && v <= high {
		return true
	}

	return false
}

func validateStringRange(s string, low, high int) bool {
	if y, err := strconv.Atoi(s); err != nil {
		return false
	} else {
		return validateRange(y, low, high)
	}
}

func validateHeight(height string) bool {
	if s := strings.TrimSuffix(height, "cm"); s != height {
		return validateStringRange(s, 150, 193)
	} else if s := strings.TrimSuffix(height, "in"); s != height {
		return validateStringRange(s, 59, 76)
	}

	return false
}

func validateColorCode(code string) bool {
	if len(code) != 7 {
		return false
	}

	var r, g, b int

	if _, err := fmt.Sscanf(code, "#%02x%02x%02x", &r, &g, &b); err != nil {
		return false
	}

	return true
}

func validateEyeColor(color string) bool {
	switch color {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

func validatePid(num string) bool {
	if len(num) != 9 {
		return false
	}

	if _, err := strconv.Atoi(num); err != nil {
		return false
	}

	return true
}

func countValidPassports(passports []map[string]string) int {
	var valid int

	for _, pass := range passports {
		if v, ok := pass["byr"]; !ok {
			continue
		} else {
			if !validateStringRange(v, 1920, 2002) {
				continue
			}
		}

		if v, ok := pass["iyr"]; !ok {
			continue
		} else {
			if !validateStringRange(v, 2010, 2020) {
				continue
			}
		}

		if v, ok := pass["eyr"]; !ok {
			continue
		} else {
			if !validateStringRange(v, 2020, 2030) {
				continue
			}
		}

		if v, ok := pass["hgt"]; !ok {
			continue
		} else {
			if !validateHeight(v) {
				continue
			}
		}

		if v, ok := pass["hcl"]; !ok {
			continue
		} else {
			if !validateColorCode(v) {
				continue
			}
		}

		if v, ok := pass["ecl"]; !ok {
			continue
		} else {
			if !validateEyeColor(v) {
				continue
			}
		}

		if v, ok := pass["pid"]; !ok {
			continue
		} else {
			if !validatePid(v) {
				continue
			}
		}

		valid++
	}

	return valid
}

func main() {
	var err error

	passports, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(countValidPassports(passports))
}
