package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type tile struct {
	pix [][]uint64
	id  uint64
}

func (t tile) inverse() tile {
	nt := tile{id: t.id}
	for y := 0; y < len(t.pix); y++ {
		nt.pix = append(nt.pix, []uint64{})
		for x := 0; x < len(t.pix[0]); x++ {
			nt.pix[y] = append(nt.pix[y], t.pix[y][len(t.pix[0])-x-1])
		}
	}

	return nt
}

func (t tile) rotate() tile {
	nt := tile{id: t.id}
	n := len(t.pix[0])
	for y := 0; y < n; y++ {
		nt.pix = append(nt.pix, []uint64{})
		for x := 0; x < n; x++ {
			nt.pix[y] = append(nt.pix[y], t.pix[n-x-1][y])
		}
	}

	return nt
}

func (t tile) isRightCompatible(nt tile) bool {
	n := len(t.pix[0])
	for y := 0; y < n; y++ {
		if t.pix[y][n-1] != nt.pix[y][0] {
			return false
		}
	}

	return true
}

func (t tile) isBottomCompatible(nt tile) bool {
	n := len(t.pix[0])
	for x := 0; x < n; x++ {
		if t.pix[n-1][x] != nt.pix[0][x] {
			return false
		}
	}

	return true
}

func readFile(filename string) ([]tile, error) {
	var err error
	var tiles []tile

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var t tile
	var y int
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Tile") {
			t = tile{}
			y = 0
			line = strings.TrimSuffix(line, ":")
			split := strings.Split(line, " ")
			tmp, _ := strconv.Atoi(split[1])
			t.id = uint64(tmp)
		} else if line == "" {
			tiles = append(tiles, t)
		} else {
			split := strings.Split(line, "")
			t.pix = append(t.pix, []uint64{})
			for _, x := range split {
				switch x {
				case ".":
					t.pix[y] = append(t.pix[y], 0)
				case "#":
					t.pix[y] = append(t.pix[y], 1)
				}
			}
			y++
		}
	}
	tiles = append(tiles, t)

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tiles, nil
}

func getTop(t tile) uint64 {
	var res uint64
	for x := 0; x < len(t.pix); x++ {
		res = res << 1
		res = res | t.pix[0][x]
	}

	return res
}

func getBottom(t tile) uint64 {
	var res uint64
	for x := 0; x < len(t.pix); x++ {
		res = res << 1
		res = res | t.pix[len(t.pix)-1][x]
	}

	return res
}

func getLeft(t tile) uint64 {
	var res uint64
	for y := 0; y < len(t.pix); y++ {
		res = res << 1
		res = res | t.pix[y][0]
	}

	return res
}

func getRight(t tile) uint64 {
	var res uint64
	for y := 0; y < len(t.pix); y++ {
		res = res << 1
		res = res | t.pix[y][len(t.pix)-1]
	}

	return res
}

func reverse10Bits(num uint64) uint64 {
	var res uint64
	for i := 0; i < 10; i++ {
		res = res<<1 | num&1
		num >>= 1
	}
	return res
}

func pasteTile(grid [][]uint64, ystart, xstart int, t tile) {
	for y := 0; y < len(t.pix[0])-2; y++ {
		for x := 0; x < len(t.pix[0])-2; x++ {
			grid[ystart+y][xstart+x] = t.pix[y+1][x+1]
		}
	}
}

func isSeaMonsterInGrid(seaMonster, grid [][]uint64, xstart, ystart int) bool {
	for y := 0; y < len(seaMonster); y++ {
		for x := 0; x < len(seaMonster[0]); x++ {
			if seaMonster[y][x] == 1 && grid[ystart+y][xstart+x] != 1 {
				return false
			}
		}
	}

	return true
}

func count1(array [][]uint64) int {
	var cnt int
	for y := 0; y < len(array); y++ {
		for x := 0; x < len(array[0]); x++ {
			if array[y][x] == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func solve1(tiles []tile) (uint64, []uint64) {
	cache := make(map[uint64]int)
	cache2 := make(map[uint64][]uint64)

	for _, t := range tiles {
		val := getTop(t)
		cache[val]++
		cache2[t.id] = append(cache2[t.id], val)
		val = reverse10Bits(val)
		cache[val]++
		cache2[t.id] = append(cache2[t.id], val)
		val = getBottom(t)
		cache[val]++
		cache2[t.id] = append(cache2[t.id], val)
		val = reverse10Bits(val)
		cache[val]++
		cache2[t.id] = append(cache2[t.id], val)
		val = getLeft(t)
		cache[val]++
		cache2[t.id] = append(cache2[t.id], val)
		val = reverse10Bits(val)
		cache[val]++
		cache2[t.id] = append(cache2[t.id], val)
		val = getRight(t)
		cache[val]++
		cache2[t.id] = append(cache2[t.id], val)
		val = reverse10Bits(val)
		cache[val]++
		cache2[t.id] = append(cache2[t.id], val)
	}

	mul := uint64(1)
	var corners []uint64
	for tid, vals := range cache2 {
		unique := 0
		for _, val := range vals {
			if cache[val] == 1 {
				unique++
			}
		}

		if unique > 2 {
			mul *= tid
			corners = append(corners, tid)
		}
	}

	return mul, corners
}

func solve2(tiles []tile, corners []uint64) int {
	// get a map on tiles
	maptiles := make(map[uint64]tile)
	for _, t := range tiles {
		maptiles[t.id] = t
	}

	// prepare grid
	sq := math.Sqrt(float64(len(tiles)))
	grid := make([][]uint64, int(sq)*(len(tiles[0].pix[0])-2))
	for y := range grid {
		grid[y] = make([]uint64, int(sq)*(len(tiles[0].pix[0])-2))
	}

	// map of side values to neighbors
	cache := make(map[uint64][]uint64)
	for _, t := range tiles {
		val := getTop(t)
		cache[val] = append(cache[val], t.id)
		val = reverse10Bits(val)
		cache[val] = append(cache[val], t.id)
		val = getBottom(t)
		cache[val] = append(cache[val], t.id)
		val = reverse10Bits(val)
		cache[val] = append(cache[val], t.id)
		val = getLeft(t)
		cache[val] = append(cache[val], t.id)
		val = reverse10Bits(val)
		cache[val] = append(cache[val], t.id)
		val = getRight(t)
		cache[val] = append(cache[val], t.id)
		val = reverse10Bits(val)
		cache[val] = append(cache[val], t.id)
	}

	// find first upper left corner
	var next uint64
loop:
	for _, tid := range corners {
		for i := 0; i < 4; i++ {
			if len(cache[getRight(maptiles[tid])]) > 1 && len(cache[getBottom(maptiles[tid])]) > 1 {
				next = tid
				break loop
			}
			maptiles[tid] = maptiles[tid].rotate()
		}
	}

	/*
		Fill row by taking right side and search for matching then fill grid until
		row is complete and go to the next row
	*/
	var nei, tmp uint64
	var ygrid, xgrid int
	left := next
	for range tiles {
		retToCol := false
		pasteTile(grid, ygrid, xgrid, maptiles[next])
		if xgrid+len(tiles[0].pix)-2 >= len(grid) {
			xgrid = 0
			ygrid += len(tiles[0].pix) - 2
			retToCol = true
		} else {
			xgrid += len(tiles[0].pix) - 2
		}
		if ygrid+len(tiles[0].pix)-2 > len(grid) {
			break
		}
		var val uint64
		if retToCol {
			tmp = left
			val = getBottom(maptiles[left])
		} else {
			tmp = next
			val = getRight(maptiles[next])
		}
		if cache[val][0] == tmp {
			nei = cache[val][1]
		} else {
			nei = cache[val][0]
		}
		for i := 0; i < 8; i++ {
			if !retToCol && maptiles[next].isRightCompatible(maptiles[nei]) {
				next = nei
				break
			} else if retToCol && maptiles[left].isBottomCompatible(maptiles[nei]) {
				next = nei
				left = nei
				break
			}
			if i == 3 {
				maptiles[nei] = maptiles[nei].inverse()
			} else {
				maptiles[nei] = maptiles[nei].rotate()
			}
		}
	}

	// Now find sea monster
	seaMonster := [][]uint64{
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2},
		{1, 2, 2, 2, 2, 1, 1, 2, 2, 2, 2, 1, 1, 2, 2, 2, 2, 1, 1, 1},
		{2, 1, 2, 2, 1, 2, 2, 1, 2, 2, 1, 2, 2, 1, 2, 2, 1, 2, 2, 2},
	}

	// fake the grid as a tile so we can easily rotate and inverse it
	tgrid := tile{pix: grid, id: 0}
	var seaMonsterCount int
	var isMonsterGrid bool
	for i := 0; i < 8; i++ {
		for y := 0; y < len(tgrid.pix)-len(seaMonster); y++ {
			for x := 0; x < len(tgrid.pix[0])-len(seaMonster[0]); x++ {
				if isSeaMonsterInGrid(seaMonster, tgrid.pix, x, y) {
					seaMonsterCount++
					isMonsterGrid = true
				}
			}
		}

		if isMonsterGrid {
			break
		}
		if i == 3 {
			tgrid = tgrid.inverse()
		} else {
			tgrid = tgrid.rotate()
		}
	}

	// Now that we have the grid that goes with the monster, count
	// each # (1) in the grid and substract the number of # in monster
	// multiplied by the number of monster

	return count1(tgrid.pix) - count1(seaMonster)*seaMonsterCount
}

func main() {
	var err error

	tiles, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	mul, corners := solve1(tiles)
	fmt.Println(mul)
	fmt.Println(solve2(tiles, corners))
}
