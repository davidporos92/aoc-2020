package ship

import (
	"fmt"
	"math"
)

type Part1 struct {
	Facing   string
	Position Coordinates
}

func NewPart1() Ship {
	return &Part1{
		Facing:   DirectionEast,
		Position: Coordinates{0, 0},
	}
}

func (s *Part1) DoAction(action string, value int) {
	switch action {
	case MoveForward:
		s.Move(s.Facing, value)
	case TurnLeft, TurnRight:
		s.Turn(action, value)
	case DirectionEast, DirectionNorth, DirectionSouth, DirectionWest:
		s.Move(action, value)
	}
}

func (s *Part1) Move(direction string, value int) {
	switch direction {
	case DirectionNorth:
		s.Position.y += value
	case DirectionSouth:
		s.Position.y -= value
	case DirectionEast:
		s.Position.x += value
	case DirectionWest:
		s.Position.x -= value
	}
}

func (s *Part1) Turn(direction string, value int) {
	currentFacing := facing[s.Facing]

	switch direction {
	case TurnRight:
		currentFacing += value
	case TurnLeft:
		currentFacing -= value
	}

	if currentFacing < 0 {
		currentFacing += 360
	}

	currentFacing %= 360

	newFacing, exists := facingReverse[currentFacing]
	if !exists {
		panic(fmt.Sprintf("Facing does not exists: %d\nTurning: %s %d\nNew facing: %s\nCurrentFacing: %d", currentFacing, direction, value, s.Facing, facing[s.Facing]))
	}

	s.Facing = newFacing
}

func (s *Part1) GetManhattanDistanceFromStart() float64 {
	x := math.Abs(float64(s.Position.x))
	y := math.Abs(float64(s.Position.y))

	return x+y
}

