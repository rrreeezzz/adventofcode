package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(filename string) ([][]string, error) {
	var err error
	var paths [][]string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var pnum int
	for scanner.Scan() {
		line := scanner.Text()

		paths = append(paths, []string{})

		for i := 0; i < len(line); i++ {
			switch line[i] {
			case 'n':
				switch line[i+1] {
				case 'e':
					paths[pnum] = append(paths[pnum], "ne")
				case 'w':
					paths[pnum] = append(paths[pnum], "nw")
				}
				i++
			case 's':
				switch line[i+1] {
				case 'e':
					paths[pnum] = append(paths[pnum], "se")
				case 'w':
					paths[pnum] = append(paths[pnum], "sw")
				}
				i++
			case 'e':
				paths[pnum] = append(paths[pnum], "e")
			case 'w':
				paths[pnum] = append(paths[pnum], "w")
			}
		}
		pnum++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return paths, nil
}

func solve(paths [][]string) (int, int) {
	var nblack1, nblack2 int

	type pos struct {
		x, y int
	}

	type tile struct {
		color int
		pos   pos
	}

	cachepos := make(map[pos]tile)

	ref := tile{pos: pos{0, 0}}
	cachepos[ref.pos] = ref

	for _, path := range paths {
		curr := ref
		for _, dir := range path {
			npos := curr.pos
			switch dir {
			case "ne":
				npos.y++
			case "nw":
				npos.x--
				npos.y++
			case "se":
				npos.x++
				npos.y--
			case "sw":
				npos.y--
			case "e":
				npos.x++
			case "w":
				npos.x--
			}

			if _, ok := cachepos[npos]; !ok {
				cachepos[npos] = tile{pos: npos}
			}
			curr = cachepos[npos]
		}

		if curr.color == 0 {
			curr.color = 1
			nblack1++
		} else {
			curr.color = 0
			nblack1--
		}
		cachepos[curr.pos] = curr
	}

	// TO REFACTOR this is taking too much time
	// use a cache to remember where the black tile are
	// and delete those who go white
	for i := 0; i < 100; i++ {
		prevstate := make(map[pos]tile)
		for p, t := range cachepos {
			prevstate[p] = t
		}

		for p := range prevstate {
			if _, ok := prevstate[pos{p.x - 1, p.y + 1}]; !ok {
				prevstate[pos{p.x - 1, p.y + 1}] = tile{0, pos{p.x - 1, p.y + 1}}
			}
			if _, ok := prevstate[pos{p.x + 1, p.y - 1}]; !ok {
				prevstate[pos{p.x + 1, p.y - 1}] = tile{0, pos{p.x + 1, p.y - 1}}
			}
			if _, ok := prevstate[pos{p.x, p.y - 1}]; !ok {
				prevstate[pos{p.x, p.y - 1}] = tile{0, pos{p.x, p.y - 1}}
			}
			if _, ok := prevstate[pos{p.x, p.y + 1}]; !ok {
				prevstate[pos{p.x, p.y + 1}] = tile{0, pos{p.x, p.y + 1}}
			}
			if _, ok := prevstate[pos{p.x + 1, p.y}]; !ok {
				prevstate[pos{p.x + 1, p.y}] = tile{0, pos{p.x + 1, p.y}}
			}
			if _, ok := prevstate[pos{p.x - 1, p.y}]; !ok {
				prevstate[pos{p.x - 1, p.y}] = tile{0, pos{p.x - 1, p.y}}
			}
		}

		for p, t := range prevstate {
			var adj int
			if prevstate[pos{p.x - 1, p.y + 1}].color == 1 {
				adj++
			}
			if prevstate[pos{p.x + 1, p.y - 1}].color == 1 {
				adj++
			}
			if prevstate[pos{p.x, p.y - 1}].color == 1 {
				adj++
			}
			if prevstate[pos{p.x, p.y + 1}].color == 1 {
				adj++
			}
			if prevstate[pos{p.x + 1, p.y}].color == 1 {
				adj++
			}
			if prevstate[pos{p.x - 1, p.y}].color == 1 {
				adj++
			}

			if t.color == 0 && adj == 2 {
				t.color = 1
			} else if adj == 0 || adj > 2 {
				t.color = 0
			}
			cachepos[p] = t
		}
	}

	for _, t := range cachepos {
		if t.color == 1 {
			nblack2++
		}
	}

	return nblack1, nblack2
}

func main() {
	var err error

	paths, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(paths))
}
