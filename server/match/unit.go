package match

type Unit struct {
	UnitID     string
	UnitType   string
	Tier       int
	Rank       int
	HP         int
	Mana       int
	AttackRate int

	Placement *Point
}
