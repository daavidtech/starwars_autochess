package game

type Placement struct {
	x    int
	y    int
	unit Unit
}

type Player struct {
	id         string
	credits    int
	health     int
	placements []Placement
	barrack    []Unit
}
