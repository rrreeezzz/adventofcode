package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var neighborhood []cube

type cube struct {
	x, y, z, w int
}

func (c cube) add(nc cube) cube {
	return cube{c.x + nc.x, c.y + nc.y, c.z + nc.z, c.w + nc.w}
}

func readFile(filename string) (map[cube]bool, error) {
	var err error
	cubes := make(map[cube]bool)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for x, elt := range line {
			c := cube{}
			if elt == '#' {
				c.x = x
				c.y = y
				cubes[c] = true
			}
		}
		y++

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cubes, nil
}

func buildNeighborhood() {
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			for z := -1; z < 2; z++ {
				for w := -1; w < 2; w++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					neighborhood = append(neighborhood, cube{x, y, z, w})
				}
			}
		}
	}
}

func activeNeighbors(c cube, cubes map[cube]bool) int {
	var active int

	for _, n := range neighborhood {
		if _, ok := cubes[c.add(n)]; ok {
			active++
		}
	}

	return active
}

func getNeighbors(c cube, cubes map[cube]bool) []cube {
	var nei []cube

	for _, n := range neighborhood {
		nei = append(nei, c.add(n))
	}

	return nei
}

func conway(cubes map[cube]bool, cycles int) int {

	for cycle := 0; cycle < cycles; cycle++ {
		prevCubes := make(map[cube]bool)
		for k, v := range cubes {
			prevCubes[k] = v
		}

		cmap := []cube{}
		for c := range cubes {
			cmap = append(cmap, getNeighbors(c, prevCubes)...)
			cmap = append(cmap, c)
		}

		for _, c := range cmap {
			countn := activeNeighbors(c, prevCubes)
			if prevCubes[c] {
				if countn < 2 || countn > 3 {
					delete(cubes, c)
				}
			} else {
				if countn == 3 {
					cubes[c] = true
				}
			}
		}
	}

	return len(cubes)
}

func main() {
	var err error

	cubes, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	buildNeighborhood()
	fmt.Println(conway(cubes, 6))
}
