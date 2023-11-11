package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func max(values []int) (int, int) {
	if len(values) == 0 {
		panic("Can't compute max of empty slice")
	}
	max := values[0]
	index := 0
	for i, n := range values {
		if n > max {
			max = n
			index = i
		}
	}
	return max, index
}

func remove(values []int, index int) ([]int, int) {
	return append(values[:index], values[index+1:]...), values[index]
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	s := make([]int, 1)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			s = append(s, 0)
			continue
		}
		num, err := strconv.Atoi(text)
		check(err)
		s[len(s)-1] += num
	}

	// a
	top, _ := max(s)
	fmt.Println("Top:", top)

	// b
	sort.Slice(s, func(i int, j int) bool {
		return s[i] > s[j]
	})
	s = s[0:3]
	fmt.Println("Top 3:", s)
	fmt.Println("Top 3 sum:", s[0]+s[1]+s[2])
}
