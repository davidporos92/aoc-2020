package main

import (
	"fmt"
	"github.com/davidporos92/aoc-2020/day-12-rain-risk/ship"
	"github.com/davidporos92/aoc-2020/utils"
	"strconv"
	"strings"
)

func main() {
	ship1 := ship.NewPart1()
	ship2 := ship.NewPart2()

	utils.NewReader("./input-1.dat").MustReadFile(func(nav string) {
		action := string([]rune(nav)[0])
		value, err := strconv.Atoi(strings.Trim(nav, action))
		if err != nil {
			panic(err)
		}

		ship1.DoAction(action, value)
		ship2.DoAction(action, value)
	})

	fmt.Printf("Solution 1: %0.f\n", ship1.GetManhattanDistanceFromStart())
	fmt.Printf("Solution 2: %0.f\n", ship2.GetManhattanDistanceFromStart())
}
