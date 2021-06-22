package game

type MatchPhase string

const (
	Lobby          MatchPhase = "Lobby"
	ShoppingPhase  MatchPhase = "ShoppingPhase"
	PlacementPhase MatchPhase = "PlacementPhase"
	BattlePhase    MatchPhase = "BattlePhase"
)

type Match struct {
	id                 string
	currentRoundNumber int
	phase              MatchPhase
	players            []Player

	shop              Shop
	TierProbabilities TierProbabilities
}

func NewMatch() Match {
	return Match{}
}

func (match *Match) start() {

}
