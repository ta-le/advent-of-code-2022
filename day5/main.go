package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func push(arr []string, toPush string) []string {
	arr = append(arr, toPush)
	return arr
}

func pop(arr []string, amount int) ([]string, []string) {
	arr, popped := arr[:len(arr)-amount], arr[len(arr)-amount:]
	return arr, popped
}

func move9000(arr [9][]string, from int, to int, amount int) [9][]string {
	newArr, popped := pop(arr[from-1], amount)
	arr[from-1] = newArr
	for i := range popped {
		arr[to-1] = push(arr[to-1], popped[len(popped)-1-i]) // push in reverse order
	}
	return arr
}

func move9001(arr [9][]string, from int, to int, amount int) [9][]string {
	newArr, popped := pop(arr[from-1], amount)
	arr[from-1] = newArr
	for _, v := range popped {
		arr[to-1] = push(arr[to-1], v) // push in normal order
	}
	return arr
}

func printStacks(stacks [9][]string) {
	for _, x := range stacks {
		fmt.Println(x)
	}
	fmt.Print("Result string: ")
	for _, x := range stacks {
		fmt.Print(x[len(x)-1])
	}
	fmt.Println("\n-----------")
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	scanner := bufio.NewScanner(file)

	var stacks [9][]string
	for scanner.Scan() {
		text := scanner.Text()
		if text == " 1   2   3   4   5   6   7   8   9 " {
			break
		}
		for i, r := range text {
			if (i-1)%4 == 0 && unicode.IsLetter(r) {
				stacks[(i-1)/4] = append([]string{string(r)}, stacks[(i-1)/4]...)
			}
		}
	}
	fmt.Println("BEFORE")
	printStacks(stacks)

	// skip empty line
	scanner.Scan()

	stacksA := stacks // Part One
	stacksB := stacks // Part Two

	for scanner.Scan() {
		text := scanner.Text()
		var amount, from, to int
		_, err := fmt.Sscanf(text, "move %d from %d to %d", &amount, &from, &to)
		check(err)
		stacksA = move9000(stacksA, from, to, amount)
		stacksB = move9001(stacksB, from, to, amount)
	}

	fmt.Println("AFTER (Part One)")
	printStacks(stacksA)

	fmt.Println("AFTER (Part Two)")
	printStacks(stacksB)
}
