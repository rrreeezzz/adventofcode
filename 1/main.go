package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(filename string) ([]int, error) {
	var err error
	var numlist []int

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numlist = append(numlist, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numlist, nil
}

func findMultiple(numlist []int, year int) int {
	intmap := make(map[int]bool, len(numlist))
	multiple := -1

	for _, elt := range numlist {
		if _, found := intmap[elt]; found {
			continue
		}
		intmap[elt] = true
		diff := year - elt
		if _, found := intmap[diff]; found {
			multiple = diff * elt
			break
		}
	}

	return multiple
}

func main() {
	var err error

	numlist, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(findMultiple(numlist, 2020))
}
