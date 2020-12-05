package main

import (
	"fmt"
	"github.com/davidporos92/aoc-2020/utils"
)

const tree = "#"

type Slope struct {
	moveRight int
	moveDown  int
	treeCount int
}

var slopes = []Slope{
	{
		moveRight: 1,
		moveDown:  1,
	},
	{
		moveRight: 3,
		moveDown:  1,
	},
	{
		moveRight: 5,
		moveDown:  1,
	},
	{
		moveRight: 7,
		moveDown:  1,
	},
	{
		moveRight: 1,
		moveDown:  2,
	},
}

func main() {
	treeMultiplication := 1
	myMap := utils.NewReader("./input-1.dat").MustReadStringMapFromFile()

	for _, slope := range slopes {
		currentPositionX := 0
		currentPositionY := 0

		for {
			if myMap[currentPositionY][currentPositionX] == tree {
				slope.treeCount++
			}

			currentPositionX += slope.moveRight
			currentPositionY += slope.moveDown

			if currentPositionY >= len(myMap) {
				break
			}

			if currentPositionX >= len(myMap[currentPositionY]) {
				currentPositionX -= len(myMap[currentPositionY])
			}
		}

		fmt.Printf("Tree count for slope: %+v\n", slope)
		treeMultiplication *= slope.treeCount
	}

	fmt.Printf("Tree multiplication: %d\n", treeMultiplication)
}
