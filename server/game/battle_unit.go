package game

type Placement struct {
	x int
	y int
}

type BattleUnit struct {
	unitId     string
	unitType   string
	tier       int
	rank       int
	hp         int
	mana       int
	attackRate int

	placement *Placement
}
