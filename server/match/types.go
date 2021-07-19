package match

type MatchPhase string

const (
	LobbyPhase     MatchPhase = "LobbyPhase"
	ShoppingPhase  MatchPhase = "ShoppingPhase"
	PlacementPhase MatchPhase = "PlacementPhase"
	BattlePhase    MatchPhase = "BattlePhase"
	EndPhase       MatchPhase = "EndPhase"
)

type MatchSnapshot struct {
}

type Point struct {
	X float32
	Y float32
}

func (point *Point) isEqual(p *Point) bool {
	return int(point.X) == int(p.X) && int(point.Y) == int(p.Y)
}
