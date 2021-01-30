package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) ([]int, []int, error) {
	var err error
	var deck1, deck2 []int

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	player2 := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Player 1") {
			continue
		} else if strings.Contains(line, "Player 2") {
			player2 = true
			continue
		} else if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, nil, err
		}

		if player2 {
			deck2 = append(deck2, num)
		} else {
			deck1 = append(deck1, num)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return deck1, deck2, nil
}

func getResult(deck1, deck2 []int) int {
	var sum int

	if len(deck1) == 0 {
		deck1 = deck2
	}

	for i := 0; i < len(deck1); i++ {
		sum += (i + 1) * deck1[len(deck1)-1-i]
	}

	return sum
}

func solve1(deck1, deck2 []int) int {
	for len(deck1) > 0 && len(deck2) > 0 {
		elt1 := deck1[0]
		elt2 := deck2[0]

		deck1 = deck1[1:]
		deck2 = deck2[1:]

		if elt1 > elt2 {
			deck1 = append(deck1, elt1, elt2)
		} else {
			deck2 = append(deck2, elt2, elt1)
		}
	}

	return getResult(deck1, deck2)
}

func solve2(deck1, deck2 []int) ([]int, []int) {
	cache := make(map[string]bool)

	for len(deck1) > 0 && len(deck2) > 0 {
		elt1 := deck1[0]
		elt2 := deck2[0]

		deck1 = deck1[1:]
		deck2 = deck2[1:]

		key := fmt.Sprintf("%+v / %+v", deck1, deck2)
		if cache[key] {
			deck1 = append(deck1, deck2...)
			return deck1, []int{}
		}
		cache[key] = true

		var p1Winner bool
		if len(deck1) >= elt1 && len(deck2) >= elt2 {
			d1 := make([]int, len(deck1[:elt1]))
			d2 := make([]int, len(deck2[:elt2]))
			copy(d1, deck1[:elt1])
			copy(d2, deck2[:elt2])
			nd1, nd2 := solve2(d1, d2)
			if len(nd1) > len(nd2) {
				p1Winner = true
			}
		} else {
			if elt1 > elt2 {
				p1Winner = true
			}
		}

		if p1Winner {
			deck1 = append(deck1, elt1, elt2)
		} else {
			deck2 = append(deck2, elt2, elt1)
		}
	}

	return deck1, deck2
}

func main() {
	var err error

	deck1, deck2, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve1(deck1, deck2))
	fmt.Println(getResult(solve2(deck1, deck2)))
}
