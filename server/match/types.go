package match

type MatchSnapshot struct {
}

type Point struct {
	x float32
	y float32
}

func (point *Point) isEqual(p *Point) bool {
	return int(point.x) == int(p.x) && int(point.y) == int(p.y)
}
