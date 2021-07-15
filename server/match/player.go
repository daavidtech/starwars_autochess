package match

import (
	"log"

	"github.com/google/uuid"
)

type Player struct {
	id         string
	name       string
	credits    int
	health     int
	accepted   bool
	dead       bool
	xp         int
	lobbyAdmin bool
	units      map[string]*Unit
	shop       *Shop
}

func NewPlayer() *Player {
	shop := NewShop()

	shop.SetSize(5)

	return &Player{
		id:     uuid.NewString(),
		units:  make(map[string]*Unit),
		shop:   shop,
		health: 100,
		dead:   false,
	}
}

func (player *Player) GetID() string {
	return player.id
}

func (player *Player) GetName() string {
	return player.name
}

func (player *Player) SetName(name string) {
	player.name = name
}

func (player *Player) AddShopUnit(shopUnit ShopUnit) []MatchEvent {
	count := player.countUnitType(shopUnit.UnitType, 1)

	if count != 2 {
		if player.IsBarrackFull() {
			return []MatchEvent{}
		}

		newUnitID := uuid.New().String()

		player.units[newUnitID] = &Unit{
			UnitID:   newUnitID,
			UnitType: shopUnit.UnitType,
			//tier:       shopUnit.Tier,
			Rank:       1,
			HP:         shopUnit.HP,
			Mana:       shopUnit.Mana,
			AttackRate: shopUnit.AttackRate,
		}

		return []MatchEvent{
			MatchEvent{
				BarrackUnitAdded: &BarrackUnitAdded{
					PlayerID:   player.id,
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
		if unit.UnitType != shopUnit.UnitType {
			continue
		}

		if !upgraded && upgradeRank == unit.Rank {
			unit.Rank += 1
			unit.HP = shopUnit.HP
			unit.Mana = shopUnit.Mana
			unit.AttackRate = shopUnit.AttackRate

			log.Printf("Unit %v upgraded to rank %v", unitID, unit.Rank)

			events = append(events, MatchEvent{
				BarrackUnitUpgraded: &BarrackUnitUpgraded{
					PlayerID: player.id,
					UnitID:   unitID,
					UnitType: shopUnit.UnitType,
					//Tier:       shopUnit.Tier,
					Rank:       unit.Rank,
					HP:         shopUnit.HP,
					Mana:       shopUnit.Mana,
					AttackRate: shopUnit.AttackRate,
				},
			})

			upgraded = true

			continue
		}

		if unit.Rank == 1 {
			if removeOneRanks == 0 {
				continue
			}

			removeOneRanks -= 1
		}

		if unit.Rank == 2 {
			if removeTwoRanks == 0 {
				continue
			}

			removeTwoRanks -= 1
		}

		if unit.Rank == 3 {
			continue
		}

		log.Println("Removing unit " + unitID)

		delete(player.units, unitID)

		events = append(events, MatchEvent{
			BarrackUnitRemoved: &BarrackUnitRemoved{
				PlayerID: player.id,
				UnitID:   unitID,
			},
		})

	}

	return events
}

func (player *Player) countUnitType(unitType string, rank int) int {
	count := 0

	for _, battleUnit := range player.units {
		if battleUnit.UnitType != unitType || battleUnit.Rank != rank {
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

func (player *Player) payDay() {
	player.credits += 100
}
