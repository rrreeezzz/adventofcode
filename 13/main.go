package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type infos struct {
	arrival int
	buses   []int
}

func readFile(filename string) (*infos, error) {
	var err error
	infos := &infos{}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if infos.arrival, err = strconv.Atoi(scanner.Text()); err != nil {
		return nil, err
	}
	scanner.Scan()
	buses := strings.Split(scanner.Text(), ",")
	for _, elt := range buses {
		if elt != "x" {
			var num int
			if num, err = strconv.Atoi(elt); err != nil {
				return nil, err
			}
			infos.buses = append(infos.buses, num)
		} else {
			infos.buses = append(infos.buses, -1)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return infos, nil
}

func findBus(infos *infos) int {
	nearest := infos.buses[0]
	min := -1

	for _, bus := range infos.buses {
		if bus < 0 {
			continue
		}
		tmp := (bus*(infos.arrival/bus+1) - infos.arrival)
		if (tmp >= 0 && tmp < min) || min < 0 {
			min = tmp
			nearest = bus
		}
	}

	return min * nearest
}

func isSubsequentBus(timestamp, bus int) bool {
	if timestamp%bus == 0 {
		return true
	}

	return false
}

func findBus2(infos *infos) int {
	var timestamp int

	mul := 1
	for i, bus := range infos.buses {
		if bus < 0 {
			continue
		}
		for !isSubsequentBus(timestamp+i, bus) {
			timestamp += mul
		}
		mul *= bus
	}

	return timestamp
}

func main() {
	var err error

	infos, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(findBus(infos))
	fmt.Println(findBus2(infos))
}
