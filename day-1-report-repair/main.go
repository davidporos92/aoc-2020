package main

import (
	"fmt"
	"github.com/davidporos92/aoc-2020/utils"
	"log"
)

func main() {
	targetSum := 2020
	lines := utils.MustReadIntSliceFromFile("./input-1.dat", '\n', "\n\r")

	fmt.Printf("Solution 1: %d\n", puzzle1(lines, targetSum))
	fmt.Printf("Solution 2: %d\n", puzzle2(lines, targetSum))
}

func puzzle1(lines []int, targetSum int) int {
	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			if lines[i]+lines[j] == targetSum {
				return lines[i] * lines[j]
			}
		}
	}

	log.Fatalf("Target sum not found: %d", targetSum)
	return 0
}

func puzzle2(lines []int, targetSum int) int {
	for i := 0; i < len(lines)-2; i++ {
		for j := i + 1; j < len(lines)-1; j++ {
			for k := j + 1; k < len(lines); k++ {
				if lines[i]+lines[j]+lines[k] == targetSum {
					return lines[i] * lines[j] * lines[k]
				}
			}
		}
	}

	log.Fatalf("Target sum not found: %d", targetSum)
	return 0
}
