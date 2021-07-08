package match

import (
	"log"

	"github.com/google/uuid"
)

type Player struct {
	id       string
	credits  int
	health   int
	accepted bool
	xp       int
	units    map[string]*Unit
	shop     *Shop
}

func NewPlayer() *Player {
	shop := NewShop()

	shop.SetSize(5)

	return &Player{
		units: make(map[string]*Unit),
		shop:  shop,
	}
}

func (player *Player) GetID() string {
	return player.id
}

func (player *Player) AddShopUnit(shopUnit ShopUnit) []MatchEvent {
	count := player.countUnitType(shopUnit.UnitType, 1)

	if count != 2 {
		if player.IsBarrackFull() {
			return []MatchEvent{}
		}

		newUnitID := uuid.New().String()

		player.units[newUnitID] = &Unit{
			unitId:   newUnitID,
			unitType: shopUnit.UnitType,
			//tier:       shopUnit.Tier,
			rank:       1,
			hp:         shopUnit.HP,
			mana:       shopUnit.Mana,
			attackRate: shopUnit.AttackRate,
		}

		return []MatchEvent{
			MatchEvent{
				BarrackUnitAdded: &BarrackUnitAdded{
					UnitID:     newUnitID,
					UnitType:   shopUnit.UnitType,
					Rank:       1,
					HP:         shopUnit.HP,
					Mana:       shopUnit.Mana,
					AttackRate: shopUnit.AttackRate,
				},
			},
		}
	}

	events := []MatchEvent{}

	count = player.countUnitType(shopUnit.UnitType, 2)

	upgraded := false
	removeTwoRanks := 0
	removeOneRanks := 1
	upgradeRank := 1

	if count == 2 {
		upgradeRank = 2
		removeTwoRanks = 1
		removeOneRanks = 2
	}

	for unitID, unit := range player.units {
		if unit.unitType != shopUnit.UnitType {
			continue
		}

		if !upgraded && upgradeRank == unit.rank {
			unit.rank += 1
			unit.hp = shopUnit.HP
			unit.mana = shopUnit.Mana
			unit.attackRate = shopUnit.AttackRate

			log.Printf("Unit %v upgraded to rank %v", unitID, unit.rank)

			events = append(events, MatchEvent{
				BarrackUnitUpgraded: &BarrackUnitUpgraded{
					UnitID:   unitID,
					UnitType: shopUnit.UnitType,
					//Tier:       shopUnit.Tier,
					Rank:       unit.rank,
					HP:         shopUnit.HP,
					Mana:       shopUnit.Mana,
					AttackRate: shopUnit.AttackRate,
				},
			})

			upgraded = true

			continue
		}

		if unit.rank == 1 {
			if removeOneRanks == 0 {
				continue
			}

			removeOneRanks -= 1
		}

		if unit.rank == 2 {
			if removeTwoRanks == 0 {
				continue
			}

			removeTwoRanks -= 1
		}

		if unit.rank == 3 {
			continue
		}

		log.Println("Removing unit " + unitID)

		delete(player.units, unitID)

		events = append(events, MatchEvent{
			BarrackUnitRemoved: &BarrackUnitRemoved{
				UnitID: unitID,
			},
		})

	}

	return events
}

func (player *Player) countUnitType(unitType string, rank int) int {
	count := 0

	for _, battleUnit := range player.units {
		if battleUnit.unitType != unitType || battleUnit.rank != rank {
			continue
		}

		count += 1
	}

	return count
}

func (player *Player) getBattleUnits() []*Unit {
	battleUnits := []*Unit{}

	for _, unit := range player.units {
		battleUnits = append(battleUnits, unit)
	}

	return battleUnits
}

func (player *Player) GetUnit(unitID string) *Unit {
	return player.units[unitID]
}

func (player *Player) RemoveUnit(unitID string) {
	delete(player.units, unitID)
}

func (player *Player) AddXP(amount int) {
	player.xp += amount
}

func (player *Player) GetLevel() int {
	if player.xp == 0 {
		return 1
	}

	return player.xp/100 + 1
}

func (player *Player) UseCredits(amount int) {
	player.credits -= amount
}

func (player *Player) IsBarrackFull() bool {
	return len(player.units) > 8
}
