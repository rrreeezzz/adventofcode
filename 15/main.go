package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	scanner.Scan()
	line := strings.Split(scanner.Text(), ",")

	for _, elt := range line {
		var num int

		if num, err = strconv.Atoi(elt); err != nil {
			return nil, err
		}

		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nums, nil
}

func lastNsaid(nums []int, n int) int {
	type tuple struct {
		inf, sup int
	}
	memory := make(map[int]*tuple)
	var queue []int

	turn := 1
	for _, num := range nums[:len(nums)-1] {
		memory[num] = &tuple{inf: turn}
		turn++
	}

	queue = append(queue, nums[len(nums)-1])

	for len(queue) > 0 {
		num := queue[0]
		if turn == n {
			break
		}

		if t, ok := memory[num]; !ok {
			memory[num] = &tuple{inf: turn}
			queue = append(queue, 0)
		} else {
			queue = append(queue, turn-t.inf)
			t.inf = turn
		}
		queue = queue[1:]
		turn++
	}

	return queue[len(queue)-1]

}

func main() {
	var err error

	nums, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lastNsaid(nums, 30000000))
}
