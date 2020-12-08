package main

import (
	"fmt"
	"github.com/davidporos92/aoc-2020/utils"
	"strings"
)

type Group struct {
	answers map[string]int
	personCount int
}

func (g *Group) CountAnswersSolution1() int {
	return len(g.answers)
}

func (g *Group) CountAnswersSolution2() int {
	count := 0
	for _, c := range g.answers {
		if c == g.personCount {
			count++
		}
	}
	return count
}

func main() {
	personGroups := utils.NewReader("./input-1.dat").MustReadStringBatchesFromFile(utils.BatchSeparatorBlankLine)
	solution1 := 0
	solution2 := 0

	for _, persons := range personGroups {
		group := &Group{
			personCount: len(strings.Split(persons, " ")),
			answers: utils.GetUniqueCharCounts(strings.ReplaceAll(persons, " ", "")),
		}

		solution1 += group.CountAnswersSolution1()
		solution2 += group.CountAnswersSolution2()
	}

	fmt.Printf("Solution 1: %d\n", solution1)
	fmt.Printf("Solution 2: %d\n", solution2)
}
