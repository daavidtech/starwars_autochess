package match

type MatchSnapshot struct {
}

type Point struct {
	X float32
	Y float32
}

func (point *Point) isEqual(p *Point) bool {
	return int(point.X) == int(p.X) && int(point.Y) == int(p.Y)
}
