package main

import (
	"fmt"
	"github.com/davidporos92/aoc-2020/utils"
)

const (
	SeatEmpty    = "L"
	SeatOccupied = "#"
	SeatFloor    = "."
)

type Coordinates struct {
	x int
	y int
}

type Seats struct {
	seats    map[Coordinates]string
	maxX     int
	maxY     int
	tolerance int
}

func NewSeats(input [][]string) *Seats {
	s := &Seats{seats: make(map[Coordinates]string, len(input))}

	for y := range input {
		for x := range input[y] {
			s.seats[Coordinates{x, y}] = input[y][x]

			if x > s.maxX {
				s.maxX = x
			}
		}

		if y > s.maxY {
			s.maxY = y
		}
	}

	return s
}

func (s *Seats) SetTolerance(t int) {
	s.tolerance = t
}

func (s *Seats) SetEmpty(c Coordinates) {
	s.seats[c] = SeatEmpty
}

func (s *Seats) SetOccupied(c Coordinates) {
	s.seats[c] = SeatOccupied
}

func (s *Seats) TotalOccupied() int {
	count := 0

	for c := range s.seats {
		if s.IsOccupied(c) {
			count++
		}
	}

	return count
}

func (s *Seats) IsEmpty(c Coordinates) bool {
	return s.seats[c] == SeatEmpty
}

func (s *Seats) IsOccupied(c Coordinates) bool {
	return s.seats[c] == SeatOccupied
}

func (s *Seats) IsFloor(c Coordinates) bool {
	return s.seats[c] == SeatFloor
}

func (s *Seats) NumberOfOccupiedSeatsAdjacent(c Coordinates) int {
	count := 0

	for y := c.y - 1; y <= c.y+1; y++ {
		for x := c.x - 1; x <= c.x+1; x++ {
			if x == c.x && y == c.y {
				continue
			}

			c := Coordinates{x, y}
			if _, exists := s.seats[c]; exists && s.IsOccupied(c) {
				count++
			}
		}
	}

	return count
}

func (s *Seats) NumberOfOccupiedSeatsSeen(c Coordinates) int {
	count := 0
	distance := 1
	directionsDone := map[string]Coordinates{
		"cUpLeft": {c.x, c.y},
		"cUp": {c.x, c.y},
		"cUpRight": {c.x, c.y},
		"cRight": {c.x, c.y},
		"cDownRight": {c.x, c.y},
		"cDown": {c.x, c.y},
		"cDownLeft": {c.x, c.y},
		"cLeft": {c.x, c.y},
	}

	for len(directionsDone) > 0 {
		if _, e := directionsDone["cUpLeft"]; e { directionsDone["cUpLeft"] = Coordinates{c.x-distance, c.y-distance} }
		if _, e := directionsDone["cUp"]; e { directionsDone["cUp"] = Coordinates{c.x, c.y-distance} }
		if _, e := directionsDone["cUpRight"]; e { directionsDone["cUpRight"] = Coordinates{c.x+distance, c.y-distance} }
		if _, e := directionsDone["cRight"]; e { directionsDone["cRight"] = Coordinates{c.x+distance, c.y} }
		if _, e := directionsDone["cDownRight"]; e { directionsDone["cDownRight"] = Coordinates{c.x+distance, c.y+distance} }
		if _, e := directionsDone["cDown"]; e { directionsDone["cDown"] = Coordinates{c.x, c.y+distance} }
		if _, e := directionsDone["cDownLeft"]; e { directionsDone["cDownLeft"] = Coordinates{c.x-distance, c.y+distance} }
		if _, e := directionsDone["cLeft"]; e { directionsDone["cLeft"] = Coordinates{c.x-distance, c.y} }

		for k, dir := range directionsDone {
			if _, e := s.seats[dir]; e && s.IsOccupied(dir) {
				count++
				delete(directionsDone, k)
			} else if !e || s.IsEmpty(dir) {
				delete(directionsDone, k)
			}
		}

		distance++
	}

	return count
}

func (s *Seats) Iterate() *Seats {
	newState := &Seats{seats: make(map[Coordinates]string), maxX: s.maxX, maxY: s.maxY, tolerance: s.tolerance}

	for c := range s.seats {
		empty := s.IsEmpty(c)
		occupied := s.IsOccupied(c)
		numOfOccupied := s.NumberOfOccupiedSeatsAdjacent(c)

		if empty && numOfOccupied == 0 {
			newState.SetOccupied(c)
			continue
		}

		if occupied && numOfOccupied >= s.tolerance {
			newState.SetEmpty(c)
			continue
		}

		newState.seats[c] = s.seats[c]
	}

	return newState
}

func (s *Seats) IterateWithSeen() *Seats {
	newState := &Seats{seats: make(map[Coordinates]string), maxX: s.maxX, maxY: s.maxY, tolerance: s.tolerance}

	for c := range s.seats {
		empty := s.IsEmpty(c)
		occupied := s.IsOccupied(c)
		numOfOccupied := s.NumberOfOccupiedSeatsSeen(c)

		if empty && numOfOccupied == 0 {
			newState.SetOccupied(c)
			continue
		}

		if occupied && numOfOccupied >= s.tolerance {
			newState.SetEmpty(c)
			continue
		}

		newState.seats[c] = s.seats[c]
	}

	return newState
}

func (s *Seats) IsDifferFrom(state *Seats) bool {
	for c := range s.seats {
		if v, exists := state.seats[c]; !exists || v != s.seats[c] {
			return true
		}
	}

	return false
}

func (s *Seats) String() string {
	str := ""

	for y := 0; y <= s.maxY; y++ {
		for x := 0; x <= s.maxX; x++ {
			str += fmt.Sprintf("%s", s.seats[Coordinates{x, y}])
		}

		str += fmt.Sprintf("\n")
	}

	return str
}

func main() {
	var newSeats *Seats
	input := utils.NewReader("./input-1.dat").MustReadStringMapFromFile()

	// Part 1
	seats := NewSeats(input)
	seats.tolerance = 4

	for {
		newSeats = seats.Iterate()
		if !newSeats.IsDifferFrom(seats) {
			break
		}

		seats = newSeats
	}
	fmt.Printf("Solution 1: %d\n", newSeats.TotalOccupied())

	// Part 2
	seats = NewSeats(input)
	seats.tolerance = 5

	for {
		newSeats = seats.IterateWithSeen()
		if !newSeats.IsDifferFrom(seats) {
			break
		}

		seats = newSeats
	}

	fmt.Printf("Solution 2: %d\n", newSeats.TotalOccupied())
}
