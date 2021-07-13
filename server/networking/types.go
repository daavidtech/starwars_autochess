package networking

import "github.com/daavidtech/starwars_autochess/match"

type CreateUnit struct {
	ID       string `json:"id"`
	UnitType string `json:"unitType"`
	Enemy    bool   `json:"enemy"`
	MaxHP    int    `json:"maxHp"`
	CurrHP   int    `json:"currHp"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}

type RemoveUnit struct {
	ID string `json:"id"`
}

type ChangeUnitPosition struct {
	ID string `json:"id"`
	X  int    `json:"x"`
	Y  int    `json:"y"`
}

type GamePhase string

const (
	BuyingPhase    GamePhase = "BuyingPhase"
	PlacementPhase           = "PlacementPhase"
	BattlePhase              = "BattlePhase"
)

type MatchPhaseChanged struct {
	MatchPhase match.MatchPhase `json:"matchPhase"`
}

type UnitAdded struct {
	UnitID     string `json:"unitId"`
	UnitType   string `json:"unitType"`
	Rank       int    `json:"rank"`
	HP         int    `json:"hp"`
	Mana       int    `json:"mana"`
	AttackRate int    `json:"attackRate"`
}

type UnitRemoved struct {
	UnitID string `json:"unitId"`
}

type UnitUpgraded struct {
	UnitID     string `json:"unitId"`
	Rank       int    `json:"rank"`
	HP         int    `json:"hp"`
	Mana       int    `json:"mana"`
	AttackRate int    `json:"attackRate"`
}

type UnitSold struct {
	UnitID string `json:"unitId"`
}

type StartTimerTimeChanged struct {
	NewTimerValue int `json:"newTimerValue"`
}

type UnitDied struct {
	UnitID string `json:"unitId"`
}

type UnitPlaced struct {
	UnitID string `json:"unitId"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

type UnitStartedMovingTo struct {
	UnitID string `json:"unitId"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

type UnitArrivedTo struct {
	UnitID string `json:"unitId"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

type UnitStartedAttacking struct {
	UnitID string `json:"unitId"`
}

type UnitStoppedAttacking struct {
	UnitID string `json:"unitId"`
}

type LaunchProjectile struct {
	StartPointX  int    `json:"startPointX"`
	StartPointY  int    `json:"startPointY"`
	EndPointX    int    `json:"endPointX"`
	EndPointY    int    `json:"endPointY"`
	Speed        int    `json:"speed"`
	ParticleType string `json:"particleType"`
}

type UnitTookDamage struct {
	UnitID string `json:"unitId"`
	Amount int    `json:"amount"`
}

type UnitUsedMana struct {
	UnitID string `json:"unitId"`
	Amount int    `json:"amount"`
}

type UnitUsedAbility struct {
	UnitID    string `json:"unitId"`
	AbilityID string `json:"abilityId"`
}

type PlayerMoneyChanged struct {
	PlayerID string `json:"playerId"`
	NewMoney int    `json:"newMoney"`
}

type PlayerLevelChanged struct {
	PlayerID string `json:"playerId"`
	NewLevel int    `json:"newLevel"`
}

type PlayerHealthChanged struct {
	PlayerID string `json:"playerId"`
	NewHP    int    `json:"newHp"`
}

type ShopUnit struct {
	ID       int    `json:"id"`
	UnitType string `json:"unit_type"`
	Level    int    `json:"level"`
	HP       int    `json:"hp"`
	Mana     int    `json:"mana"`
	Rank     int    `json:"rank"`
	Cost     int    `json:"cost"`
}

type ShopUnitRemoved struct {
	ShopUnitID int `json:"shopUnitId"`
}

type ShopRefilled struct {
	ShopUnits []ShopUnit `json:"shop_units"`
}

type CountdownStarted struct {
	StartValue int     `json:"startValue"`
	Interval   float32 `json:"interval"`
}

type BattleUnit struct {
	Team          int    `json:"team"`
	UnitID        string `json:"unitId"`
	UnitType      string `json:"unitType"`
	Rank          int    `json:"rank"`
	MaxHP         int    `json:"maxHp"`
	HP            int    `json:"hp"`
	MaxMana       int    `json:"maxMana"`
	Mana          int    `json:"mana"`
	AttackRate    int    `json:"attackRate"`
	AttackRange   int    `json:"attackRange"`
	AttackDamage  int    `json:"attackDamage"`
	InstantAttack bool   `json:"instantAttack"`
	MoveSpeed     int    `json:"moveSpeed"`
	Dead          bool   `json:"dead"`
	Placement     Point  `json:"placement"`
}

type RoundCreated struct {
	PlayerID string       `json:"playerId"`
	Units    []BattleUnit `json:"units"`
}

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Unit struct {
	Team       int    `json:"team"`
	UnitID     string `json:"unitId"`
	UnitType   string `json:"unitType"`
	Tier       int    `json:"tier"`
	Rank       int    `json:"rank"`
	HP         int    `json:"hp"`
	Mana       int    `json:"mana"`
	AttackRate int    `json:"attackRate"`
	Placement  *Point `json:"placement"`
}

type RoundFinished struct {
	PlayerID         string `json:"playerId"`
	NewCreditsAmount int    `json:"newCreditsAmount"`
	NewPlayerHealth  int    `json:"newPlayerHealth"`
	Units            []Unit `json:"units"`
}

type MessageToClient struct {
	UnitAdded             *UnitAdded             `json:"unitAdded"`
	UnitRemoved           *UnitRemoved           `json:"unitRemoved"`
	UnitSold              *UnitSold              `json:"unitSold"`
	UnitUpgraded          *UnitUpgraded          `json:"unitUpgraded"`
	StartTimerTimeChanged *StartTimerTimeChanged `json:"startTimerTimeChanged"`
	UnitDied              *UnitDied              `json:"unitDied"`
	UnitPlaced            *UnitPlaced            `json:"unitPlaced"`
	UnitTookDamage        *UnitTookDamage        `json:"unitTookDamage"`
	UnitUsedMana          *UnitUsedMana          `json:"unitUsedMana"`
	UnitUsedAbility       *UnitUsedAbility       `json:"unitUsedAbility"`
	UnitStartedMovingTo   *UnitStartedMovingTo   `json:"unitStartedMovingTo"`
	UnitArrivedTo         *UnitArrivedTo         `json:"unitArrivedToPosition"`
	UnitStartedAttacking  *UnitStartedAttacking  `json:"unitStartedAttacking"`
	UnitStoppedAttacking  *UnitStoppedAttacking  `json:"unitStoppedAttacking"`
	LaunchProjectile      *LaunchProjectile      `json:"launchParticle"`
	PlayerMoneyChanged    *PlayerMoneyChanged    `json:"playerMoneyChanged"`
	PlayerLevelChanged    *PlayerLevelChanged    `json:"playerLevelChanged"`
	PlayerHealthChanged   *PlayerHealthChanged   `json:"playerHealthChanged"`
	ShopUnitRemoved       *ShopUnitRemoved       `json:"shopUnitRemoved"`
	ShopRefilled          *ShopRefilled          `json:"shopRefilled"`
	CountdownStarted      *CountdownStarted      `json:"countdownStarted"`
	RoundCreated          *RoundCreated          `json:"roundCreated"`
	RoundFinished         *RoundFinished         `json:"roundFinished"`

	MatchPhaseChanged *MatchPhaseChanged `json:"matchPhaseChanged"`
}

type BuyUnit struct {
	ShopUnitIndex int `json:"shopUnitIndex"`
}

type PlaceUnit struct {
	UnitID string `json:"unitId"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

type SellUnit struct {
	UnitID string
}

type BuyLevelUp struct {
	ShopUnitID int `json:"shop_unit_id"`
}

type RecycleShopUnits struct{}

type JoinGame struct{}

type SeekMatch struct{}

type Login struct {
	Username string
	Password string
}

type StartMatch struct {
}

type MessageFromClient struct {
	BuyUnit          *BuyUnit          `json:"buyUnit"`
	PlaceUnit        *PlaceUnit        `json:"placeUnit"`
	SellUnit         *SellUnit         `json:"sellUnit"`
	BuyLevelUp       *BuyLevelUp       `json:"buyLevelUp"`
	RecycleShopUnits *RecycleShopUnits `json:"recycleShopUnits"`
	JoinGame         *JoinGame         `json:"joinGame"`
	SeekMatch        *SeekMatch        `json:"seekMatch"`
	Login            *Login            `json:"login"`
	StartMatch       *StartMatch       `json:"startMatch"`
}

type GameSnapshot struct {
}
