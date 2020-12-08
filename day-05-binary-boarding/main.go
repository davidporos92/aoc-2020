package main

import (
	"fmt"
	"github.com/davidporos92/aoc-2020/utils"
	"math"
	"sort"
)

const (
	lower = iota
	upper

	front = "F"
	back  = "B"
	left  = "L"
	right = "R"

	maxRows    int = 127
	maxColumns int = 7
)

func main() {
	maxSeatID := 0.0
	seatIDs := make([]float64, 0)

	utils.NewReader("./input-1.dat").MustReadFile(func(line string) {
		seatID := float64(calculateSeatID(calculateRowAndColumn(line)))

		seatIDs = append(seatIDs, seatID)
		maxSeatID = math.Max(seatID, maxSeatID)
	})

	fmt.Printf("Max SeatID: %0.f\n", maxSeatID)

	sort.Float64s(seatIDs)
	for i := 1; i < len(seatIDs)-1; i++ {
		next := seatIDs[i+1]
		actual := seatIDs[i]

		if actual+1 != next && actual+2 == next {
			fmt.Printf("My seat: %0.f\n", actual+1)
			return
		}
	}
}

func calculateRowAndColumn(seat string) (int, int) {
	rowRange := map[uint]int{
		lower: 0,
		upper: maxRows,
	}
	columnRange := map[uint]int{
		lower: 0,
		upper: maxColumns,
	}
	for _, char := range []rune(seat) {
		switch string(char) {
		case back:
			rowRange[lower] += (rowRange[upper]-rowRange[lower])/2 + 1
		case front:
			rowRange[upper] -= (rowRange[upper]-rowRange[lower])/2 + 1
		case right:
			columnRange[lower] += (columnRange[upper]-columnRange[lower])/2 + 1
		case left:
			columnRange[upper] -= (columnRange[upper]-columnRange[lower])/2 + 1
		}
	}

	return rowRange[upper], columnRange[upper]
}

func calculateSeatID(row, column int) int {
	return (row * 8) + column
}
