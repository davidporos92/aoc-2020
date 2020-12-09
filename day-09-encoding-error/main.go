package main

import (
	"fmt"
	"github.com/davidporos92/aoc-2020/utils"
)

const preambleLength = 25

func main() {
	input := utils.NewReader("./input-1.dat").MustReadIntSliceFromFile()
	solution1 := part1(input)
	solution2 := part2(input, solution1)

	fmt.Printf("Solution 1: %d\n", solution1)
	fmt.Printf("Solution 2: %d\n", solution2)
}

func part1(input []int) int {
	for i := preambleLength; i < len(input); i++ {
		if !isPartOfCypherList(input[i], input[i-preambleLength:i]) {
			return input[i]
		}
	}

	panic("No solution found")
}

func part2(input []int, search int) int {
	for i := 0; i < len(input)-2; i++ {
		for j := i + 2; j < len(input); j++ {
			list := input[i:j]
			sum := utils.SumInts(list)

			if sum > search {
				break
			}

			if sum == search {
				min, max := utils.MinMax(list)
				return min + max
			}
		}
	}

	panic("No solution found")
}

func isPartOfCypherList(search int, haystack []int) bool {
	for i := 0; i < len(haystack)-1; i++ {
		a := haystack[i]
		for j := i + 1; j < len(haystack); j++ {
			if a+haystack[j] == search {
				return true
			}
		}
	}

	return false
}
