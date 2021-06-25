package game

import "github.com/google/uuid"

type Player struct {
	id          string
	credits     int
	health      int
	accepted    bool
	xp          int
	battleUnits map[string]*BattleUnit
}

func NewPlayer() *Player {
	return &Player{
		battleUnits: make(map[string]*BattleUnit),
	}
}

func (player *Player) AddShopUnit(shopUnit ShopUnit) {
	unitID := uuid.New().String()

	count := player.countUnitType(shopUnit.UnitType, 1)

	if count != 2 {
		player.battleUnits[unitID] = &BattleUnit{
			unitId:     unitID,
			unitType:   shopUnit.UnitType,
			tier:       shopUnit.Tier,
			rank:       1,
			hp:         shopUnit.HP,
			mana:       shopUnit.Mana,
			attackRate: shopUnit.AttackRate,
		}

		return
	}

	removedCount := 0

	for unitID, unit := range player.battleUnits {
		if unit.unitType != shopUnit.UnitType && unit.rank != 1 || removedCount > 2 {
			continue
		}

		removedCount += 1

		delete(player.battleUnits, unitID)
	}

	player.battleUnits[unitID] = &BattleUnit{
		unitId:     unitID,
		unitType:   shopUnit.UnitType,
		tier:       shopUnit.Tier,
		rank:       2,
		hp:         shopUnit.HP,
		mana:       shopUnit.Mana,
		attackRate: shopUnit.AttackRate,
	}
}

func (player *Player) countUnitType(unitType string, rank int) int {
	count := 0

	for _, battleUnit := range player.battleUnits {
		if battleUnit.unitType != unitType || battleUnit.rank != rank {
			continue
		}

		count += 1
	}

	return count
}

func (player *Player) getBattleUnits() []*BattleUnit {
	battleUnits := []*BattleUnit{}

	for _, unit := range player.battleUnits {
		battleUnits = append(battleUnits, unit)
	}

	return battleUnits
}

func (player *Player) GetUnit(unitID string) *BattleUnit {
	return player.battleUnits[unitID]
}

func (player *Player) RemoveUnit(unitID string) {
	delete(player.battleUnits, unitID)
}

func (player *Player) AddXP(amount int) {
	player.xp += amount
}

func (player *Player) GetLevel() int {
	return player.xp / 100
}

