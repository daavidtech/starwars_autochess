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
		units: map[string]*Unit{
			"1": &Unit{
				unitType: "unit_droid",
				rank:     1,
			},
			"2": &Unit{
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

func TestCannotUpgradeUnitTooBigRank(t *testing.T) {
	player := Player{
		units: map[string]*Unit{
			"1": &Unit{
				unitId:   "1",
				unitType: "unit_droid",
				rank:     2,
			},
			"2": &Unit{
				unitId:   "2",
				unitType: "unit_droid",
				rank:     1,
			},
		},
	}

	player.AddShopUnit(ShopUnit{
		UnitType: "unit_droid",
	})

	units := player.getBattleUnits()

	if len(units) != 3 {
		t.Error("Invalid number of units")
	}

	unit1 := player.GetUnit("1")

	if unit1.rank != 2 {
		t.Errorf("Unit 1 has invalid rank %v", unit1.rank)
	}

	unit2 := player.GetUnit("2")

	if unit2.rank != 1 {
		t.Error("Unit 2 has invalid rank")
	}
}

func Test_updates_unit_rank_to_3(t *testing.T) {
	player := Player{
		units: map[string]*Unit{
			"1": &Unit{
				unitType: "unit_droid",
				rank:     2,
			},
			"2": &Unit{
				unitType: "unit_droid",
				rank:     2,
			},
			"3": &Unit{
				unitType: "unit_droid",
				rank:     1,
			},
			"4": &Unit{
				unitType: "unit_droid",
				rank:     1,
			},
		},
	}

	events := player.AddShopUnit(ShopUnit{
		UnitType: "unit_droid",
	})

	if len(events) != 4 {
		t.Error("Invalid number of events")
	}

	units := player.getBattleUnits()

	if len(units) != 1 {
		t.Error("Invalid number of units")
	}

	rank3Unit := false

	for _, unit := range units {
		if unit.rank == 3 {
			rank3Unit = true
		}
	}

	if !rank3Unit {
		t.Error("Unit did not get upgraded correctly")
	}
}

func Test_diffrent_unit_types_are_ignored(t *testing.T) {
	player := Player{
		units: map[string]*Unit{
			"1": &Unit{
				unitType: "unit_droid",
				rank:     1,
			},
			"2": &Unit{
				unitType: "unit_clone",
				rank:     1,
			},
		},
	}

	events := player.AddShopUnit(ShopUnit{
		UnitType: "unit_droid",
	})

	if len(events) != 1 {
		t.Error("Invalid number of events")
	}

	units := player.getBattleUnits()

	if len(units) != 3 {
		t.Error("Invalid number of units")
	}
}

func Test_cannot_upgrade_to_rank3_when_different_unit_type(t *testing.T) {
	player := Player{
		units: map[string]*Unit{
			"1": &Unit{
				unitType: "unit_droid",
				rank:     2,
			},
			"2": &Unit{
				unitType: "unit_clone",
				rank:     2,
			},
			"4": &Unit{
				unitType: "unit_clone",
				rank:     1,
			},
			"5": &Unit{
				unitType: "unit_droid",
				rank:     1,
			},
			"6": &Unit{
				unitType: "unit_droid",
				rank:     1,
			},
		},
	}

	events := player.AddShopUnit(ShopUnit{
		UnitType: "unit_droid",
	})

	if len(events) != 2 {
		t.Errorf("Invalid number of events %v", len(events))
	}

	upgradeEvent := false
	removeEvent := false

	for _, event := range events {
		if event.BarrackUnitUpgraded != nil {
			upgradeEvent = true
		}

		if event.BarrackUnitRemoved != nil {
			removeEvent = true
		}
	}

	if !upgradeEvent {
		t.Error("Upgrade event not found")
	}

	if !removeEvent {
		t.Error("Remove event not found")
	}

	units := player.getBattleUnits()

	if len(units) != 4 {
		t.Errorf("Invalid number of units %v", len(units))
	}

	droidRank2Units := 0

	for _, unit := range units {
		if unit.rank == 2 && unit.unitType == "unit_droid" {
			droidRank2Units += 1
		}
	}

	if droidRank2Units != 2 {
		t.Error("Unit was not upgraded correctly")
	}
}

func Test_upgrades_to_rank2_unit_when_there_is_rank3_unit(t *testing.T) {
	player := Player{
		units: map[string]*Unit{
			"1": &Unit{
				unitType: "unit_droid",
				rank:     3,
			},
			"2": &Unit{
				unitType: "unit_droid",
				rank:     1,
			},
			"3": &Unit{
				unitType: "unit_droid",
				rank:     1,
			},
		},
	}

	events := player.AddShopUnit(ShopUnit{
		UnitType: "unit_droid",
	})

	if len(events) != 2 {
		t.Errorf("Invalid number of events %v", len(events))
	}

	units := player.getBattleUnits()

	if len(units) != 2 {
		t.Errorf("Invalid number of units %v", len(units))
	}

	rank3Unit := 0
	rank2Unit := 0

	for _, unit := range units {
		if unit.rank == 2 {
			rank2Unit += 1
		}

		if unit.rank == 3 {
			rank3Unit += 1
		}
	}

	if rank2Unit != 1 {
		t.Errorf("Invalid number of rank2Units %v", rank2Unit)
	}

	if rank3Unit != 1 {
		t.Errorf("Invalid number of rank3Units %v", rank3Unit)
	}
}
