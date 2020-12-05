package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(filename string) ([]string, error) {
	var err error
	var seats []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seats = append(seats, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return seats, nil
}

func findRow(seat string) int {
	h := 127
	l := 0
	var value int

	for _, c := range seat[0:7] {
		if c == 'F' {
			h = (h + l) / 2
			value = h
		} else {
			l = (h + l + 1) / 2
			value = l
		}
	}
	return value
}

func findColumn(seat string) int {
	h := 7
	l := 0
	var value int

	for _, c := range seat[7:] {
		if c == 'L' {
			h = (h + l) / 2
			value = h
		} else {
			l = (h + l + 1) / 2
			value = l
		}
	}
	return value
}

func findHighestSeatID(seats []string) int {
	var max int

	for _, seat := range seats {
		val := findRow(seat)*8 + findColumn(seat)
		if val >= max {
			max = val
		}
	}

	return max
}

func findMySeat(seats []string) int {
	idmap := make(map[int]bool)

	for _, seat := range seats {
		idmap[findRow(seat)*8+findColumn(seat)] = true
	}

	for i := 1; i < (127*8 + 7); i++ {
		if _, ok := idmap[i]; ok {
			continue
		}
		_, prev := idmap[i-1]
		_, next := idmap[i+1]

		if next && prev {
			return i
		}
	}

	return 0
}

func main() {
	var err error

	seats, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(findHighestSeatID(seats))
	fmt.Println(findMySeat(seats))
}
