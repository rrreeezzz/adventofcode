package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type seat int

const (
	floor seat = iota
	empty
	occupied
)

func readFile(filename string) ([][]seat, error) {
	var err error
	var seats [][]seat

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		seats = append(seats, []seat{})

		for _, s := range line {
			switch s {
			case '.':
				seats[i] = append(seats[i], floor)
			case 'L':
				seats[i] = append(seats[i], empty)
			}
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return seats, nil
}

func checkDir(prevState [][]seat, i, j, x, y int) seat {

	if i+x >= 0 && j+y >= 0 &&
		i+x < len(prevState) && j+y < len(prevState[i]) {
		s := prevState[i+x][j+y]
		if s != floor {
			return s
		}
	}

	return floor
}

func isSeatable(prevState [][]seat, i, j int) bool {
	var free int
	potential := 8

	dirs := []struct{ x, y int }{
		{-1, -1}, {-1, 0}, {0, -1}, {+1, 0},
		{0, +1}, {-1, +1}, {+1, -1}, {+1, +1},
	}

	for _, dir := range dirs {
		s := checkDir(prevState, i, j, dir.x, dir.y)
		if s < occupied {
			free++
		}
	}

	if free == potential {
		return true
	}
	return false
}

func isEmptyable(prevState [][]seat, i, j int) bool {
	var seated int
	potential := 4

	dirs := []struct{ x, y int }{
		{-1, -1}, {-1, 0}, {0, -1}, {+1, 0},
		{0, +1}, {-1, +1}, {+1, -1}, {+1, +1},
	}

	for _, dir := range dirs {
		s := checkDir(prevState, i, j, dir.x, dir.y)
		if s == occupied {
			seated++
		}
	}

	if seated >= potential {
		return true
	}
	return false
}

func copyTable(seats [][]seat) [][]seat {
	ret := make([][]seat, len(seats))
	for i := range seats {
		ret[i] = make([]seat, len(seats[i]))
		copy(ret[i], seats[i])
	}

	return ret
}

func occupiedSeats(seats [][]seat) int {
	var changed bool
	var prevState [][]seat
	var seated int

	for {
		changed = false
		prevState = copyTable(seats)

		for i := 0; i < len(seats); i++ {
			for j := 0; j < len(seats[i]); j++ {
				if seats[i][j] == empty && isSeatable(prevState, i, j) {
					seats[i][j] = occupied
					seated++
					changed = true
				} else if seats[i][j] == occupied && isEmptyable(prevState, i, j) {
					seats[i][j] = empty
					seated--
					changed = true
				}
			}
		}

		if !changed {
			break
		}
	}

	return seated
}

func checkDir2(prevState [][]seat, i, j, x, y int) seat {
	a := x
	b := y

	for i+a >= 0 && j+b >= 0 &&
		i+a < len(prevState) && j+b < len(prevState[i]) {
		s := prevState[i+a][j+b]
		if s != floor {
			return s
		}
		a += x
		b += y
	}

	return floor
}

func isSeatable2(prevState [][]seat, i, j int) bool {
	var free int
	potential := 8

	dirs := []struct{ x, y int }{
		{-1, -1}, {-1, 0}, {0, -1}, {+1, 0},
		{0, +1}, {-1, +1}, {+1, -1}, {+1, +1},
	}

	for _, dir := range dirs {
		s := checkDir2(prevState, i, j, dir.x, dir.y)
		if s < occupied {
			free++
		}
	}

	if free >= potential {
		return true
	}
	return false
}

func isEmptyable2(prevState [][]seat, i, j int) bool {
	var seated int
	potential := 5

	dirs := []struct{ x, y int }{
		{-1, -1}, {-1, 0}, {0, -1}, {+1, 0},
		{0, +1}, {-1, +1}, {+1, -1}, {+1, +1},
	}

	for _, dir := range dirs {
		s := checkDir2(prevState, i, j, dir.x, dir.y)
		if s == occupied {
			seated++
		}
	}

	if seated >= potential {
		return true
	}
	return false
}

func occupiedSeats2(seats [][]seat) int {
	var changed bool
	var prevState [][]seat
	var seated int

	for {
		changed = false
		prevState = copyTable(seats)

		for i := 0; i < len(seats); i++ {
			for j := 0; j < len(seats[i]); j++ {
				if seats[i][j] == empty && isSeatable2(prevState, i, j) {
					seats[i][j] = occupied
					seated++
					changed = true
				} else if seats[i][j] == occupied && isEmptyable2(prevState, i, j) {
					seats[i][j] = empty
					seated--
					changed = true
				}
			}
		}

		if !changed {
			break
		}
	}

	return seated
}

func main() {
	var err error

	seats, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	orig := copyTable(seats)

	fmt.Println(occupiedSeats(seats))
	fmt.Println(occupiedSeats2(orig))
}
