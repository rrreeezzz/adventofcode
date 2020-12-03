package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var TREE = '#'
var NOTTREE = '.'

func readFile(filename string) ([][]bool, error) {
	var err error
	var trees [][]bool

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		trees = append(trees, []bool{})
		line := scanner.Text()
		for _, elt := range line {
			if elt == TREE {
				trees[i] = append(trees[i], true)
			} else {
				trees[i] = append(trees[i], false)
			}
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return trees, nil
}

func countTree(trees [][]bool, right, down int) int {
	ntrees := 0
	shift := right

	for row := down; row < len(trees); row += down {
		if trees[row][shift] == true {
			ntrees++
		}

		shift = (shift + right) % len(trees[0])
	}

	return ntrees
}

func main() {
	var err error

	trees, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(countTree(trees, 3, 1))

	fmt.Println(countTree(trees, 1, 1) *
		countTree(trees, 3, 1) *
		countTree(trees, 5, 1) *
		countTree(trees, 7, 1) *
		countTree(trees, 1, 2))
}
