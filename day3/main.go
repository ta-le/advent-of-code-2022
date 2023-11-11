package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findMatchingRune(a string, b string) rune {
	for _, rune1 := range a {
		for _, rune2 := range b {
			if rune1 == rune2 {
				return rune1
			}
		}
	}
	return 0
}

func createSetOfItems(items string) (set map[rune]bool) {
	set = make(map[rune]bool)
	for _, item := range items {
		set[item] = true
	}
	return
}

func findMatchingRune3(a string, b string, c string) rune {
	first := createSetOfItems(a)
	second := createSetOfItems(b)
	third := createSetOfItems(c)
	for key := range first {
		if second[key] && third[key] {
			return key
		}
	}
	return 0
}

func mapRuneToInt(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	}
	if r >= 'A' && r <= 'Z' {
		return int(r - 'A' + 27)
	}
	return 0
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	sum2 := 0
	var group []string
	for i := 1; scanner.Scan(); i++ {
		text := strings.Trim(scanner.Text(), " ")
		first := text[:len(text)/2]
		second := text[len(text)/2:]

		match := findMatchingRune(first, second)
		sum += mapRuneToInt(match)

		group = append(group, text)
		if i%3 == 0 {
			match := findMatchingRune3(group[0], group[1], group[2])
			sum2 += mapRuneToInt(match)
			group = nil
		}
	}
	fmt.Println("Part 1", sum)
	fmt.Println("Part 2", sum2)
}
