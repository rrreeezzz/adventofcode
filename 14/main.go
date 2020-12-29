package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	mask   string
	memory uint64
	value  uint64
}

func readFile(filename string) ([]*instruction, error) {
	var err error
	var insts []*instruction

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var curMask string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inst := &instruction{}

		splitted := strings.Split(line, " = ")
		if splitted[0] == "mask" {
			curMask = strings.TrimSpace(splitted[1])
			continue
		}

		inst.mask = curMask

		splitted[0] = strings.ReplaceAll(splitted[0], "mem[", "")
		splitted[0] = strings.ReplaceAll(splitted[0], "]", "")

		if inst.memory, err = strconv.ParseUint(splitted[0], 10, 64); err != nil {
			return nil, err
		}

		if inst.value, err = strconv.ParseUint(splitted[1], 10, 64); err != nil {
			return nil, err
		}

		insts = append(insts, inst)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return insts, nil
}

func applyMask(mask string, value uint64) uint64 {
	uintbytes := []byte(strconv.FormatUint(value, 2))
	maskbytes := []byte(mask)

	zeroToAdd := len(mask) - len(uintbytes)
	for i := 0; i < zeroToAdd; i++ {
		uintbytes = append([]byte{'0'}, uintbytes...)
	}

	i := len(maskbytes) - 1
	for i >= 0 {
		switch mask[i] {
		case '1', '0':
			uintbytes[i] = maskbytes[i]
		}
		i--
	}

	newval, err := strconv.ParseUint(string(uintbytes), 2, 64)
	if err != nil {
		return 0
	}

	return newval
}

func applyMask2(mask string, value uint64) []uint64 {
	var memories []uint64
	var masks [][]byte
	var m []byte
	var index []int
	maskbytes := []byte(mask)

	for i := 0; i < len(mask); i++ {
		m = append(m, 'X')
	}

	pos := 0
	for i := len(maskbytes) - 1; i >= 0; i-- {
		switch mask[i] {
		case '1':
			m[i] = '1'
		case 'X':
			m[i] = '0'
			index = append(index, i)
			pos++
		case '0':
		}
	}

	masks = append(masks, m)
	for i := 0; i < len(index); i++ {
		lenmask := len(masks)
		for j := 0; j < lenmask; j++ {
			tmp := make([]byte, len(masks[j]))
			copy(tmp, masks[j])
			tmp[index[i]] = '1'
			masks = append(masks, tmp)
		}
	}

	for _, elt := range masks {
		mem := applyMask(string(elt), value)
		memories = append(memories, mem)
	}

	return memories
}

func computeMemory(insts []*instruction) uint64 {
	var sum uint64
	memorymap := make(map[uint64]uint64)

	for _, inst := range insts {
		memorymap[inst.memory] = applyMask(inst.mask, inst.value)
	}

	for _, val := range memorymap {
		sum += val
	}

	return sum
}

func computeMemory2(insts []*instruction) uint64 {
	var sum uint64
	memorymap := make(map[uint64]uint64)

	for _, inst := range insts {
		memories := applyMask2(inst.mask, inst.memory)

		for _, memory := range memories {
			memorymap[memory] = inst.value
		}
	}

	for _, val := range memorymap {
		sum += val
	}

	return sum
}

func main() {
	var err error

	insts, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(computeMemory(insts))
	fmt.Println(computeMemory2(insts))
}
