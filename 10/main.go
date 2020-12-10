package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func readFile(filename string) ([]int, error) {
	var err error
	var adaptaters []int

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if num, err := strconv.Atoi(line); err != nil {
			return nil, err
		} else {
			adaptaters = append(adaptaters, num)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return adaptaters, nil
}

func find13Differences(adaptaters []int) int {
	var diff1, diff3, prev int

	sort.Ints(adaptaters)

	for _, adpt := range adaptaters {
		diff := adpt - prev

		if diff == 1 {
			diff1++
		} else if diff == 3 {
			diff3++
		} else if diff > 3 {
			fmt.Println("ERROR")
		}

		prev = adpt
	}

	diff3++
	return diff1 * diff3
}

func findComb(adaptaters []int) int {
	adaptaters = append(adaptaters, 0)
	sort.Ints(adaptaters)

	comb := make([]int, len(adaptaters))
	comb[0]++
	for i := 0; i < len(adaptaters); i++ {
		for j := i + 1; j < len(adaptaters); j++ {
			diff := adaptaters[j] - adaptaters[i]
			if diff > 3 {
				break
			}
			comb[j] += comb[i]
		}
	}
	return comb[len(comb)-1]
}

func main() {
	var err error

	adaptaters, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(find13Differences(adaptaters))
	fmt.Println(findComb(adaptaters))
}
