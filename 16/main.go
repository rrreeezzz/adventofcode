package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	rule  string
	rnges []int
	field int
}

func setRange(r *rule, str string) error {
	var err error
	var l, u int
	rnge := strings.Split(str, "-")

	if l, err = strconv.Atoi(rnge[0]); err != nil {
		return err
	}
	r.rnges = append(r.rnges, l)
	if u, err = strconv.Atoi(rnge[1]); err != nil {
		return err
	}
	r.rnges = append(r.rnges, u)

	return nil
}

func readFile(filename string) ([]*rule, []int, [][]int, error) {
	var err error
	var ticket []int
	var tickets [][]int
	var rules []*rule
	var level int

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, err
	}
	defer file.Close()

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			level++
			continue
		} else if line == "your ticket:" {
			continue
		} else if line == "nearby tickets:" {
			continue
		}

		switch level {
		case 0:
			r := &rule{}
			splitted := strings.Split(line, ":")

			r.rule = splitted[0]

			splitted = strings.Split(splitted[1], " ")

			if err = setRange(r, splitted[1]); err != nil {
				return nil, nil, nil, err
			}

			if err = setRange(r, splitted[3]); err != nil {
				return nil, nil, nil, err
			}

			rules = append(rules, r)
		case 1:
			splitted := strings.Split(line, ",")
			for _, num := range splitted {
				if n, err := strconv.Atoi(num); err != nil {
					return nil, nil, nil, err
				} else {
					ticket = append(ticket, n)
				}
			}
		case 2:
			tickets = append(tickets, []int{})
			splitted := strings.Split(line, ",")
			for _, num := range splitted {
				if n, err := strconv.Atoi(num); err != nil {
					return nil, nil, nil, err
				} else {
					tickets[i] = append(tickets[i], n)
				}
			}
			i++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, nil, err
	}

	return rules, ticket, tickets, nil
}

func isInRanges(rules []*rule, num int) bool {
	for _, r := range rules {
		if num >= r.rnges[0] && num <= r.rnges[1] {
			return true
		} else if num >= r.rnges[2] && num <= r.rnges[3] {
			return true
		}
	}

	return false
}

func isInRange(r *rule, num int) bool {
	if num >= r.rnges[0] && num <= r.rnges[1] {
		return true
	} else if num >= r.rnges[2] && num <= r.rnges[3] {
		return true
	}

	return false
}

func invalids(rules []*rule, tickets [][]int) int {
	var sum int

	for _, ticket := range tickets {
		for _, num := range ticket {
			if !isInRanges(rules, num) {
				sum += num
			}
		}
	}

	return sum
}

func discardInvalids(rules []*rule, tickets [][]int) [][]int {
	var newtickets [][]int

	for _, ticket := range tickets {
		isValid := true
		for _, num := range ticket {
			if !isInRanges(rules, num) {
				isValid = false
				break
			}
		}
		if isValid {
			newtickets = append(newtickets, ticket)
		}
	}

	return newtickets
}

func findRules(rules []*rule, myticket []int, tickets [][]int) int {
	var field, irule int
	mul := 1
	possible := make(map[*rule]map[int]bool)

	for irule < len(rules) {
		rule := rules[irule]
		all := true
		for _, ticket := range tickets {
			/* if irule == 0 && field == 1 {
				fmt.Println(rule.rnges, ticket[field])
			} */
			if !isInRange(rule, ticket[field]) {
				all = false
				break
			}
		}

		if all {
			if _, ok := possible[rule]; !ok {
				possible[rule] = make(map[int]bool)
			}
			possible[rule][field] = true
		}
		field++
		if field == len(rules) {
			field = 0
			irule++
		}
	}

	todel := 0
	mapped := len(rules)
	for mapped > 0 {
		for r, fs := range possible {
			if len(fs) == 1 {
				for f := range fs {
					r.field = f
					todel = f
				}
				mapped--
			}
		}

		for _, fs := range possible {
			delete(fs, todel)
		}
	}

	for _, rule := range rules {
		splitted := strings.Split(rule.rule, " ")
		if len(splitted) < 2 {
			continue
		} else if splitted[0] == "departure" {
			mul *= myticket[rule.field]
		}
	}

	return mul
}

func main() {
	var err error

	rules, ticket, tickets, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(invalids(rules, tickets))
	newtickets := discardInvalids(rules, tickets)
	fmt.Println(findRules(rules, ticket, newtickets))
}
