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
	var numbers []int

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
			numbers = append(numbers, num)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

func isSumInList(slice []int, num int) bool {
	var found bool

	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] == slice[j] {
				continue
			} else if slice[i]+slice[j] == num {
				return true
			}
		}
	}

	return found
}

func findWrongAndWeak(numbers []int) (int, int) {
	var wrong, weak, min, max int

	l := 0
	h := 25

	for i := 25; i < len(numbers); i++ {
		if !isSumInList(numbers[l:h], numbers[i]) {
			wrong = numbers[i]
			break
		}
		l++
		h++
	}

loop1:
	for i := 0; i < len(numbers[:h+1])-1; i++ {
		sum := numbers[i]
		min = numbers[i]
		max = numbers[i]
	loop2:
		for j := i + 1; j < len(numbers[:h+1]); j++ {
			sum += numbers[j]
			if numbers[j] < min {
				min = numbers[j]
			}
			if numbers[j] > max {
				max = numbers[j]
			}
			if sum > wrong {
				break loop2
			} else if sum == wrong {
				weak = min + max
				break loop1
			}
		}
	}

	return wrong, weak
}

func main() {
	var err error

	numbers, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(findWrongAndWeak(numbers))
}
