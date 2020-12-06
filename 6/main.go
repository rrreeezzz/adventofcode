package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(filename string) ([][]string, error) {
	var err error
	var answers [][]string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	answers = append(answers, []string{})
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			answers = append(answers, []string{})
			i++
			continue
		}
		answers[i] = append(answers[i], line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return answers, nil
}

func countYesGroup(group []string) int {
	mmap := make(map[rune]bool)
	for _, line := range group {
		for _, char := range line {
			mmap[char] = true
		}
	}

	return len(mmap)
}

func countAllYes(answers [][]string) int {
	var sum int

	for _, grp := range answers {
		sum += countYesGroup(grp)
	}

	return sum
}

func countEveryoneYesGroup(group []string) int {
	mmap := make(map[rune]int)
	for _, line := range group {
		for _, char := range line {
			mmap[char]++
		}
	}

	for key, cnt := range mmap {
		if cnt != len(group) {
			delete(mmap, key)
		}
	}

	return len(mmap)
}

func countEveryoneYes(answers [][]string) int {
	var sum int

	for _, grp := range answers {
		sum += countEveryoneYesGroup(grp)
	}

	return sum
}

func main() {
	var err error

	answers, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(countAllYes(answers))
	fmt.Println(countEveryoneYes(answers))
}
