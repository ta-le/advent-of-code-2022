package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const LOSS = 0
const DRAW = 3
const WIN = 6
const ROCK = 1
const PAPER = 2
const SCISSOR = 3

// Part one
var points = map[string]int{
	"X": ROCK,
	"Y": PAPER,
	"Z": SCISSOR,
}

var matchup = map[string]int{
	"A X": DRAW,
	"A Y": WIN,
	"A Z": LOSS,
	"B X": LOSS,
	"B Y": DRAW,
	"B Z": WIN,
	"C X": WIN,
	"C Y": LOSS,
	"C Z": DRAW,
}

// Part two
var points2 = map[string]int{
	"X": LOSS,
	"Y": DRAW,
	"Z": WIN,
}
var matchup2 = map[string]int{
	"A X": SCISSOR,
	"A Y": ROCK,
	"A Z": PAPER,
	"B X": ROCK,
	"B Y": PAPER,
	"B Z": SCISSOR,
	"C X": PAPER,
	"C Y": SCISSOR,
	"C Z": ROCK,
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	sum := 0
	sum2 := 0
	for scanner.Scan() {
		text := scanner.Text()
		s := strings.Fields(text)[1]

		// Part one
		sum += matchup[text] + points[s]

		// Part two
		sum2 += matchup2[text] + points2[s]
	}
	fmt.Println("Part one", (sum))
	fmt.Println("Part two", (sum2))
}
