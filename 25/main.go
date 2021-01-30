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
	var nums []int

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		nums = append(nums, n)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nums, nil
}

func getKey(subjnum, loop int) int {
	num := 1
	for i := 0; i < loop; i++ {
		num = num * subjnum
		num = num % 20201227
	}

	return num
}

func findLoop(pub int) int {
	subjnum := 7

	var loop int
	for {
		if pub == 1 {
			break
		}
		for pub%subjnum != 0 {
			pub += 20201227
		}
		pub = pub / subjnum

		loop++
	}

	return loop
}

func solve(pubs []int) int {

	loop := findLoop(pubs[0])

	return getKey(pubs[1], loop)
}

func main() {
	var err error

	nums, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(nums))
}
