package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) ([]string, error) {
	var err error
	var exs []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		exs = append(exs, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return exs, nil
}

func delPar(ex string, index int) string {
	hasp := true

	for hasp {
		var pc, i, prev, next int
		hasp = false

		for _, c := range ex {
			switch c {
			case '(':
				hasp = true
				if pc == 0 {
					prev = i
				}
				pc++
			case ')':
				pc--
				if pc == 0 {
					next = i
				}
			}
			i++
		}

		if hasp {
			r := strconv.Itoa(evaluate(ex[prev+1 : next]))
			ex = ex[0:prev] + r + ex[next+1:]
		}
	}

	return ex
}

func evaluate(ex string) int {
	var res int
	ex = delPar(ex, 1)

	splitted := strings.Split(ex, " ")
	res = toInt(splitted[0])
	splitted = splitted[1:]
	for len(splitted) > 1 {
		op := splitted[0]
		op2 := string(splitted[1])

		switch op {
		case "+":
			res += toInt(op2)
		case "*":
			res *= toInt(op2)
		}
		splitted = splitted[2:]
	}
	return res
}

func solve(exs []string) int {
	var sum int

	for _, ex := range exs {
		sum += evaluate(ex)
	}

	return sum
}

func delPar2(ex string, index int) string {
	hasp := true

	for hasp {
		var pc, i, prev, next int
		hasp = false

		for _, c := range ex {
			switch c {
			case '(':
				hasp = true
				if pc == 0 {
					prev = i
				}
				pc++
			case ')':
				pc--
				if pc == 0 {
					next = i
				}
			}
			i++
		}

		if hasp {
			r := strconv.Itoa(evaluate2(ex[prev+1 : next]))
			ex = ex[0:prev] + r + ex[next+1:]
		}
	}

	return ex
}

func evaluate2(ex string) int {
	var res int
	ex = delPar2(ex, 1)

	ch := make(chan int, strings.Count(ex, "*"))
	splitted := strings.Split(ex, " ")
	res = toInt(splitted[0])
	splitted = splitted[1:]
	for len(splitted) > 1 {
		op := splitted[0]
		op2 := string(splitted[1])

		switch op {
		case "+":
			res += toInt(op2)
		case "*":
			ch <- res
			res = toInt(op2)
		}
		splitted = splitted[2:]
	}
	close(ch)

	for num := range ch {
		res *= num
	}

	return res
}

func solve2(exs []string) int {
	var sum int

	for _, ex := range exs {
		sum += evaluate2(ex)
	}

	return sum
}

func toInt(c string) int {
	n, _ := strconv.Atoi(c)

	return n
}

func main() {
	var err error

	exs, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(exs))
	fmt.Println(solve2(exs))
}
