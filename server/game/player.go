package game

type Player struct {
	id          string
	credits     int
	health      int
	accepted    bool
	battleUnits []BattleUnit
}

func (player *Player) AddShopUnit(shopUnit ShopUnit) {
	count := player.countUnitType(shopUnit.UnitType, 1)

	if count != 2 {
		player.battleUnits = append(player.battleUnits, BattleUnit{
			unitId:     shopUnit.UnitID,
			unitType:   shopUnit.UnitType,
			tier:       shopUnit.Tier,
			rank:       1,
			hp:         shopUnit.HP,
			mana:       shopUnit.Mana,
			attackRate: shopUnit.AttackRate,
		})

		return
	}

	newBattleUnits := []BattleUnit{}

	removedCount := 0

	for _, unit := range player.battleUnits {
		if unit.unitType == shopUnit.UnitType || unit.rank == 1 || removedCount < 2 {
			removedCount += 1

			continue
		}

		newBattleUnits = append(newBattleUnits, unit)
	}

	newBattleUnits = append(newBattleUnits, BattleUnit{
		unitId:     shopUnit.UnitID,
		unitType:   shopUnit.UnitType,
		tier:       shopUnit.Tier,
		rank:       2,
		hp:         shopUnit.HP,
		mana:       shopUnit.Mana,
		attackRate: shopUnit.AttackRate,
	})

	player.battleUnits = newBattleUnits
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

func (player *Player) getBattleUnits() []BattleUnit {
	return player.battleUnits
}
