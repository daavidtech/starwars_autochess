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

type ShopUnitRemoved struct {
	ShopUnitID int
}

type CountdownStarted struct {
	StartValue int
	Interval   float32
}

type UnitPlaced struct {
	UnitID string
	X      int
	Y      int
}

type UnitStartedMovingTo struct {
	UnitID string
	X      int
	Y      int
}

type UnitArrivedTo struct {
	UnitID string
	X      int
	Y      int
}

type UnitDied struct {
	UnitID string
}

type RoundCreated struct {
	PlayerID string
	Units    []BattleUnit
}

type RoundFinished struct {
	PlayerID         string
	NewCreditsAmount int
	NewPlayerHealth  int
	Units            []Unit
}

type PlayerJoined struct {
	Player Player
}

type PlayerLeft struct {
	Player Player
}

type MatchEvent struct {
	UnitBought          *UnitBought
	BarrackUnitAdded    *BarrackUnitAdded
	BarrackUnitRemoved  *BarrackUnitRemoved
	BarrackUnitUpgraded *BarrackUnitUpgraded
	ShopRefilled        *ShopRefilled
	PhaseChanged        *PhaseChanged
	ShopUnitRemoved     *ShopUnitRemoved
	CountdownStarted    *CountdownStarted
	UnitPlaced          *UnitPlaced
	UnitStartedMovingTo *UnitStartedMovingTo
	UnitArrivedTo       *UnitArrivedTo
	UnitDied            *UnitDied
	RoundCreated        *RoundCreated
	RoundFinished       *RoundFinished
	PlayerJoined        *PlayerJoined
	PlayerLeft          *PlayerLeft
}
