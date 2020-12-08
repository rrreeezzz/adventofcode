package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	name    string
	value   int
	visited int
}

func readFile(filename string) ([]*instruction, error) {
	var err error
	var instructions []*instruction

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inst := &instruction{}

		splitted := strings.Split(line, " ")
		inst.name = splitted[0]

		if inst.value, err = strconv.Atoi(splitted[1]); err != nil {
			return nil, err
		}

		instructions = append(instructions, inst)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return instructions, nil
}

func findAccBeforeLoop(instructions []*instruction) int {
	var ip, acc int

loop:
	for {
		inst := instructions[ip]
		inst.visited++
		if inst.visited > 1 || ip > len(instructions) {
			break loop
		}

		switch inst.name {
		case "nop":
			ip++
		case "acc":
			ip++
			acc += inst.value
		case "jmp":
			ip += inst.value
		}
	}

	return acc
}

func checkLoop(instructions []*instruction) (bool, int) {
	var ip, acc int
	var isloop bool

loop:
	for {
		if ip >= len(instructions) {
			break loop
		}
		inst := instructions[ip]
		inst.visited++
		if inst.visited > 1 {
			isloop = true
			break loop
		}

		switch inst.name {
		case "nop":
			ip++
		case "acc":
			ip++
			acc += inst.value
		case "jmp":
			ip += inst.value
		}
	}

	return isloop, acc
}

func resetCounter(instructions []*instruction) {
	for _, inst := range instructions {
		inst.visited = 0
	}
}

func changeOneInst(instructions []*instruction) int {
	var acc int
	var isloop bool

	for _, inst := range instructions {
		resetCounter(instructions)
		switch inst.name {
		case "nop":
			inst.name = "jmp"
		case "jmp":
			inst.name = "nop"
		}

		isloop, acc = checkLoop(instructions)
		if !isloop {
			break
		}

		switch inst.name {
		case "nop":
			inst.name = "jmp"
		case "jmp":
			inst.name = "nop"
		}
	}

	return acc
}

func main() {
	var err error

	instructions, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(findAccBeforeLoop(instructions))
	fmt.Println(changeOneInst(instructions))
}
