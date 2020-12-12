package ship

import (
	"math"
)

type Part2 struct {
	Facing   string
	Position Coordinates
	Waypoint Coordinates
}

func NewPart2() Ship {
	return &Part2{
		Facing:   DirectionEast,
		Position: Coordinates{0, 0},
		Waypoint: Coordinates{10, 1},
	}
}

func (s *Part2) DoAction(action string, value int) {
	switch action {
	case MoveForward:
		s.Position.x += s.Waypoint.x * value
		s.Position.y += s.Waypoint.y * value
	case TurnLeft, TurnRight:
		s.Turn(action, value)
	case DirectionEast, DirectionNorth, DirectionSouth, DirectionWest:
		s.Move(action, value)
	}
}

func (s *Part2) Move(direction string, value int) {
	switch direction {
	case DirectionNorth:
		s.Waypoint.y += value
	case DirectionSouth:
		s.Waypoint.y -= value
	case DirectionEast:
		s.Waypoint.x += value
	case DirectionWest:
		s.Waypoint.x -= value
	}
}

func (s *Part2) Turn(direction string, value int) {
	turn := value / 90

	switch direction {
	case TurnLeft:
		turn = 4 - turn
	}

	switch turn {
	case 1:
		s.Waypoint.x, s.Waypoint.y = s.Waypoint.y, -s.Waypoint.x
	case 2:
		s.Waypoint.x, s.Waypoint.y = -s.Waypoint.x, -s.Waypoint.y
	case 3:
		s.Waypoint.x, s.Waypoint.y = -s.Waypoint.y, s.Waypoint.x
	}
}

func (s *Part2) GetManhattanDistanceFromStart() float64 {
	x := math.Abs(float64(s.Position.x))
	y := math.Abs(float64(s.Position.y))

	return x + y
}
