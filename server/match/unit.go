package match

type Unit struct {
	unitId     string
	unitType   string
	tier       int
	rank       int
	hp         int
	mana       int
	attackRate int

	placement *Point
}
