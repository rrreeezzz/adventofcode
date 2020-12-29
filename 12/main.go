package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type instruction struct {
	action byte
	value  int
}

func readFile(filename string) ([]*instruction, error) {
	var err error
	var insts []*instruction

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inst := &instruction{}

		inst.action = line[0]
		if inst.value, err = strconv.Atoi(line[1:]); err != nil {
			return nil, err
		}

		insts = append(insts, inst)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return insts, nil
}

func getDir(dir int) byte {
	switch dir {
	case 0:
		return 'E'
	case 90:
		return 'N'
	case 180:
		return 'W'
	}

	return 'S'
}

func getManhattan(insts []*instruction) int {
	var north, south, east, west int
	var dir, pos int

	for _, inst := range insts {
		action := inst.action

		switch action {
		case 'R':
			dir = (360 + (dir - (inst.value))) % 360
			continue
		case 'L':
			dir = (dir + (inst.value)) % 360
			continue
		case 'F':
			action = getDir(dir)
			fallthrough
		default:
			switch action {
			case 'N':
				north += inst.value
			case 'S':
				south += inst.value
			case 'E':
				east += inst.value
			case 'W':
				west += inst.value
			}
		}
	}

	if north > south {
		pos += north - south
	} else {
		pos += south - north
	}
	if east > west {
		pos += east - west
	} else {
		pos += west - east
	}

	return pos
}

func getManhattan2(insts []*instruction) int {
	var waypX, waypY, boatX, boatY, pos int

	waypX = 10
	waypY = 1

	for _, inst := range insts {
		action := inst.action

		switch action {
		case 'R':
			rot := 90
			for inst.value-rot >= 0 {
				waypX, waypY = waypY, -waypX
				rot += 90
			}

		case 'L':
			rot := 90
			for inst.value-rot >= 0 {
				waypX, waypY = -waypY, waypX
				rot += 90
			}

		case 'F':
			boatX = boatX + inst.value*waypX
			boatY = boatY + inst.value*waypY
		case 'N':
			waypY += inst.value
		case 'S':
			waypY -= inst.value
		case 'E':
			waypX += inst.value
		case 'W':
			waypX -= inst.value
		}
	}

	if boatX > 0 {
		pos += boatX
	} else {
		pos -= boatX
	}
	if boatY > 0 {
		pos += boatY
	} else {
		pos -= boatY
	}

	return pos
}

func main() {
	var err error

	insts, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(getManhattan(insts))
	fmt.Println(getManhattan2(insts))

}
