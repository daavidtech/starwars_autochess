package match

import "testing"

func TestAddsNewShopUnitWhenNoOthers(t *testing.T) {
	player := NewPlayer()

	player.AddShopUnit(ShopUnit{
		UnitType: "unit_droid",
	})

	battleUnits := player.getBattleUnits()

	if len(battleUnits) != 1 {
		t.Errorf("Invalid number of battle units")
	}
}

func TestUpgradeUnit(t *testing.T) {
	player := Player{
		battleUnits: map[string]*BattleUnit{
			"1": &BattleUnit{
				unitType: "unit_droid",
				rank:     1,
			},
			"2": &BattleUnit{
				unitType: "unit_droid",
				rank:     1,
			},
		},
	}

	player.AddShopUnit(ShopUnit{
		UnitType: "unit_droid",
	})

	units := player.getBattleUnits()

	if len(units) != 1 {
		t.Errorf("Invalid number of units")
	}

	unit := units[0]

	if unit.rank != 2 {
		t.Errorf("Unit did not get upgraded correctly")
	}
}
