package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pass struct {
	lower    int
	higher   int
	letter   string
	password string
}

func readFile(filename string) ([]*pass, error) {
	var err error
	var passlist []*pass

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var rng, l, h string
		p := &pass{}
		if _, err := fmt.Sscanf(scanner.Text(), "%s %s %s",
			&rng, &p.letter, &p.password); err != nil {
			continue
		}
		l = strings.Split(rng, "-")[0]
		if p.lower, err = strconv.Atoi(l); err != nil {
			return nil, err
		}
		h = strings.Split(rng, "-")[1]
		if p.higher, err = strconv.Atoi(h); err != nil {
			return nil, err
		}
		p.letter = strings.TrimRight(p.letter, ":")
		passlist = append(passlist, p)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return passlist, nil
}

func verifyPasswordPartOne(passlist []*pass) int {
	var valid int

	for _, p := range passlist {
		cnt := strings.Count(p.password, p.letter)
		if cnt >= p.lower && cnt <= p.higher {
			valid++
		}
	}

	return valid
}

func verifyPasswordPartTwo(passlist []*pass) int {
	var valid int

	for _, p := range passlist {
		bletter := []byte(p.letter)[0]
		l := p.password[p.lower-1]
		h := p.password[p.higher-1]

		if (l == bletter || h == bletter) &&
			!(l == bletter && h == bletter) {
			valid++
		}
	}

	return valid
}

func main() {
	var err error

	passlist, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(verifyPasswordPartOne(passlist))
	fmt.Println(verifyPasswordPartTwo(passlist))

}
