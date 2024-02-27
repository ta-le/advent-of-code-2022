package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// initial solution
func main2() {
	file, err := os.Open("./input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		text := scanner.Text()
		var strPair []string = strings.Split(text, ",")
		if len(strPair) != 2 {
			panic("not a valid assignment pair")
		}
		var pair [2][2]int
		for i, v := range strPair {
			split := strings.Split(v, "-")
			first, err := strconv.Atoi(split[0])
			check(err)
			second, err := strconv.Atoi(split[1])
			check(err)
			pair[i] = [2]int{first, second}
		}

		var aContainsB bool = pair[0][0] <= pair[1][0] && pair[0][1] >= pair[1][1]
		var bContainsA bool = pair[0][0] >= pair[1][0] && pair[0][1] <= pair[1][1]
		if aContainsB || bContainsA {
			fmt.Println("match:", text)
			count++
		}
	}
	fmt.Println("Result:", count)
}

// cool solution
func main() {
	file, err := os.Open("./input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	count := 0
	count2 := 0

	for scanner.Scan() {
		var oneLeft, oneRight, twoLeft, twoRight int
		// use Sscanf to pass a template, parse into ints and write them into variables
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &oneLeft, &oneRight, &twoLeft, &twoRight)

		// Part One
		if (oneLeft >= twoLeft && oneRight <= twoRight) ||
			(oneLeft <= twoLeft && oneRight >= twoRight) {
			count++
		}

		// Part two
		if oneLeft <= twoRight && twoLeft <= oneRight {
			count2++
		}
	}
	fmt.Println("Result", count)
	fmt.Println("Result", count2)
}
