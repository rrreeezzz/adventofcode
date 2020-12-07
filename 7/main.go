package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type inBagTuple struct {
	count int
	bag   string
}

func buildInverseMap(filename string) (map[string][]*inBagTuple, error) {
	var err error
	bagInBags := make(map[string][]*inBagTuple)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, " bags contain ")
		if splitted[1] == "no other bags." {
			continue
		}

		container := splitted[0]

		for _, sub := range strings.Split(splitted[1], ", ") {
			tuple := &inBagTuple{}
			tmp := strings.Split(sub, " ")
			if tuple.count, err = strconv.Atoi(tmp[0]); err != nil {
				return nil, err
			}
			tuple.bag = container
			bagInBags[fmt.Sprintf("%s %s", tmp[1], tmp[2])] = append(bagInBags[fmt.Sprintf("%s %s", tmp[1], tmp[2])], tuple)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return bagInBags, nil
}

func buildMap(filename string) (map[string][]*inBagTuple, error) {
	var err error
	bagInBags := make(map[string][]*inBagTuple)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, " bags contain ")
		if splitted[1] == "no other bags." {
			continue
		}

		container := splitted[0]

		for _, sub := range strings.Split(splitted[1], ", ") {
			tuple := &inBagTuple{}
			tmp := strings.Split(sub, " ")
			if tuple.count, err = strconv.Atoi(tmp[0]); err != nil {
				return nil, err
			}
			tuple.bag = fmt.Sprintf("%s %s", tmp[1], tmp[2])
			bagInBags[container] = append(bagInBags[container], tuple)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return bagInBags, nil
}

func howManyShiny(bagInBags map[string][]*inBagTuple) int {
	queue := []string{"shiny gold"}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		for _, b := range bagInBags[queue[0]] {
			queue = append(queue, b.bag)
			visited[b.bag] = true
		}
		queue = queue[1:]
	}

	return len(visited)
}

func bagContain(bagInBags map[string][]*inBagTuple, bag string) int {
	sum := 0
	for _, b := range bagInBags[bag] {
		sum += b.count + b.count*bagContain(bagInBags, b.bag)
	}

	return sum
}

func shinyContains(bagInBags map[string][]*inBagTuple) int {
	var sum int

	for _, b := range bagInBags["shiny gold"] {
		sum += b.count + b.count*bagContain(bagInBags, b.bag)
	}

	return sum
}

func main() {
	var err error

	inverseMap, err := buildInverseMap("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(howManyShiny(inverseMap))

	bagMap, err := buildMap("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(shinyContains(bagMap))

}
