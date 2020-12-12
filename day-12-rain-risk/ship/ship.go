package ship

const (
	DirectionNorth = "N"
	DirectionSouth = "S"
	DirectionEast  = "E"
	DirectionWest  = "W"

	TurnLeft  = "L"
	TurnRight = "R"

	MoveForward = "F"
)

var facing = map[string]int{
	DirectionNorth: 0,
	DirectionEast:  90,
	DirectionSouth: 180,
	DirectionWest:  270,
}

var facingReverse = map[int]string{
	0:   DirectionNorth,
	90:  DirectionEast,
	180: DirectionSouth,
	270: DirectionWest,
}

type Coordinates struct {
	x int
	y int
}

type Ship interface {
	DoAction(action string, value int)
	Move(direction string, value int)
	Turn(direction string, value int)
	GetManhattanDistanceFromStart() float64
}
