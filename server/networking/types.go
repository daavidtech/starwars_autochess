package networking

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

type GamePhaseChanged struct {
	NewPhaseType string
}

type UnitBought struct {
	UnitType string `json:"unitType"`
}

type UnitUpgraded struct {
	UnitId string `json:"unitId"`
	Rank   int    `json:"rank"`
}

type UnitSold struct {
	UnitId string `json:"unitId"`
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

type MessageToClient struct {
	GamePhaseChanged      *GamePhaseChanged      `json:"gamePhaseChanged"`
	UnitBought            *UnitBought            `json:"unitBought"`
	UnitSold              *UnitSold              `json:"unitSold"`
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
}

type BuyUnit struct {
	ShopUnitID string `json:"shopUnitId"`
}

type PlaceUnit struct {
	UnitID string
}

type SellUnit struct {
	UnitID string
}

type BuyLevelUp struct {
}

type RecycleShopUnits struct{}

type MessageFromClient struct {
	BuyUnit          *BuyUnit          `json:"buyUnit"`
	PlaceUnit        *PlaceUnit        `json:"placeUnit"`
	SellUnit         *SellUnit         `json:"sellUnit"`
	BuyLevelUp       *BuyLevelUp       `json:"buyLevelUp"`
	RecycleShopUnits *RecycleShopUnits `json:"recycleShopUnits"`
}

type GameSnapshot struct {
}
