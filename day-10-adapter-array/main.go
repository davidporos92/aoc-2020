package main

import (
	"fmt"
	"github.com/davidporos92/aoc-2020/utils"
	"sort"
)

func main() {
	input := utils.NewReader("./input-1.dat").MustReadIntSliceFromFile()
	sort.Ints(input)
	input = append([]int{0}, input...)
	input = append(input, input[len(input)-1]+3)
	solution1 := part1(input)
	solution2 := part2(input)

	fmt.Printf("Solution 1: %d\n", solution1)
	fmt.Printf("Solution 2: %d\n", solution2)
}

func part1(input []int) int {
	diff1 := 0
	diff3 := 0

	if input[0] == 1 {
		diff1++
	} else if input[0] == 3 {
		diff3++
	}

	for i := 1; i < len(input); i++ {
		switch input[i] - input[i-1] {
		case 1:
			diff1++
		case 3:
			diff3++
		}
	}

	return diff1 * diff3
}

func part2(input []int) int {
	a := make([]int, len(input))
	a[len(input)-1] = 1

	for i := len(input)-2; i >= 0; i-- {
		a[i] = a[i + 1]

		if i+3 < len(input) && input[i+3] <= input[i] + 3 {
			a[i] += a[i+3]
		}

		if i+2 < len(input) && input[i+2] <= input[i] + 3 {
			a[i] += a[i+2]
		}
	}

	return a[0]
}
