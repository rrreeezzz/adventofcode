package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type node struct {
	value int
	prev  *node
	next  *node
}

func readFile(filename string) (*node, error) {
	var err error
	nums := &node{}
	front := nums

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	for _, c := range line {
		n, err := strconv.Atoi(string(c))
		if err != nil {
			return nil, err
		}

		nums.value = n
		nums.next = &node{}
		nums.next.prev = nums
		nums = nums.next
	}

	// link last and first element
	nums.prev.next = front
	front.prev = nums.prev

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return front, nil
}

func solve1(curr *node, iterations int, b bool) int {
	var top int
	cache := make(map[int]*node)
	tmp := curr
	for {
		tmp = tmp.next
		cache[tmp.value] = tmp
		if curr == tmp {
			break
		}
		if tmp.value > top {
			top = tmp.value
		}
	}

	for turn := 0; turn < iterations; turn++ {
		n1 := curr.next
		n3 := curr.next.next.next

		cachedel := make(map[int]*node)
		cachedel[n1.value] = n1
		cachedel[n1.next.value] = n1.next
		cachedel[n3.value] = n3

		curr.next = curr.next.next.next.next

		val := curr.value
		pivot := curr.next
		var found bool
		for !found {
			val--
			if _, ok := cachedel[val]; ok {
				continue
			} else if val < 1 {
				for i := top; i > top-4; i-- {
					if _, ok := cachedel[i]; ok {
						continue
					}
					pivot = cache[i]
					found = true
					break
				}
			} else {
				pivot = cache[val]
				found = true
			}
		}
		tmp := n3.next
		n3.next = pivot.next
		pivot.next = n1
		curr = tmp
	}

	n1 := cache[1]
	var num int
	if !b {
		var numstring string
		curr = n1.next
		for {
			numstring = numstring + strconv.Itoa(curr.value)
			curr = curr.next
			if curr == n1 {
				break
			}
		}
		num, _ = strconv.Atoi(numstring)
	} else {
		num = n1.next.value * n1.next.next.value
	}

	return num
}

func main() {
	var err error

	nums, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve1(nums, 100, false))

	nums, err = readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}
	curr := nums.prev
	for i := 10; i < 1000001; i++ {
		n := &node{value: i}
		curr.next = n
		curr = curr.next
	}
	curr.next = nums
	fmt.Println(solve1(nums, 10000000, true))
}
