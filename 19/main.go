package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile(filename string) ([]string, []string, error) {
	var err error
	var rules, msgs []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var isRule bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isRule = true
			continue
		}
		if !isRule {
			splitted := strings.Split(line, ":")

			n, _ := strconv.Atoi(splitted[0])

			l := len(rules)
			for i := 0; i < n-l+1; i++ {
				rules = append(rules, "")
			}

			rules[n] = strings.ReplaceAll(splitted[1], "\"", "")
		}
		if isRule {
			msgs = append(msgs, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return rules, msgs, nil
}

func expandRule(cache map[string][]string, rule string, rules []string) []string {
	var exp, right, left []string
	var or bool

	splitted := strings.Split(strings.TrimSpace(rule), " ")

	for _, c := range splitted {
		if c == "|" {
			or = true
			continue
		}

		if n, err := strconv.Atoi(c); err == nil {
			tmp, ok := cache[rules[n]]
			if !ok {
				tmp = expandRule(cache, rules[n], rules)
			}
			ptr := &left
			if or {
				ptr = &right
			}

			if len(*ptr) == 0 {
				*ptr = tmp
			} else {

				var newexp []string
				for _, str1 := range tmp {
					for _, str2 := range *ptr {
						newexp = append(newexp, str2+str1)
					}
				}

				*ptr = newexp
			}
		} else {
			left = append(left, c)
		}
	}

	exp = append(exp, left...)
	exp = append(exp, right...)

	cache[rule] = exp

	return exp
}

func solve(rules, msgs []string) int {
	var expRules []string
	var sum int

	cache := make(map[string][]string)

	expRules = append(expRules, expandRule(cache, rules[0], rules)...)

	for _, msg := range msgs {
		for _, rule := range expRules {
			if msg == rule {
				sum++
				break
			}
		}
	}

	return sum
}

func reg(rulesdict map[int]string, num int) string {
	if num == 8 {
		return reg(rulesdict, 42) + "+"
	} else if num == 11 {
		a := reg(rulesdict, 42)
		b := reg(rulesdict, 31)
		tmp := "(?:"
		for i := 1; i < 99; i++ {
			is := strconv.Itoa(i)
			tmp += a + "{" + is + "}" + b + "{" + is + "}" + "|"
		}
		tmp += a + "{99}" + b + "{99}" + ")"
		return tmp
	}

	rule := rulesdict[num]
	if rule == " a" || rule == " b" {
		return string(rule[1])
	}

	rule = strings.TrimSpace(rule)

	splitted := strings.Split(rule, " | ")
	var res []string
	for _, p := range splitted {
		nums := strings.Split(p, " ")
		var tmp string
		for _, n := range nums {
			nint, _ := strconv.Atoi(n)
			tmp += reg(rulesdict, nint)
		}
		res = append(res, tmp)
	}

	tmp := "(?:"
	for i := 0; i < len(res)-1; i++ {
		tmp += res[i] + "|"
	}
	tmp += res[len(res)-1] + ")"

	return tmp
}

func solve2(rules, msgs []string) int {
	var sum int

	rulesdict := make(map[int]string)
	for i, rule := range rules {
		rulesdict[i] = rule
	}

	re := regexp.MustCompile("^" + reg(rulesdict, 0) + "$")

	for _, msg := range msgs {
		if ok := re.MatchString(msg); ok {
			sum++
		}
	}

	return sum
}

func main() {
	var err error

	rules, msgs, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(rules, msgs))
	rules[8] = " 42 | 42 8"
	rules[11] = " 42 31 | 42 11 31"

	// did not understand part 2, solution heavily inspired from
	// https://github.com/sophiebits/adventofcode/blob/main/2020/day19.py
	// using regex
	fmt.Println(solve2(rules, msgs))
}
