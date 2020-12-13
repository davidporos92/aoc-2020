package main

import (
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"

	"github.com/davidporos92/aoc-2020/utils"
	"github.com/deanveloper/modmath/v1/bigmod"
)

func main() {
	var earliestLeaveTime int
	var err error
	part1Bus := 0
	part1WaitTime := math.MaxInt64
	var busses []bigmod.CrtEntry

	utils.NewReader("./input-1.dat").MustReadFile(func(line string) {
		if earliestLeaveTime == 0 {
			earliestLeaveTime, err = strconv.Atoi(line)
			if err != nil {
				panic(fmt.Sprintf("Cannot convert %s to int: %v", line, err))
			}

			return
		}

		buses := regexp.MustCompile(`[^,]+`).FindAllString(line, -1)
		for k, bus := range buses {
			if bus == "x" {
				continue
			}

			bus, err := strconv.Atoi(bus)
			if err != nil {
				panic(fmt.Sprintf("Cannot convert %s to int: %v", line, err))
			}

			waitTime := bus-earliestLeaveTime%bus
			if waitTime < part1WaitTime {
				part1WaitTime = waitTime
				part1Bus = bus
			}

			busses = append(busses, bigmod.CrtEntry{A: big.NewInt(int64(bus - k)), N: big.NewInt(int64(bus))})
		}
	})

	fmt.Printf("Solution 1: %d\n", part1WaitTime*part1Bus)
	fmt.Printf("Solution 2: %d\n", bigmod.SolveCrtMany(busses))
}
