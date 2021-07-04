package match

type UnitBought struct {
}

type BarrackUnitAdded struct {
	UnitID     string
	UnitType   string
	Rank       int
	HP         int
	Mana       int
	AttackRate int
}

type BarrackUnitRemoved struct {
	UnitID string
}

type BarrackUnitUpgraded struct {
	UnitID     string
	UnitType   string
	Tier       int
	Rank       int
	HP         int
	Mana       int
	AttackRate int
}

type ShopRefilled struct {
	ShopUnits []ShopUnit
}

type PhaseChanged struct {
	MatchPhase MatchPhase
}

type MatchEvent struct {
	UnitBought          *UnitBought
	BarrackUnitAdded    *BarrackUnitAdded
	BarrackUnitRemoved  *BarrackUnitRemoved
	BarrackUnitUpgraded *BarrackUnitUpgraded
	ShopRefilled        *ShopRefilled
	PhaseChanged        *PhaseChanged
}
