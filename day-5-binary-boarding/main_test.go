package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay5(t *testing.T) {
	testCases := map[string]struct {
		row    int
		column int
		seatID int
	}{
		"FBFBBFFRLR": {44, 5, 357},
		"BFFFBBFRRR": {70, 7, 567},
		"FFFBBBFRRR": {14, 7, 119},
		"BBFFBBFRLL": {102, 4, 820},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			row, col := calculateRowAndColumn(name)
			seatID := calculateSeatID(row, col)

			assert.Equal(t, tc.row, row)
			assert.Equal(t, tc.column, col)
			assert.Equal(t, tc.seatID, seatID)
		})
	}
}
