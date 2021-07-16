package match

type UnitBought struct {
}

type BarrackUnitAdded struct {
	PlayerID string

	Unit
}

type BarrackUnitRemoved struct {
	PlayerID string
	UnitID   string
}

type BarrackUnitUpgraded struct {
	PlayerID string

	Unit
}

type ShopRefilled struct {
	PlayerID  string
	ShopUnits []ShopUnit
}

type PhaseChanged struct {
	MatchPhase MatchPhase
}

type ShopUnitRemoved struct {
	PlayerID   string
	ShopUnitID int
}

type CountdownStarted struct {
	StartValue int
	Interval   float32
}

type UnitPlaced struct {
	PlayerID string
	UnitID   string
	X        int
	Y        int
}

type UnitStartedMovingTo struct {
	PlayerID string
	UnitID   string
	X        int
	Y        int
}

type UnitArrivedTo struct {
	PlayerID string
	UnitID   string
	X        int
	Y        int
}

type UnitDied struct {
	PlayerID string
	UnitID   string
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

type BattleUnitHealthChanged struct {
	PlayerID  string
	UnitID    string
	NewHealth int
}

type BattleUnitManaChanged struct {
	PlayerID string
	UnitID   string
	NewMana  int
}

type MatchEvent struct {
	UnitBought              *UnitBought
	BarrackUnitAdded        *BarrackUnitAdded
	BarrackUnitRemoved      *BarrackUnitRemoved
	BarrackUnitUpgraded     *BarrackUnitUpgraded
	ShopRefilled            *ShopRefilled
	PhaseChanged            *PhaseChanged
	ShopUnitRemoved         *ShopUnitRemoved
	CountdownStarted        *CountdownStarted
	UnitPlaced              *UnitPlaced
	UnitStartedMovingTo     *UnitStartedMovingTo
	UnitArrivedTo           *UnitArrivedTo
	UnitDied                *UnitDied
	RoundCreated            *RoundCreated
	RoundFinished           *RoundFinished
	PlayerJoined            *PlayerJoined
	PlayerLeft              *PlayerLeft
	BattleUnitHealthChanged *BattleUnitHealthChanged
	BattleUnitManaChanged   *BattleUnitManaChanged
}
