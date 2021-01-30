package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type food struct {
	ingredients []string
	allergens   []string
}

func readFile(filename string) ([]food, error) {
	var err error
	var foods []food

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		f := food{}

		var isAllergen bool
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		for _, w := range strings.Split(line, " ") {
			if w == "contains" {
				isAllergen = true
				continue
			}
			if isAllergen {
				w = strings.ReplaceAll(w, ",", "")
				f.allergens = append(f.allergens, w)
			} else {
				f.ingredients = append(f.ingredients, w)
			}
		}

		foods = append(foods, f)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return foods, nil
}

func isInArray(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func countArray(a string, list []string) int {
	var n int
	for _, b := range list {
		if b == a {
			n++
		}
	}
	return n
}

func solve(foods []food) (int, string) {
	var num int

	cache := make(map[string][]string)

	for _, f := range foods {
		for _, a := range f.allergens {
			if _, ok := cache[a]; !ok {
				cache[a] = f.ingredients
				continue
			}
			toKeep := []string{}
			for _, i := range f.ingredients {
				if isInArray(i, cache[a]) {
					toKeep = append(toKeep, i)
				}
			}

			cache[a] = toKeep
		}
	}

	var allIngredients []string
	for _, f := range foods {
		allIngredients = append(allIngredients, f.ingredients...)
	}

	num = len(allIngredients)
	counted := make(map[string]bool)
	for _, a := range cache {
		for _, i := range a {
			if ok := counted[i]; ok {
				continue
			}
			num -= countArray(i, allIngredients)
			counted[i] = true
		}
	}

	// Part 2
	ingToAll := make(map[string]string)
	allToIng := make(map[string]string)
	for len(cache) > 0 {
		for a, is := range cache {
			if len(is) == 1 {
				ingToAll[is[0]] = a
				allToIng[a] = is[0]
				delete(cache, a)
				break
			}

			for idx, i := range is {
				if _, ok := ingToAll[i]; ok {
					cache[a] = append(cache[a][:idx], cache[a][idx+1:]...)
				}
			}
		}
	}

	allergensOrder := []string{}
	dangerous := []string{}
	for _, i := range ingToAll {
		allergensOrder = append(allergensOrder, i)
	}
	sort.Strings(allergensOrder)

	for _, a := range allergensOrder {
		dangerous = append(dangerous, allToIng[a])
	}

	return num, strings.Join(dangerous, ",")
}

func main() {
	var err error

	foods, err := readFile("data.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(foods))
}
