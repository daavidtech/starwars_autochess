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
	UintID   string `json:"unitId"`
	UnitType string `json:"unitType"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}

type UnitStartedMovingTo struct {
	UnitID string `json:"unitId"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

type UnitArrivedToPosition struct {
	UnitID string `json:"unitId"`
	X      int    `json:"x"`
	Y      int    `json:"Y"`
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
	UnitType string `json:"unit_type"`
	Level    int    `json:"level"`
	HP       int    `json:"hp"`
	Mana     int    `json:"mana"`
	Rank     int    `json:"rank"`
	Cost     int    `json:"cost"`
}

type ShopRefilled struct {
	ShopUnits []ShopUnit `json:"shop_units"`
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
	UnitArrivedToPosition *UnitArrivedToPosition `json:"unitArrivedToPosition"`
	UnitStartedAttacking  *UnitStartedAttacking  `json:"unitStartedAttacking"`
	UnitStoppedAttacking  *UnitStoppedAttacking  `json:"unitStoppedAttacking"`
	LaunchProjectile      *LaunchProjectile      `json:"launchParticle"`
	PlayerMoneyChanged    *PlayerMoneyChanged    `json:"playerMoneyChanged"`
	PlayerLevelChanged    *PlayerLevelChanged    `json:"playerLevelChanged"`
	PlayerHealthChanged   *PlayerHealthChanged   `json:"playerHealthChanged"`
	ShopRefilled          *ShopRefilled          `json:"shopRefilled"`
	MatchPhaseChanged     *MatchPhaseChanged     `json:"matchPhaseChanged"`
}

type BuyUnit struct {
	ShopUnitID string `json:"shopUnitId"`
}

type PlaceUnit struct {
	UnitID string
	X      int
	Y      int
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

type MessageFromClient struct {
	BuyUnit          *BuyUnit          `json:"buyUnit"`
	PlaceUnit        *PlaceUnit        `json:"placeUnit"`
	SellUnit         *SellUnit         `json:"sellUnit"`
	BuyLevelUp       *BuyLevelUp       `json:"buyLevelUp"`
	RecycleShopUnits *RecycleShopUnits `json:"recycleShopUnits"`
	JoinGame         *JoinGame         `json:"joinGame"`
	SeekMatch        *SeekMatch        `json:"seekMatch"`
}

type GameSnapshot struct {
}
